/* eslint-disable react-refresh/only-export-components */
import { getStorage, setStorage } from "@/utils/storages";
import { createContext, useContext, useEffect, useMemo, useState } from "react";

interface UserContextType {
  userId: string | null;
  setUserId: (id: string | null) => void;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [userId, setUserId] = useState<string | null>(() =>
    getStorage("user_id", window.localStorage)
  );

  useEffect(() => {
    setStorage("user_id", userId, window.localStorage);
  }, [userId]);

  const contextValue = useMemo(
    () => ({ userId, setUserId }),
    [userId, setUserId]
  );

  return (
    <UserContext.Provider value={contextValue}>{children}</UserContext.Provider>
  );
};

export const useUser = (): UserContextType => {
  const context = useContext(UserContext);

  if (!context) throw new Error("useUser must be used within a UserProvider");
  return context;
};
