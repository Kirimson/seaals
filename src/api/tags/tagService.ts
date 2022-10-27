import { Tag } from "./tagModel";
import { prisma } from "app";
export type TagCreationParams = Pick<Tag, "name">;

export class TagService {
  async get(id?: number, name?: string, includeSeals?: boolean): Promise<Tag> {
    const tag = await prisma.tag.findFirst({
      where: {
        id: id,
        name: name,
      },
      include: {
        seals: includeSeals,
      },
    });
    return tag as Tag;
  }
}
