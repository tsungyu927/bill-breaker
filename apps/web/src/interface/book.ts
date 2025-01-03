export enum RegisterAction {
  Create = "create",
  Restore = "restore",
}

export interface BookMember {
  userId: string;
  userName: string;
}

export interface Book {
  id: string;
  name: string;
  description?: string;
  creatorId: string;
  createdAt: string;
  lastModifiedAt: string;
}

export interface Cost {
  id: string;
  bookId: string;
  title: string;
  amount: number;
  currency: string;
  date: Date;
  creatorId: string;
  createdAt: string;
  description?: string;
}

export interface CostDetail extends Cost {
  payers: CostMember[];
  sharers: CostMember[];
}

export interface BookDetail extends Book {
  members: BookMember[];
  costs: Cost[];
}

export interface CostMember {
  userId: string;
  amount: number;
}

// Create book form
export type CreateBookForm = Pick<Book, "name" | "description">;

// Join book form
export type JoinBookForm = Pick<Book, "id">;

// Create cost form
export interface CreateCostForm
  extends Omit<Cost, "id" | "creatorId" | "createdAt"> {
  time: Date;
  payers: CostMember[];
  sharers: CostMember[];
}
