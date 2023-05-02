import express, { NextFunction, Request, Response, json, urlencoded } from "express";
import { RegisterRoutes } from "api/routes";
import { PrismaClient } from "@prisma/client";
import path from "path";
import { ValidateError } from "tsoa";

import jwt from "jsonwebtoken";

export const app = express();
export const prisma = new PrismaClient();
// Use body parser to read sent json payloads
app.use(
  urlencoded({
    extended: true,
  })
);
app.use(json());

interface IUser {
  username: string
}

function generateAccessToken(username:IUser) {
  if (process.env.TOKEN_SECRET) {
    return jwt.sign(username, process.env.TOKEN_SECRET, { expiresIn: '1800s' });
  }
  return ""
}

app.post('/api/createNewUser', (req, res) => {
  const token = generateAccessToken({ username: req.body.username });
  res.json(token);
});


app.set("views", path.join(__dirname, "views"));
app.set("view engine", "pug");
RegisterRoutes(app);

app.use(function errorHandler(
  err: unknown,
  req: Request,
  res: Response,
  next: NextFunction
): Response | void {
  if (err instanceof ValidateError) {
    console.warn(`Caught Validation Error for ${req.path}:`, err.fields);
    return res.status(422).json({
      message: "Validation Failed",
      details: err?.fields,
    });
  }
  if (err instanceof Error) {
    return res.status(500).json({
      message: err.message,
    });
  }

  next();
});