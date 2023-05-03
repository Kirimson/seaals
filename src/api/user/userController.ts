import {
  Controller,
  Tags,
  Route,
  Path,
  Delete,
  Security,
  Get,
  Patch,
  Body,
} from "tsoa";
import { UserService } from "./userService";
import { User } from "@prisma/client";

@Route("/api/user")
@Tags("User API")
@Security("jwt", ["ADMIN"])
export class UserController extends Controller {
    /**
   * Delete an existing User by ID
   * @returns {UserResponse} Response if user is deleted
   */
    @Delete("/id/{id}")
    public async deleteByID(
      @Path() id: number
    ): Promise<any> {  
      return new UserService().deleteByID(id)
    }

    /**
   * Delete an existing User by Username
   * @returns {UserResponse} Response if user is deleted
   */
    @Delete("/username/{username}")
    public async deleteByName(
      @Path() username: string
    ): Promise<any> {  
      return new UserService().deleteByName(username)
    }

    /**
   * Get an Existing User by ID
   * @returns {UserResponse} Response if user is deleted
   */
    @Get("/id/{id}")
    public async getByID(
      @Path() id: number
    ): Promise<any> {  
      return new UserService().getByID(id)
    }

    /**
   * Get an Existing User by name
   * @returns {UserResponse} Response if user is deleted
   */
    @Get("/username/{username}")
    public async getByName(
      @Path() username: string
    ): Promise<any> {  
      return new UserService().getByName(username)
    }

    /**
   * Update an Existing User by ID
   * @returns {UserResponse} Response if user is deleted
   */
    @Patch("/id/{id}")
    public async patchByID(
      @Path() id: number,
      @Body() requestBody: Partial<Pick<User, "username" | "role">>
    ): Promise<any> {  
      return new UserService().updateUserByID(id, requestBody)
    }

    /**
   * Update an Existing User by ID
   * @returns {UserResponse} Response if user is deleted
   */
    @Patch("/username/{username}")
    public async patchByName(
      @Path() username: string,
      @Body() requestBody: Partial<Pick<User, "username" | "role">>
    ): Promise<any> {  
      return new UserService().updateUserByName(username, requestBody)
    }
}