import { RegisterAction } from "@/interface/book";
import { createUser } from "@/server/axios/user";
import { useMutation } from "@tanstack/react-query";
import { Outlet } from "react-router-dom";
import RegisterDialog from "@/modules/book/RegisterDialog";
import { useUser } from "@/contexts/UserContext";

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
    <div>
      <Outlet />
      <RegisterDialog open={!userId} onSubmit={handleRegisterSubmit} />
    </div>
  );
}

export default Layout;
