import {
  Controller,
  Tags,
  Route,
  Path,
  Delete,
  Security,
} from "tsoa";
import { UserService } from "./userService";
import { UserResponse } from "./userModel";

@Route("/api/user")
@Tags("Auth API")
export class UserController extends Controller {
    /**
   * Delete an existing User by ID
   * @returns {UserResponse} Response if user is deleted
   */
    @Delete("/id/{id}")
    @Security("jwt", ["ADMIN"])
    public async deleteByID(
      @Path() id: number
    ): Promise<any> {  
      return new UserService().deleteByID(id)
    }

    /**
   * Delete an existing User by IName
   * @returns {UserResponse} Response if user is deleted
   */
    @Delete("/username/{username}")
    @Security("jwt", ["ADMIN"])
    public async deleteByName(
      @Path() username: string
    ): Promise<any> {  
      return new UserService().deleteByName(username)
    }
}