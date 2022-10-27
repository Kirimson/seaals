import express, { json, urlencoded } from "express";
import { RegisterRoutes } from "api/routes";
import { PrismaClient } from "@prisma/client";
import path from "path";

export const app = express();
export const prisma = new PrismaClient();
// Use body parser to read sent json payloads
app.use(
  urlencoded({
    extended: true,
  })
);
app.use(json());

app.set("views", path.join(__dirname, "views"));
app.set("view engine", "pug");
RegisterRoutes(app);
