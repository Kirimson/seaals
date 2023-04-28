import * as express from "express";
import * as jwt from "jsonwebtoken";

export function expressAuthentication(
  request: express.Request,
  securityName: string,
  scopes?: string[]
): Promise<any> {
  if (securityName === "jwt") {
    const token =
      request.body.token ||
      request.query.token ||
      request.headers["authorization"]?.split("Bearer ")[1];
    
    return new Promise((resolve, reject) => {
      if (!token) {
        reject(new Error("No token provided"));
      }
      if (process.env.TOKEN_SECRET) {
        jwt.verify(token, process.env.TOKEN_SECRET, function (err: any, decoded: any) {
          if (err) {
            reject(err);
          } else {
            // Check if JWT contains all required scopes
            if (scopes) {
              for (let scope of scopes) {
                if (!decoded.scopes.includes(scope)) {
                  reject({"status": 403, "msg": "unauthorised"});
                }
              }
              resolve(decoded);
            }
          }
        });
      }
    });
  }
  return new Promise((res, rej) => rej("No valid security name provided"));
}