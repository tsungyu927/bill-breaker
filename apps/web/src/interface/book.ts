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
  createAt: string;
  lastModifiedAt: string;
}

export interface BookDetail extends Book {
  members: BookMember[];
}

// Create book form
export type CreateBookForm = Pick<Book, "name" | "description">;

// Join book form
export type JoinBookForm = Pick<Book, "id">;
