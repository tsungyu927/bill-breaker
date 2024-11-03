import { isNil } from "./helpers";

type StorageType = typeof window.localStorage | typeof window.sessionStorage;

export const getStorage = (
  key: string,
  storageType: StorageType = window.localStorage
): string | null => {
  const jsonValue = storageType.getItem(key);

  if (jsonValue !== null) return JSON.parse(jsonValue);
  return null;
};

export const setStorage = <T>(
  key: string,
  value: T,
  storageType: StorageType = window.localStorage
): void => {
  if (isNil(value)) {
    deleteStorage(key);
    return;
  }

  const jsonValue = JSON.stringify(value);
  storageType.setItem(key, jsonValue);
};

export const deleteStorage = (
  key: string,
  storageType: StorageType = window.localStorage
): void => {
  storageType.removeItem(key);
};
