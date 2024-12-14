import {
  Book,
  BookDetail,
  Cost,
  CreateBookForm,
  CreateCostForm,
  JoinBookForm,
} from "@/interface/book";
import {
  BookDetailModel,
  BookModel,
  CreateBookPayload,
  CreateCostPayload,
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

export const convertToCreateCostPayload = (
  cost: CreateCostForm
): CreateCostPayload => ({
  book_id: cost.bookId,
  amount: cost.amount,
  currency: cost.currency,
  description: cost.description,
  payers: cost.payers.map((item) => ({
    user_id: item.userId,
    amount: item.amount,
  })),
  sharers: cost.sharers.map((item) => ({
    user_id: item.userId,
    share_amount: item.amount,
  })),
});
