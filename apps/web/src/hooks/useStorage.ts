import {
  useCallback,
  useState,
  useEffect,
  Dispatch,
  SetStateAction,
} from "react";

type StorageType = typeof window.localStorage | typeof window.sessionStorage;

export function useLocalStorage<T>(key: string, defaultValue: T) {
  return useStorage(key, defaultValue, window.localStorage);
}

export function useSessionStorage<T>(key: string, defaultValue: T) {
  return useStorage(key, defaultValue, window.sessionStorage);
}

function useStorage<T>(
  key: string,
  defaultValue: T,
  storageObject: StorageType
): [T, Dispatch<SetStateAction<T>>, () => void] {
  const [value, setValue] = useState(() => {
    const jsonValue = storageObject.getItem(key);

    if (jsonValue != null) return JSON.parse(jsonValue);
    return defaultValue;
  });

  useEffect(() => {
    if (value === undefined) return storageObject.removeItem(key);
    storageObject.setItem(key, JSON.stringify(value));
  }, [key, value, storageObject]);

  const remove = useCallback(() => {
    setValue(undefined as unknown as T);
  }, []);

  return [value, setValue, remove];
}
