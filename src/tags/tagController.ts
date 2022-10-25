import {
  Controller,
  Get,
  Path,
  Query,
  Route,
} from "tsoa";
import { Tag } from "./tagModel";
import { TagService } from "./tagService";

@Route("tags")
export class TagController extends Controller {
  @Get("{tagId}")
  public async getTag(
    @Path() tagId?: number,
  ): Promise<Tag> {
    return new TagService().get(tagId);
  }
}
