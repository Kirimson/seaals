import { SealApiService } from "api/seals/sealApiService";
import { Router } from "express";
import { config } from "config";
import path from "path";
export const sealRouter = Router();

/**
 * Get a random Seal
 */
sealRouter.get("/", async (req, res, next) => {
  const seal = await new SealApiService().getRandom();
  const sealPath = path.join(config.sealDir, seal.slug);
  res.sendFile(sealPath);
});

/**
 * Get a random Seal
 */
sealRouter.get("/:tag", async (req, res, next) => {
  const seal = await new SealApiService().getByTag(req.params.tag);
  const sealPath = path.join(config.sealDir, seal.slug);
  res.sendFile(sealPath);
});
