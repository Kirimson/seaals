import {
  Controller,
  Get,
  Post,
  Path,
  Query,
  Tags,
  Route,
  UploadedFile,
  FormField,
  Delete,
} from "tsoa";
import { Seals, Seal } from "./sealModel";
import {
  SealCreationParams,
  SealError,
  SealResponse,
  SealService,
} from "./sealService";
@Route("/api/seals")
@Tags("Seals")
export class SealAPIController extends Controller {
  /**
   * Get a Specific Seal
   * @param id Id of a seal
   * @returns {Seals} Data for a Seal
   */
  @Get("/id/{id}")
  public async getSeal(@Path() id: number): Promise<Seal> {
    return new SealService().getById(id);
  }

  /**
   * Delete a Seal
   * @param id Id of a seal to delete
   * @returns {SealResponse} Data for a Seal
   */
  @Delete("/id/{id}")
  public async deleteSeal(@Path() id: number): Promise<SealResponse> {
    return new SealService().deleteById(id);
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

  /**
   * Create a seal
   * @param tags comma delimited tags for the seal
   * @returns {Seal} Data for a Seal
   */
  @Post("/")
  public async createSeal(
    @FormField() tags: string,
    @UploadedFile() file: Express.Multer.File
  ): Promise<Seal | SealError> {
    const params: SealCreationParams = {
      file: file.buffer,
      tags: tags.split(",").map((tag) => tag.trim()),
    };
    return new SealService().create(params);
  }
}
