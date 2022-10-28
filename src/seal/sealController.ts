import { Controller, Request, Get, Route, Tags, Path, Query } from "tsoa";
import express from "express";
import { SealService } from "./sealService";

@Route("/seal")
@Tags("Seals")
export class SealController extends Controller {
  /**
   * Get a random Seal image
   * @param html Render the seal in a HTML document
   * @returns Image of a seal
   */
  @Get("/")
  public async getRandomSeal(
    @Request() request: express.Request,
    @Query() html?: boolean
  ) {
    // Some silly any casting so we can make ourselves a express.Response instance
    let res = (<any>request).res as express.Response;
    return new SealService().getRandom(res, html);
  }

  /**
   * Get a random Seal image with a set tag
   * @param tag Type of seal to get
   * @param html Render the seal in a HTML document
   * @returns Image of a seal
   */
  @Get("/tag/{tag}")
  public async getRandomSealByTag(
    @Request() request: express.Request,
    @Path() tag: string,
    @Query() html?: boolean
  ) {
    // Some silly any casting so we can make ourselves a express.Response instance
    let res = (<any>request).res as express.Response;
    return new SealService().getRandomByTag(res, tag, html);
  }

  /**
   * Get a Seal with a specific ID
   * @param id ID of seal to get
   * @returns Image of a seal
   */
  @Get("/id/{id}")
  public async getSealById(
    @Request() request: express.Request,
    @Path() id: number
  ) {
    // Some silly any casting so we can make ourselves a express.Response instance
    let res = (<any>request).res as express.Response;
    return new SealService().getById(res, id);
  }
}
