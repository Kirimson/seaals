import * as express from "express";
import * as jwt from "jsonwebtoken";
import { config } from "config";
import { prisma } from "app";
import { TokenExpiredError } from "jsonwebtoken";
import bcrypt from 'bcryptjs';
import { TokenInfo } from "api/auth/authModel";

type SealToken = {
  username: string
  role: string
  iat: number
  exp: number
}

export async function expressAuthentication(
  request: express.Request,
  securityName: string,
  roles?: string[]
): Promise<any> {
  if (securityName == "jwt") {
    // Try and grab a token from some places
    const token =
      request.body.token ||
      request.query.token ||
      request.headers["authorization"]?.split("Bearer ")[1];
      
    return new Promise(async (resolve, reject) => {
      // If no token was found, return right away
      if (!token) reject(new Error("No token provided!"));
      try {
        // Try and decode the JWT with the secret from config
        const decoded = jwt.verify(token, config.jwtSecret) as SealToken;
        
        // If a role is needed, go and check it. Else just resolve
        if (roles) {
          // If user is not part of one of the required roles, reject
          if (!roles.includes(decoded.role||"")) reject(new Error("Incorrect Role"));
        }
        resolve(decoded)
      // If decoding fails
      } catch (e) {
        // If token has expired, it will fail
        if (e instanceof TokenExpiredError) {
          reject(new Error(e.message));
        // Some other token error
        } else if (e instanceof jwt.JsonWebTokenError) {
          reject(new Error(e.message));
        // Something else entirely
        } else {
          reject(new Error("Bad Token"));
        }
      }
    })
  } else if (securityName == "basic") {
    // Get Basic Auth header
    const basicAuth = request.headers["authorization"]?.split("Basic ")[1];
    
    return new Promise(async (resolve, reject) => {
      if (!basicAuth) return reject(new Error("Invalid Auth Header"));
      // Decode basic auth
      const [username, password] = Buffer.from(basicAuth, 'base64').toString().split(/:(.*)/);

      // Get the user for the username provided
      const user = await prisma.user.findUnique({
        where: {
          username: username
        }
      });

      // If that user exists, check the password
      if (user) {
        const passwordIsValid = bcrypt.compareSync(password, user.password);

        if (!passwordIsValid) {
          reject(new Error("Password Incorrect"))
        }

        const tokenData = { username: user.username, role: user.role };
        const tokenOptions = { expiresIn: 86400 }
        const token = jwt.sign(tokenData, config.jwtSecret, tokenOptions);
        const userinfo: TokenInfo = {
          username: username,
          token: token
        }
        resolve(userinfo);
      } else {
        reject(new Error("This User does not exist"))
      }
    });
  } else {
    return new Promise((res, rej) => rej(new Error("Invalid Security Type")));
  }
}