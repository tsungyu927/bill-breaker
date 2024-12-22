import { Button } from "@/components/ui/button";
import { IMatches } from "@/interface/router";
import { ArrowLeftIcon } from "@radix-ui/react-icons";
import { Params, useLocation, useMatches, useNavigate } from "react-router-dom";
import Settings from "./Settings";

const resolvePath = (path: string, params?: Params<string>) => {
  return path.replace(/\/:([^/]+)/g, (_match, key) => {
    return `/${params?.[key] || ""}`;
  });
};

function Header() {
  const location = useLocation();
  const navigate = useNavigate();
  const matches = useMatches() as IMatches[];

  // Get current route metadata
  const currentRoute = matches.find(
    (match) => match.pathname === location.pathname
  );
  const params = currentRoute?.params;
  const routeTitle = currentRoute?.handle?.title; // Define `handle.title` in your routes
  const goBackPath = currentRoute?.handle?.goBackPath;

  const handleGoBack = () => {
    const path = resolvePath(goBackPath!, params);
    if (path) {
      navigate(path);
      return;
    }
    navigate(-1);
  };

  return (
    <header className="flex h-16 w-full items-center border-b border-gray-200 dark:border-gray-800">
      <div className="mx-auto max-w-7xl w-full px-4 md:px-6 grid grid-cols-3">
        <div className="flex justify-start">
          {goBackPath && (
            <Button onClick={handleGoBack} variant="outline" size="icon">
              <ArrowLeftIcon />
            </Button>
          )}
        </div>
        <h1 className="text-xl font-semibold flex justify-center items-center">
          {routeTitle}
        </h1>
        <div className="flex justify-end">
          <Settings />
        </div>
      </div>
    </header>
  );
}

export default Header;
