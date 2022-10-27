import { Controller, Request, Get, Route, Tags, Path } from "tsoa";
import express from "express";
import { SealService } from "./sealService";

@Route("/seal")
@Tags("Seal")
export class TestController extends Controller {
  /**
   * Get a random Seal image
   * @returns Image of a seal
   */
  @Get("/")
  public async getRandomSeal(@Request() request: express.Request) {
    // Some silly any casting so we can make ourselves a express.Response instance
    let res = (<any>request).res as express.Response;
    return new SealService().getRandom(res);
  }

  /**
   * Get a random Seal image with a set tag
   * @returns Image of a seal
   */
  @Get("/{tag}")
  public async getRandomSealWithTag(
    @Request() request: express.Request,
    @Path() tag: string
  ) {
    // Some silly any casting so we can make ourselves a express.Response instance
    let res = (<any>request).res as express.Response;
    return new SealService().getRandomByTag(res, tag);
  }
}
