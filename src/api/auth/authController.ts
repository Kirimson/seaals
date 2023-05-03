import {
  Controller,
  Post,
  Tags,
  Route,
  Header,
  Security,
  Request,
} from "tsoa";

import { TokenInfo, AuthResponse } from "./authModel";
import { AuthService } from "./authService";

@Route("/api/auth")
@Tags("Auth API")
export class AuthController extends Controller {
    /**
   * Create a new User
   * @returns {TokenInfo} Authentication for the user
   */
    @Post("/register")
    public async signUp(
      @Header('authorization') basicAuth: string,
    ): Promise<AuthResponse> {  
      return new AuthService().create({basicAuth: basicAuth})
    }

    /**
   * Request a token for an existing User
   * @returns {TokenInfo} Authentication for the user
   */
    @Post("/token")
    @Security("basic")
    public async signIn(
      @Request() request: any
    ): Promise<TokenInfo | AuthResponse> {
      return Promise.resolve(request.user)
    }
}