import { SealApiService } from "../api/seals/sealApiService";
import { config } from "config";
import path from "path";
import fs from "fs";
import mime from "mime-types";
import express from "express";
import { Seal, SealResponse } from "api/seals/sealApiModel";

export class SealService {
  async getRandom(res: express.Response) {
    // Get a random Seal model using the underlying Seal API and get the path for it
    const seal = await new SealApiService().getRandom();
    this.handleSeal(res, seal);
  }

  async getRandomByTag(res: express.Response, tag: string) {
    // Get a random Seal model using the underlying Seal API and get the path for it
    const seal = await new SealApiService().getRandomByTag(tag);
    this.handleSeal(res, seal);
  }

  private handleSeal(res: express.Response, seal: Seal) {
    if (seal) {
      const filePath = path.join(config.sealDir, seal.slug);

      // Load the seal into memory
      const data = fs.readFileSync(filePath);
      // Get the Content Type of the image to set in the response
      const contType = mime.contentType(filePath);
      // So TS won't complain about potential null/false/undefined vars
      if (data && contType) {
        // Set content type, and send the Buffer of the image
        res.contentType(contType);
        res.end(data);
      }
    } else {
      res.json({ message: "No Seal found" } as SealResponse);
    }
  }
}
