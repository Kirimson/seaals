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
    return seal as Seal
  }
}