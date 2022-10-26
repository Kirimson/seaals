import { Controller, Get, Path, Tags, Route } from "tsoa";
import { Seal } from "./sealModel";
import { SealService } from "./sealService";

@Route("seal")
@Tags("Seal")
export class SealController extends Controller {
  /**
   * Get a Specific Seal
   * @param id Id of a seal
   * @returns {Seal} Data for a Seal
   */
  @Get("/id/{id}")
  public async getSeal(@Path() id?: number): Promise<Seal> {
    return new SealService().get(id);
  }

  /**
   * Get a random Seal
   * @returns {Seal} Data for a Seal
   */
  @Get("/")
  public async getRandomSeal(): Promise<Seal> {
    return new SealService().getRandom();
  }

  /**
   * Get a Seal with tag
   * @param tag tag of a seal
   * @returns {Seal} Data for a Seal
   */
   @Get("/{tag}")
   public async getSealByTag(@Path() tag: string): Promise<Seal> {
     return new SealService().getByTag(tag);
   }
}
