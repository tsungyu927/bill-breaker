import axiosInstance from "..";
import { BookListResponse, BookResponse } from "./models";

export const fetchBookList = async () => {
  const { data } = await axiosInstance.get<BookListResponse>("api/v1/books");

  return data.data;
};

export const fetchBookByBookId = async (bookId: string) => {
  const { data } = await axiosInstance.get<BookResponse>(
    `api/v1/book/${bookId}`
  );

  return data.data;
};
