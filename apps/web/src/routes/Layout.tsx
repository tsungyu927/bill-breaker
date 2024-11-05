import { RegisterAction } from "@/interface/book";
import { createUser } from "@/server/axios/user";
import { useMutation } from "@tanstack/react-query";
import { Outlet } from "react-router-dom";
import RegisterDialog from "@/modules/RegisterDialog";
import { useUser } from "@/contexts/UserContext";
import { ThemeProvider } from "@/contexts/ThemeContext";
import Settings from "@/modules/Settings";
import { Toaster } from "@/components/ui/toaster";

function Layout() {
  const { userId, setUserId } = useUser();
  const { mutate } = useMutation({
    mutationFn: createUser,
    onSuccess: (data) => {
      if (data) {
        setUserId(data.id);
      }
    },
  });

  const handleRegisterSubmit = (props: {
    type: RegisterAction;
    value: string;
  }) => {
    const { type, value } = props;

    if (type === RegisterAction.Create) {
      mutate({ name: value });
    }
    //TODO: call api to restore user
  };

  return (
    <ThemeProvider>
      <Outlet />
      <RegisterDialog open={!userId} onSubmit={handleRegisterSubmit} />
      <Settings />
      <Toaster />
    </ThemeProvider>
  );
}

export default Layout;
