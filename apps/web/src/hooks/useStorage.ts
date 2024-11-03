import {
  useCallback,
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
} from "react";
import { getStorage, setStorage } from "../utils/storages";

type StorageType = typeof window.localStorage | typeof window.sessionStorage;

export function useLocalStorage(key: string) {
  return useStorage(key, window.localStorage);
}

export function useSessionStorage(key: string) {
  return useStorage(key, window.sessionStorage);
}

function useStorage(
  key: string,
  storageObject: StorageType
): {
  value: string | null;
  setValue: Dispatch<SetStateAction<string | null>>;
  remove: () => void;
} {
  const [value, setValue] = useState(() => {
    return getStorage(key, storageObject);
  });

  useEffect(() => {
    setStorage(key, value, storageObject);
  }, [key, value, storageObject]);

  const remove = useCallback(() => {
    setValue(null);
  }, []);

  return { value, setValue, remove };
}
