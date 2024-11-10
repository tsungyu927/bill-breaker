export enum RegisterAction {
  Create = "create",
  Restore = "restore",
}

export interface Book {
  id: string;
  name: string;
  description?: string;
  creatorId: string;
  createAt: string;
  lastModifiedAt: string;
}

// Create book form
export type CreateBookForm = Pick<Book, "name" | "description">;
