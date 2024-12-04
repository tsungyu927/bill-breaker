import {
  Book,
  BookDetail,
  Cost,
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
  createdAt: bookModel.create_at,
  lastModifiedAt: bookModel.last_modified_at,
});

export const convertBookDetailModelToBookDetail = (
  bookDetailModel: BookDetailModel
): BookDetail => {
  const costs = bookDetailModel.costs.map<Cost>((cost) => ({
    id: cost.id,
    bookId: cost.book_id,
    amount: cost.amount,
    currency: cost.currency,
    creatorId: cost.creator_id,
    createdAt: cost.created_at,
    description: cost.description,
  }));

  return {
    id: bookDetailModel.id,
    name: bookDetailModel.book_name,
    description: bookDetailModel.book_description,
    members: bookDetailModel.members.map((item) => ({
      userId: item.user_id,
      userName: item.user_name,
    })),
    costs,
    creatorId: bookDetailModel.creator_id,
    createdAt: bookDetailModel.create_at,
    lastModifiedAt: bookDetailModel.last_modified_at,
  };
};

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
