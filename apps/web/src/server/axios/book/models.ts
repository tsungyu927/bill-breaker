import { APIResponse } from "@/interface/api";

export interface BookModel {
  id: string;
  creator_id: string;
  book_name: string;
  book_description?: string;
  create_at: string;
  last_modified_at: string;
}

// Response
export interface BookDetailModel extends BookModel {
  members: {
    user_id: string;
    user_name: string;
  }[];
}

export type CreateBookResponse = APIResponse<BookModel>;
export type BookListResponse = APIResponse<BookModel[]>;
export type BookResponse = APIResponse<BookDetailModel>;
export type JoinBookResponse = APIResponse<string>;

// Payload
export type CreateBookPayload = Pick<
  BookModel,
  "book_name" | "book_description"
>;
export type JoinBookPayload = {
  book_id: string;
};
