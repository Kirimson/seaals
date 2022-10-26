import express, { json, urlencoded } from "express";
import { RegisterRoutes } from "routes";
import { PrismaClient } from "@prisma/client";

export const app = express();
export const prisma = new PrismaClient();
// Use body parser to read sent json payloads
app.use(
  urlencoded({
    extended: true,
  })
);
app.use(json());

RegisterRoutes(app);
