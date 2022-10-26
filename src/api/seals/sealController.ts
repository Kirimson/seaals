import { Controller, Get, Path, Query, Tags, Route } from "tsoa";
import { Seals, Seal } from "./sealModel";
import { SealService } from "./sealService";

@Route("/api/seals")
@Tags("Seals")
export class SealAPIController extends Controller {
  /**
   * Get a Specific Seal
   * @param id Id of a seal
   * @returns {Seals} Data for a Seal
   */
  @Get("/id/{id}")
  public async getSeal(@Path() id?: number): Promise<Seals> {
    return new SealService().getAll((id = id));
  }

  /**
   * Get all Seals
   * @param offset Offset to start at
   * @param limit Amount to seals to get. Defaults to 20
   * @returns {Seal[]} Data for a Seal
   */
  @Get("/")
  public async getAllSeals(
    @Query() offset = 0,
    @Query() limit = 20,
    @Query() id?: number,
    @Query() slug?: string,
    @Query() tags?: string[]
  ): Promise<Seals> {
    return new SealService().getAll(offset, limit, id, slug, tags);
  }

  /**
   * Get a Seal with tag
   * @param tag tag of a seal
   * @returns {Seal} Data for a Seal
   */
  @Get("/tag/{tag}")
  public async getSealByTag(@Path() tag: string): Promise<Seal> {
    return new SealService().getByTag(tag);
  }
}
