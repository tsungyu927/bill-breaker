import {
  Book,
  BookDetail,
  Cost,
  CostDetail,
  CreateBookForm,
  CreateCostForm,
  JoinBookForm,
} from "@/interface/book";
import {
  BookDetailModel,
  BookModel,
  CostDetailModel,
  CreateBookPayload,
  CreateCostPayload,
  JoinBookPayload,
} from "./models";
import { formatISO, parseISO } from "date-fns";
import { combineDateTime } from "@/utils/datetime";

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
    title: cost.title,
    amount: cost.amount,
    currency: cost.currency,
    date: parseISO(cost.date),
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

export const convertCostDetailModelToCostDetail = (
  costDetailModel: CostDetailModel
): CostDetail => ({
  id: costDetailModel.id,
  bookId: costDetailModel.book_id,
  title: costDetailModel.title,
  amount: costDetailModel.amount,
  currency: costDetailModel.currency,
  date: parseISO(costDetailModel.date),
  creatorId: costDetailModel.creator_id,
  createdAt: costDetailModel.created_at,
  description: costDetailModel.description,
  payers: costDetailModel.payers.map((item) => ({
    userId: item.user_id,
    amount: item.amount,
  })),
  sharers: costDetailModel.sharers.map((item) => ({
    userId: item.user_id,
    amount: item.share_amount,
  })),
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

export const convertToCreateCostPayload = (
  cost: CreateCostForm
): CreateCostPayload => ({
  book_id: cost.bookId,
  title: cost.title,
  amount: cost.amount,
  currency: cost.currency,
  description: cost.description,
  date: formatISO(combineDateTime(cost.date, cost.time)),
  payers: cost.payers.map((item) => ({
    user_id: item.userId,
    amount: item.amount,
  })),
  sharers: cost.sharers.map((item) => ({
    user_id: item.userId,
    share_amount: item.amount,
  })),
});
