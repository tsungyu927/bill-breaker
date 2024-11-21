import BookCard from "@/modules/book/BookCard";
import CreateBookDialog from "@/modules/book/CreateBookDialog";
import { fetchBookList } from "@/server/axios/book";
import { useQuery } from "@tanstack/react-query";

function BookList() {
  const { data: bookList } = useQuery({
    queryKey: ["bookList"],
    queryFn: fetchBookList,
    initialData: [],
  });

  return (
    <div className="flex flex-col p-4">
      <div className="flex flex-1 flex-col gap-4 p-4">
        {bookList.map((book) => (
          <BookCard key={book.id} {...book} />
        ))}
      </div>
      <CreateBookDialog />
    </div>
  );
}

export default BookList;
