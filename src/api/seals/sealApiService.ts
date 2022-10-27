import { Seals, Seal } from "./sealApiModel";
import { prisma } from "app";
import { Prisma } from "@prisma/client";
import crypto from "crypto";
import { config } from "config";
import path from "path";
import fs from "fs";

export type SealCreationParams = {
  file: Buffer;
  filename: string;
  tags: string[];
};

/**
 * Response when something goes wrong
 */
export interface SealError {
  /**
   * Message detailing the error, if the error is known
   */
  message: string;
  /**
   * Error code for the error. Usually a Prisma error
   */
  error: string;
}

/**
 * Generic response message
 */
export interface SealResponse {
  /**
   * Message providing information about the response
   */
  message: string;
}

export class SealApiService {
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

  async getRandomByTag(tag: string): Promise<Seal> {
    // Get seals with that tag
    const seals = await prisma.seal.findMany({
      where: {
        tags: {
          some: {
            name: tag,
          },
        },
      },
      orderBy: {
        slug: "desc",
      },
    });
    // Get a random one from the result
    const seal = seals[Math.floor(Math.random() * seals.length)];
    return seal as Seal;
  }

  async create(sealData: SealCreationParams): Promise<Seal | SealError> {
    const hasher = crypto.createHash("md5");
    const fileHash = hasher.update(sealData.file).digest("hex");
    const extension = path.extname(sealData.filename);
    const slug = `${fileHash}${extension}`;
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
      fs.writeFileSync(path.join(config.sealDir, slug), sealData.file);

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

  async update(id: number, tags: string[]): Promise<Seal> {
    const updatedSeal = await prisma.seal.update({
      where: {
        id: id,
      },
      data: {
        tags: {
          connectOrCreate: tags.map((tag) => ({
            where: {
              name: tag,
            },
            create: {
              name: tag,
            },
          })),
        },
      },
      include: {
        tags: true,
      },
    });
    return updatedSeal as Seal;
  }

  async deleteById(id: number): Promise<SealResponse> {
    const sealToDelete = await prisma.seal.findUnique({
      where: {
        id: id,
      },
    });

    if (sealToDelete) {
      await prisma.seal.delete({
        where: { id: sealToDelete.id },
      });
      return { message: `Deleted Seal with ID ${sealToDelete.id}` };
    }
    return { message: `Seal with ID ${id} not found` };
  }
}
