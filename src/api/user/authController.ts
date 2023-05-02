import {
  Controller,
  Post,
  Tags,
  Route,
  Header
} from "tsoa";

import { TokenInfo, AuthError } from "./authModel";
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
      @Header('authorization') basicAuth: string
    ): Promise<TokenInfo | AuthError> {  
      return new AuthService().create({basicAuth: basicAuth})
    }

    /**
   * Request a token for an existing User
   * @returns {TokenInfo} Authentication for the user
   */
    @Post("/token")
    public async signIn(
      @Header('authorization') basicAuth: string
    ): Promise<TokenInfo | AuthError> {  
      return new AuthService().signIn({basicAuth: basicAuth})
    }
}