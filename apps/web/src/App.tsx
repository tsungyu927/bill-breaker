import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import BookList from "@/routes/book/BookList";
import ErrorPage from "@/routes/Error";
import { useLocalStorage } from "./hooks/useStorage";
import BookDetail from "./routes/book/BookDetail";

export default function App() {
  const { value: userId } = useLocalStorage("user_id");

  const basePathLoader = async () => {
    return redirect("/book");
  };

  const authLoader = async () => {
    if (!userId) {
      return redirect("/");
    }
    return null;
  };

  const routers = createBrowserRouter([
    {
      path: "/book",
      element: <BookList />,
    },
    {
      path: "/book/:bookId",
      element: <BookDetail />,
      loader: authLoader,
    },
    {
      path: "/",
      loader: basePathLoader,
    },
    {
      path: "*",
      errorElement: <ErrorPage />,
    },
  ]);

  return <RouterProvider router={routers} />;
}
