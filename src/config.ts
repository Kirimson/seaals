import * as dotenv from "dotenv";
import path from "path";
dotenv.config();

const defaultMimes = [
  "image/png",
  "image/jpeg",
  "image/jpg",
  "image/webp",
  "image/gif",
];

export const config = {
  sealDir: path.resolve(__dirname, process.env.SEAL_DIR || "resources/seals"),
  validMimes:
    process.env.VALID_MIMES?.split(",").map((mime) => mime.trim()) ||
    defaultMimes,
  jwtSecret: process.env.VALID_MIMES || "pleasechangeme"
};
