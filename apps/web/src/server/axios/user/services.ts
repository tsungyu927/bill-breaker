import axiosInstance from "..";
import { CreateUserPayload, CreateUserResponse, UserResponse } from "./models";

export const fetchUser = async (userId: string) => {
  const { data } = await axiosInstance.get<UserResponse>(
    `api/v1/user/${userId}`
  );

  return data.data;
};

export const createUser = async (payload: CreateUserPayload) => {
  const { data } = await axiosInstance.post<CreateUserResponse>(
    `api/v1/user`,
    payload
  );

  return data.data;
};
