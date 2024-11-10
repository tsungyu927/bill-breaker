import { Book } from "@/interface/book";
import { BookModel } from "./models";

export const convertBookModelToBook = (bookModel: BookModel): Book => ({
  id: bookModel.id,
  name: bookModel.book_name,
  description: bookModel.book_description,
  creatorId: bookModel.creator_id,
  createAt: bookModel.create_at,
  lastModifiedAt: bookModel.last_modified_at,
});

export const convertBookToBookModel = (book: Book): BookModel => ({
  id: book.id,
  book_name: book.name,
  book_description: book.description,
  creator_id: book.creatorId,
  create_at: book.createAt,
  last_modified_at: book.lastModifiedAt,
});
