export interface APIResponse<T> {
  status: "success" | "error";
  data: T | null;
  message: string | null;
}
