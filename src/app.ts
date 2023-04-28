import express, { json, urlencoded } from "express";
import { RegisterRoutes } from "api/routes";
import { PrismaClient } from "@prisma/client";
import path from "path";

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
