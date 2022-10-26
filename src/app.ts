import express, { json, urlencoded } from "express";
import { RegisterRoutes } from "api/routes";
import { PrismaClient } from "@prisma/client";
import { sealRouter } from "seal/sealRouter";

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
app.use("/seal", sealRouter);
