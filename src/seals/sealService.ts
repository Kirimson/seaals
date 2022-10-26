import { Seal } from "seals/sealModel";
import { prisma } from "app";
export type SealCreationParams = Pick<Seal, "slug" | "tags">

export class SealService {
  async get(id?: number, slug?: string, tags?: string[]): Promise<Seal> {
    const seal = await prisma.seal.findFirst({
      where: {
        id: id,
        slug: slug,
        tags: {
          some: {
            name: {
              in: tags
            }
          }
        }
      },
      include: {
        tags: true
      }
    })
    if (!seal) return {} as Seal;
    return seal as Seal
  }

  async getByTag(tag:string) {
    return this.get(undefined, undefined, [tag])
  }
  async getByTags(tag:string[]) {
    return this.get(undefined, undefined, tag)
  }

  async getRandom(): Promise<Seal> {
    const sealCount = await prisma.seal.count()
    const skip = Math.floor(Math.random() * sealCount);
    const seal = await prisma.seal.findFirst({
      take: 1,
      skip: skip,
      orderBy: {
        slug: 'desc'
      }
    })
    return seal as Seal
  }
}