import { useLocalStorage } from "@/hooks/useStorage";

function BookList() {
  const { value: userId } = useLocalStorage("user_id");

  return <div>{userId}</div>;
}

export default BookList;
