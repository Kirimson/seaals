import { Controller, Get, Path, Query, Route, Tags } from "tsoa";
import { Tag } from "./tagModel";
import { TagService } from "./tagService";

@Route("/api/tags")
@Tags("Tags")
export class TagController extends Controller {
  /**
   * Get a Specific Tag by ID
   * @param id Id of a tag
   * @param includeSeals Boolean to include associated Seals
   * @returns {Tag} Data for a Tag
   */
  @Get("/id/{id}")
  public async getTag(
    @Path() id?: number,
    @Query() includeSeals?: boolean
  ): Promise<Tag> {
    return new TagService().get(id, undefined, includeSeals);
  }

  /**
   * Get a Specific Tag by ID
   * @param name Name of a tag
   * @param includeSeals Boolean to include associated Seals
   * @returns {Tag} Data for a Tag
   */
  @Get("/tag/{name}")
  public async getTagByName(
    @Path() name?: string,
    @Query() includeSeals?: boolean
  ): Promise<Tag> {
    return new TagService().get(undefined, name, includeSeals);
  }
}
