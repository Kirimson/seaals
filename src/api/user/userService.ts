import { User } from "@prisma/client";
import { UserResponse } from "./userModel";
import { prisma } from "app";

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

  protected async deleteUser(user:User): Promise<UserResponse>{
    const deleted = await prisma.user.delete({where: {id: user.id}});
    if (deleted) {
      return {message: `${deleted.username} Deleted`, success: true}
    } else {
      return {message: `Failed to Delete ${user.username}`, success: false}
    }
  }
}
