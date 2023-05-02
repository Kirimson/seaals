import * as express from "express";
import * as jwt from "jsonwebtoken";
import { config } from "config";
import { prisma } from "app";
import { TokenExpiredError } from "jsonwebtoken";

type SealToken = {
  username: string
  iat: number
  exp: number
}

export async function expressAuthentication(
  request: express.Request,
  securityName: string,
  roles?: string[]
): Promise<any> {
  if (securityName != "jwt") return new Promise((res, rej) => rej(new Error("Only jwt authentication is allowed!")));
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
        const user = await prisma.user.findUnique({
          where: {
            username: decoded.username
          }
        })
        // If user is not part of one of the required roles, reject
        if (!roles.includes(user?.role||"")) reject(new Error("Incorrect Role"));
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
}