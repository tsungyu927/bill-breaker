import { LoaderFunction, UIMatch } from "react-router-dom";

interface Handle {
  title: string;
  goBackPath?: string;
}

export interface IMatches extends UIMatch {
  data: unknown;
  handle: Handle;
}

export interface CustomRouteObject {
  path?: string;
  element?: React.ReactNode | null;
  errorElement?: React.ReactNode | null;
  children?: CustomRouteObject[];
  loader?: LoaderFunction | boolean;
  handle?: IMatches["handle"];
}
