import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import BookList from "@/routes/book/BookList";
import ErrorPage from "@/routes/Error";
import BookDetail from "./routes/book/BookDetail";
import Layout from "./routes/Layout";
import { useUser } from "./contexts/UserContext";
import CostDetail from "./routes/book/CostDetail";

export default function App() {
  const { userId } = useUser();
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
      element: <Layout />,
      children: [
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
          path: "/book/:bookId/cost/:costId",
          element: <CostDetail />,
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
      ],
    },
  ]);

  return <RouterProvider router={routers} />;
}
