import { Prisma } from "@prisma/client";
import bcrypt from 'bcryptjs';
import jwt from "jsonwebtoken";

import { prisma } from "app";
import { config } from "config";
import { UserAuthParams, AuthError, TokenInfo } from "./authModel";

function decodeAuth(user:UserAuthParams):[string, string]{
  const [username, password] = Buffer.from(user.basicAuth.split(" ")[1], 'base64').toString().split(/:(.*)/);
  return  [username, password]
}

export class AuthService {
  async create(userAuth: UserAuthParams): Promise<TokenInfo | AuthError> {
    const [username, password] = decodeAuth(userAuth)
    try {
      const newUser = await prisma.user.create({
        data: {
          username: username,
          password: bcrypt.hashSync(password),
          role: "USER"
        }
      })
      if (newUser) {
        return new AuthService().signIn(userAuth);
      }
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        if (e.code == "P2002") {
          return {
            message: "This User already exists",
            error: e.code,
          };
        } else {
          return {
            message: "Unknown error from the deep sea",
            error: e.code,
          };
        }
      }
    }
    return {message: "bad", error: "bad"}
  }

  async signIn(userAuth: UserAuthParams): Promise<TokenInfo | AuthError> {
    try {
      const [username, password] = decodeAuth(userAuth)
      const user = await prisma.user.findUnique({
        where: {
          username: username
        }
      });
      if (user) {
        const passwordIsValid = bcrypt.compareSync(password, user.password);

        if (!passwordIsValid) {
          return {
            message: "Password incorrect",
            error: "P2001",
          };
        }
        const tokenData = { username: user.username, role: user.role };
        const tokenOptions = { expiresIn: 86400 }
        const token = jwt.sign(tokenData, config.jwtSecret, tokenOptions);
        return {username: user.username, token: token}
      } else {
        return {
          message: "This User does not exist",
          error: "P2001",
        };
      }
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        if (e.code == "P2001") {
          return {
            message: "This User does not exist",
            error: e.code,
          };
        } else {
          return {
            message: "Unknown error from the deep sea",
            error: e.code,
          };
        }
      }
    }
    return {
      message: "Unknown error from the deep sea",
      error: "Spooky",
    };
  }
}
