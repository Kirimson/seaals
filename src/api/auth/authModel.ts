export interface UserAuthParams {
  /**
   * Name of the new user
   */
  basicAuth: string;
}

export interface TokenInfo {
  /**
   * Name of the new user
   */
  username: string;
  /**
   * A JWT for this user
   */
  token: string;
}

export interface AuthResponse {
  message: string
  success: boolean
}