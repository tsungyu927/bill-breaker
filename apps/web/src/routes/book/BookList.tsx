import { BookCard } from "@/modules/book/card";
import CreateBookDialog from "@/modules/book/createBook";
import JoinBookDialog from "@/modules/book/joinBook";
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
      <div className="flex flex-col gap-2">
        <CreateBookDialog />
        <JoinBookDialog />
      </div>
    </div>
  );
}

export default BookList;
