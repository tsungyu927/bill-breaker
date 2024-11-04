import { APIResponse } from "@/interface/api";

export interface UserModel {
  id: string;
  device_id: string;
  name: string;
  email?: string;
  created_at: string;
  last_modified_at: string;
}

// Response
export type UserResponse = APIResponse<UserModel>;
export type CreateUserResponse = APIResponse<{ id: string }>;

// Payload
export interface CreateUserPayload {
  name: string;
  email?: string;
}
