import { BaseRoom, SuccessResponse } from "../@types/api";
import { protectedApiClient } from "../utils/apiClient";

export interface UserResponse {
  id: string;
  username: string;
  password: string;
  email: string;
  rooms: BaseRoom[];
}

export const getUser = () => {
  return protectedApiClient.get<SuccessResponse<UserResponse>>("/user/profile");
};
