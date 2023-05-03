import { User } from "@prisma/client"

export interface UserResponse {
  message: string,
  success: boolean
}

export interface UserInfo {
  user: Pick<User, "id" | "username" | "role">|null,
  success: boolean
}