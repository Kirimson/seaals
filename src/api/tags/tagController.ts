import { SealResponse } from "../seals/sealApiModel";
import {
  Controller,
  Delete,
  Get,
  Path,
  Query,
  Route,
  Tags as TsoaTags,
} from "tsoa";
import { Tag, Tags } from "./tagModel";
import { TagService } from "./tagService";

@Route("/api/tags")
// Named TsoaTags here due to name conflict with the Tags Model. Renaming the Tags model causes tsoa to fail
@TsoaTags("Tags")
export class TagController extends Controller {
  /**
   * Get all Tags
   * @param offset Offset to start at
   * @param limit Amount to seals to get. Defaults to 20
   * @param limit Include associated Seals with response
   * @returns {Seal[]} Data for a Seal
   */
  @Get("/")
  public async getAllSeals(
    @Query() offset = 0,
    @Query() limit = 20,
    @Query() includeSeals?: boolean
  ): Promise<Tags> {
    return new TagService().getAll(offset, limit, includeSeals);
  }

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

  /**
   * Delete a Tag by ID
   * @param id Id of a tag to delete
   * @returns {SealResponse} Deletion confirmation
   */
  @Delete("/id/{id}")
  public async deleteTag(@Path() id: number): Promise<SealResponse> {
    return new TagService().deleteById(id);
  }

  /**
   * Delete a Tag by Name
   * @param name Id of a tag to delete
   * @returns {SealResponse} Deletion confirmation
   */
  @Delete("/name/{name}")
  public async deleteTagByName(@Path() name: string): Promise<SealResponse> {
    return new TagService().deleteByName(name);
  }
}
