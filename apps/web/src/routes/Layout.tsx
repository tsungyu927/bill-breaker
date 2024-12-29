import { RegisterAction } from "@/interface/book";
import { createUser } from "@/server/axios/user";
import { useMutation } from "@tanstack/react-query";
import { Outlet } from "react-router-dom";
import RegisterDialog from "@/modules/RegisterDialog";
import { useUser } from "@/contexts/UserContext";
import { ThemeProvider } from "@/contexts/ThemeContext";
import { Toaster } from "@/components/ui/toaster";
import Header from "@/modules/Header";

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
      <div className="w-dvw h-dvh flex flex-col">
        <Header />
        <div className="w-full h-[calc(100%-64px)]">
          <Outlet />
        </div>
        <RegisterDialog open={!userId} onSubmit={handleRegisterSubmit} />
        <Toaster />
      </div>
    </ThemeProvider>
  );
}

export default Layout;
