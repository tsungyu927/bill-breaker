import axiosInstance from "..";
import { convertBookModelToBook } from "./converter";
import {
  BookListResponse,
  CreateBookPayload,
  CreateBookResponse,
} from "./models";

export const createBook = async (payload: CreateBookPayload) => {
  const { data } = await axiosInstance.post<CreateBookResponse>(
    `api/v1/book`,
    payload
  );

  if (data.data) {
    return convertBookModelToBook(data.data);
  }
  return null;
};

export const fetchBookList = async () => {
  const { data } = await axiosInstance.get<BookListResponse>(
    "api/v1/book/list"
  );

  return data.data?.map(convertBookModelToBook) || [];
};

// export const fetchBookByBookId = async (bookId: string) => {
//   const { data } = await axiosInstance.get<BookResponse>(
//     `api/v1/book/${bookId}`
//   );

//   return data.data;
// };
