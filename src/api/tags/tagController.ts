import { Controller, Get, Path, Query, Route, Tags } from "tsoa";
import { Tag } from "./tagModel";
import { TagService } from "./tagService";

@Route("/api/tags")
@Tags("Tags")
export class TagController extends Controller {
  /**
   * Get a Specific Tag
   * @param id Id of a tag
   * @returns {Tag} Data for a Tag
   */
  @Get("{id}")
  public async getTag(@Path() id?: number): Promise<Tag> {
    return new TagService().get(id);
  }
}
