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
  Put,
} from "tsoa";
import {
  ManySeals,
  Seal,
  SealCreationParams,
  SealError,
  SealResponse,
} from "./sealApiModel";
import { SealApiService } from "./sealApiService";
@Route("/api/seals")
@Tags("Seals API")
export class SealAPIController extends Controller {
  /**
   * Get a Specific Seal
   * @param id Id of a seal
   * @returns {ManySeals} Data for a Seal
   */
  @Get("/id/{id}")
  public async getSeal(@Path() id: number): Promise<Seal> {
    return new SealApiService().getById(id);
  }

  /**
   * Delete a Seal
   * @param id Id of a seal to delete
   * @returns {SealResponse} Deletion confirmation
   */
  @Delete("/id/{id}")
  public async deleteSeal(@Path() id: number): Promise<SealResponse> {
    return new SealApiService().deleteById(id);
  }

  /**
   * Get all Seals
   * @param offset Offset to start at
   * @param limit Amount to seals to get. Defaults to 20
   * @param id ID of a seal to get
   * @param slug Slug to identify seal with
   * @param tags Tags to search for
   * @returns {Seal[]} Data for a Seal
   */
  @Get("/")
  public async getAllSeals(
    @Query() offset = 0,
    @Query() limit = 20,
    @Query() id?: number,
    @Query() slug?: string,
    @Query() tags?: string[]
  ): Promise<ManySeals> {
    return new SealApiService().getAll(offset, limit, id, slug, tags);
  }

  /**
   * Create a new seal
   * @param tags comma delimited tags for the seal
   * @returns {Seal} Data for a Seal
   */
  @Post("/")
  public async createSeal(
    @FormField() tags: string,
    @UploadedFile() file: Express.Multer.File
  ): Promise<Seal | SealError> {
    const params: SealCreationParams = {
      filename: file.originalname,
      file: file.buffer,
      tags: tags.split(",").map((tag) => tag.trim()),
    };
    return new SealApiService().create(params);
  }

  /**
   * Updates a Seal's tags
   */
  @Put("/{id}")
  public async updateSeal(@Path() id: number, @Query() tags: string[]) {
    return new SealApiService().update(id, tags);
  }
}
