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
import { CustomRouteObject } from "./interface/router";

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

  const routes: CustomRouteObject[] = [
    {
      element: <Layout />,
      children: [
        {
          path: "/book",
          element: <BookList />,
          handle: { title: "Book" },
        },
        {
          path: "/book/:bookId",
          element: <BookDetail />,
          handle: { title: "Detail", goBackPath: "/book" },
          loader: authLoader,
        },
        {
          path: "/book/:bookId/cost/:costId",
          element: <CostDetail />,
          handle: {
            title: "Cost",
            goBackPath: "/book/:bookId",
          },
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
  ];

  const routers = createBrowserRouter(routes);

  return <RouterProvider router={routers} />;
}
