import {
  Controller,
  Get,
  Path,
  Query,
  Route,
} from "tsoa";
import { Seal } from "./sealModel";
import { SealService } from "./sealService";

@Route("seals")
export class SealController extends Controller {
  /**
   * Get a Specific Seal
   * @param sealId Id of a seal
   * @returns {Seal} Data for a Seal
   */
  @Get("{sealId}")
  public async getSeal(
    @Path() sealId?: number,
  ): Promise<Seal> {
    return new SealService().get(sealId);
  }
}
