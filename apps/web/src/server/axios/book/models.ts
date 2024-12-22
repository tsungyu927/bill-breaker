import { APIResponse } from "@/interface/api";

export interface BookModel {
  id: string;
  creator_id: string;
  book_name: string;
  book_description?: string;
  create_at: string;
  last_modified_at: string;
}

export interface CostModel {
  id: string;
  book_id: string;
  title: string;
  amount: number;
  currency: string;
  date: string;
  created_at: string;
  creator_id: string;
  description?: string;
}

export interface BookDetailModel extends BookModel {
  members: {
    user_id: string;
    user_name: string;
  }[];
  costs: CostModel[];
}

export interface PayerModel {
  cost_id: string;
  user_id: string;
  amount: number;
}

export interface SharerModel {
  cost_id: string;
  user_id: string;
  share_amount: number;
}

export interface CostDetailModel extends CostModel {
  payers: PayerModel[];
  sharers: SharerModel[];
}

// Response

export type CreateBookResponse = APIResponse<BookModel>;
export type BookListResponse = APIResponse<BookModel[]>;
export type BookResponse = APIResponse<BookDetailModel>;
export type JoinBookResponse = APIResponse<string>;
export type CreateCostResponse = APIResponse<{ id: string }>;

// Payload
export type CreateBookPayload = Pick<
  BookModel,
  "book_name" | "book_description"
>;
export type JoinBookPayload = {
  book_id: string;
};
export interface CreateCostPayload
  extends Omit<CostModel, "id" | "creator_id" | "created_at"> {
  payers: Omit<PayerModel, "cost_id">[];
  sharers: Omit<SharerModel, "cost_id">[];
}
