import {
  Book,
  BookDetail,
  CreateBookForm,
  JoinBookForm,
} from "@/interface/book";
import {
  BookDetailModel,
  BookModel,
  CreateBookPayload,
  JoinBookPayload,
} from "./models";

export const convertBookModelToBook = (bookModel: BookModel): Book => ({
  id: bookModel.id,
  name: bookModel.book_name,
  description: bookModel.book_description,
  creatorId: bookModel.creator_id,
  createAt: bookModel.create_at,
  lastModifiedAt: bookModel.last_modified_at,
});

export const convertBookDetailModelToBookDetail = (
  bookDetailModel: BookDetailModel
): BookDetail => ({
  id: bookDetailModel.id,
  name: bookDetailModel.book_name,
  description: bookDetailModel.book_description,
  members: bookDetailModel.members.map((item) => ({
    userId: item.user_id,
    userName: item.user_name,
  })),
  creatorId: bookDetailModel.creator_id,
  createAt: bookDetailModel.create_at,
  lastModifiedAt: bookDetailModel.last_modified_at,
});

export const convertToCreateBookPayload = (
  book: CreateBookForm
): CreateBookPayload => ({
  book_name: book.name,
  book_description: book.description,
});

export const convertToJoinBookPayload = (
  data: JoinBookForm
): JoinBookPayload => ({
  book_id: data.id,
});
