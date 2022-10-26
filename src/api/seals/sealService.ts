import { Seals, Seal } from "./sealModel";
import { prisma } from "app";
import { Prisma } from "@prisma/client";
import crypto from "crypto";
import { config } from "config";
import fs from "fs";

export type SealCreationParams = {
  file: Buffer;
  tags: string[];
};

export interface SealError {
  message: string;
  error: string;
}

export class SealService {
  async getById(id: number): Promise<Seal> {
    return this.get(id, undefined, undefined);
  }

  async get(id?: number, slug?: string, tags?: string[]): Promise<Seal> {
    const seal = await prisma.seal.findFirst({
      where: {
        id: id,
        slug: slug,
        tags: {
          some: {
            name: {
              in: tags,
            },
          },
        },
      },
      include: {
        tags: true,
      },
    });
    if (!seal) return {} as Seal;
    return seal as Seal;
  }

  async getByTag(tag: string) {
    return this.get(undefined, undefined, [tag]);
  }
  async getByTags(tag: string[]) {
    return this.get(undefined, undefined, tag);
  }

  async getAll(
    offset = 0,
    limit = 20,
    id?: number,
    slug?: string,
    tags?: string[]
  ): Promise<Seals> {
    const seals = await prisma.seal.findMany({
      take: limit,
      skip: offset,
      where: {
        id: id,
        slug: slug,
        tags: {
          some: {
            name: {
              in: tags,
            },
          },
        },
      },
      include: {
        tags: true,
      },
    });

    const allSeals = {
      count: seals.length,
      offset: offset,
      limit: limit,
      seals: seals as Seal[],
    } as Seals;

    return allSeals;
  }

  async getRandom(): Promise<Seal> {
    const sealCount = await prisma.seal.count();
    const skip = Math.floor(Math.random() * sealCount);
    const seal = await prisma.seal.findFirst({
      take: 1,
      skip: skip,
      orderBy: {
        slug: "desc",
      },
    });
    return seal as Seal;
  }

  async create(sealData: SealCreationParams): Promise<Seal | SealError> {
    const hasher = crypto.createHash("md5");
    const slug = hasher.update(sealData.file).digest("hex");
    try {
      // Create the seal, and all the tags for it as well
      const newSeal = await prisma.seal.create({
        data: {
          slug: slug,
          tags: {
            connectOrCreate: sealData.tags.map((tag) => ({
              where: {
                name: tag,
              },
              create: {
                name: tag,
              },
            })),
          },
        },
      });
      // Save the seal to file from the sealData
      fs.writeFileSync(config.sealDir, sealData.file);

      return newSeal as Seal;
    } catch (e) {
      if (e instanceof Prisma.PrismaClientKnownRequestError) {
        if (e.code == "P2002") {
          return {
            message: "This Seal already exists",
            error: e.code,
          };
        } else {
          return {
            message: "Unknown error from the deep sea",
            error: e.code,
          };
        }
      }
    }

    return {} as SealError;
  }
}
