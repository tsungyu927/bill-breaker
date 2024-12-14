import axiosInstance from "..";
import {
  convertBookDetailModelToBookDetail,
  convertBookModelToBook,
} from "./converter";
import {
  BookListResponse,
  BookResponse,
  CreateBookPayload,
  CreateBookResponse,
  CreateCostPayload,
  CreateCostResponse,
  JoinBookPayload,
  JoinBookResponse,
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

export const fetchBookByBookId = async (bookId: string) => {
  const { data } = await axiosInstance.get<BookResponse>(
    `api/v1/book/${bookId}`
  );

  if (data.data) {
    return convertBookDetailModelToBookDetail(data.data);
  }
  return null;
};

export const joinBook = async (payload: JoinBookPayload) => {
  return axiosInstance.post<JoinBookResponse>(`api/v1/book/join`, payload);
};

// Cost
export const createCost = async (payload: CreateCostPayload) => {
  const { data } = await axiosInstance.post<CreateCostResponse>(
    `api/v1/book/cost`,
    payload
  );

  return data.data;
};
