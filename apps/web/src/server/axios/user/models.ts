import { APIResponse } from "@/interface/api";

export interface UserModel {
  id: string;
  device_id: string;
  name: string;
  email?: string;
  created_at: string;
  last_modified_at: string;
}

export type UserResponse = APIResponse<UserModel>;
