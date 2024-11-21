import { Book, CreateBookForm, JoinBookForm } from "@/interface/book";
import { BookModel, CreateBookPayload, JoinBookPayload } from "./models";

export const convertBookModelToBook = (bookModel: BookModel): Book => ({
  id: bookModel.id,
  name: bookModel.book_name,
  description: bookModel.book_description,
  creatorId: bookModel.creator_id,
  createAt: bookModel.create_at,
  lastModifiedAt: bookModel.last_modified_at,
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
