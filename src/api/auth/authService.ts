import { Prisma } from "@prisma/client";
import bcrypt from 'bcryptjs';
import jwt from "jsonwebtoken";

import { prisma } from "app";
import { UserAuthParams, AuthResponse } from "./authModel";

export class AuthService {
  async create(userAuth: UserAuthParams): Promise<AuthResponse> {
    const [username, password] = Buffer.from(userAuth.basicAuth.split(" ")[1], 'base64').toString().split(/:(.*)/);
    try {
      const newUser = await prisma.user.create({
        data: {
          username: username,
          password: bcrypt.hashSync(password),
          role: "USER"
        }
      })
      if (newUser) {
        return {message: "User Created", success: true};
      }
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        if (e.code == "P2002") {
          return {
            message: "This User already exists",
            success: false,
          };
        } else {
          return {
            message: "Unknown error from the deep sea",
            success: false,
          };
        }
      }
    }
    return {message: "bad", success: false}
  }
}
