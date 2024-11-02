import axiosInstance from "..";
import { UserResponse } from "./models";

export const fetchUser = async (userId: string) => {
  const { data } = await axiosInstance.get<UserResponse>(
    `api/v1/user/${userId}`
  );

  return data.data;
};
