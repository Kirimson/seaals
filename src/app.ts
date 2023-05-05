import express, { NextFunction, Request, Response, json, urlencoded } from "express";
import { RegisterRoutes } from "api/routes";
import { PrismaClient } from "@prisma/client";
import path from "path";
import { ValidateError } from "tsoa";
import cors from "cors";

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
app.use(cors());

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