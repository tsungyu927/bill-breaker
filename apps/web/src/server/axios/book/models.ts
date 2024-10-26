import { APIResponse } from "@/src/interface/api";

export interface BookModel {
  ID: number;
  BookId: string;
  BookName: string;
  CreatedDate: number;
  LastModifiedDate: number;
}

export type BookListResponse = APIResponse<BookModel[]>;
export type BookResponse = APIResponse<BookModel>;
