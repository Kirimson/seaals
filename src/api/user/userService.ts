import { Prisma, User } from "@prisma/client";
import { UserInfo, UserResponse } from "./userModel";
import { prisma } from "app";

function exclude<User, Key extends keyof User>(
  user: User,
  keys: Key[]
): Omit<User, Key> {
  for (let key of keys) {
    delete user[key]
  }
  return user
}

export class UserService {
  async deleteByID(userID: number): Promise<UserResponse> {
    const user = await prisma.user.findUnique({where: {id: userID}});
    if (user) {
      return this.deleteUser(user);
    } else return {message: "User does not exist", success: false}
  }

  async deleteByName(username: string): Promise<UserResponse> {
    const user = await prisma.user.findUnique({where: {username: username}});
    if (user) {
      return this.deleteUser(user);
    } else return {message: "User does not exist", success: false}
  }
  
  async getByID(userID:number): Promise<UserInfo|UserResponse> {
    const user = await prisma.user.findUnique({where: {id: userID}});
    return this.sanitizeUser(user);
  }

  async getByName(username:string): Promise<UserInfo|UserResponse> {
    const user = await prisma.user.findUnique({where: {username: username}});
    return this.sanitizeUser(user);
  }

  async updateUserByID(userID:number,data:Partial<Pick<User, "username" | "role">>): Promise<UserInfo|UserResponse> {
    return this.updateUser({id: userID}, data)
  }


  async updateUserByName(username:string, data:Partial<Pick<User, "username" | "role">>): Promise<UserInfo|UserResponse> {
    return this.updateUser({username: username}, data)
  }

  protected async updateUser(select:{username: string}|{id: number}, data:Partial<Pick<User, "username" | "role">>): Promise<UserInfo|UserResponse>  {
    // Check If user exists
    const user = await prisma.user.findUnique({where: select});
    if (user) {
      try {
        const updateUser = await prisma.user.update({
          where: select,
          data: data,
        })
        if (updateUser) {
          const userWithoutPassword = exclude(updateUser, ['password'])
          return {user: userWithoutPassword, success: true}
        }
      } catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
          if (e.code == "P2002") {
            return {
              message: "User with this name already exists",
              success: false,
            };
          } else {
            return {
              message: "Unknown error from the deep sea",
              success: false,
            };
          }
        }
      }

    }
    return {message: "User does not exist", success: false}
  }

  protected sanitizeUser(user:User|null): UserInfo|UserResponse {
    if (user) {
      const userWithoutPassword = exclude(user, ['password'])
      return {user: userWithoutPassword, success: true}
    }
    return {message: "User does not exist", success: false}
  }

  protected async deleteUser(user:User): Promise<UserResponse>{
    const deleted = await prisma.user.delete({where: {id: user.id}});
    if (deleted) {
      return {message: `${deleted.username} Deleted`, success: true}
    } else {
      return {message: `Failed to Delete ${user.username}`, success: false}
    }
  }
}
