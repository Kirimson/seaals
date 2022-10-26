import * as dotenv from "dotenv";
import path from "path";
dotenv.config();

export const config = {
  sealDir: path.resolve(__dirname, process.env.SEAL_DIR || "resources/seals"),
};
