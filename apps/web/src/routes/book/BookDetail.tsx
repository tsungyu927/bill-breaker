import { fetchBookByBookId } from "@/server/axios/book";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";

function BookDetail() {
  const { bookId } = useParams();
  const { data } = useQuery({
    queryKey: ["bookDetail", bookId],
    queryFn: () => {
      if (bookId) {
        return fetchBookByBookId(bookId);
      }
    },
    enabled: !!bookId,
    initialData: null,
  });

  return <div>detail: {JSON.stringify(data)}</div>;
}

export default BookDetail;
