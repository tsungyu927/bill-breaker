import { fetchBookList } from "@/src/server/axios/book";
import { useQuery } from "@tanstack/react-query";

function Homepage() {
  const { data: books, isLoading } = useQuery({
    queryKey: ["books"],
    queryFn: fetchBookList,
  });

  if (isLoading) return <div>Loading...</div>;
  return books?.map((book) => (
    <div key={book.ID}>
      <p>ID: {book.ID}</p>
      <p>Book ID: {book.BookId}</p>
      <p>Book Name: {book.BookName}</p>
      {/* <p>Created Date: {book.CreatedDate}</p>
      <p>LastModified Date: {book.LastModifiedDate}</p> */}
    </div>
  ));
}

export default Homepage;
