import { useUser } from "@/contexts/UserContext";

function BookList() {
  const { userId } = useUser();

  return <div>Book List: {userId}</div>;
}

export default BookList;
