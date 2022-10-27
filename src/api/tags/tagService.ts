import { Tag, Tags } from "./tagModel";
import { prisma } from "app";
import { SealResponse } from "../seals/sealApiService";
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

  async getAll(offset = 0, limit = 20, includeSeals?: boolean): Promise<Tags> {
    const tags = await prisma.tag.findMany({
      take: limit,
      skip: offset,
      include: {
        seals: includeSeals,
      },
    });

    const allTags = {
      count: tags.length,
      offset: offset,
      limit: limit,
      tags: tags as Tag[],
    } as Tags;

    return allTags;
  }

  async deleteById(id: number): Promise<SealResponse> {
    const tagToDelete = await prisma.tag.findUnique({
      where: {
        id: id,
      },
    });

    if (tagToDelete) {
      await prisma.tag.delete({
        where: { id: tagToDelete.id },
      });
      return { message: `Deleted Tag with ID ${tagToDelete.id}` };
    }
    return { message: `Tag with ID ${id} not found` };
  }

  async deleteByName(name: string): Promise<SealResponse> {
    const tagToDelete = await prisma.tag.findUnique({
      where: {
        name: name,
      },
    });

    if (tagToDelete) {
      await prisma.tag.delete({
        where: { id: tagToDelete.id },
      });
      return { message: `Deleted Tag with Name ${tagToDelete.id}` };
    }
    return { message: `Tag with name ${name} not found` };
  }
}
