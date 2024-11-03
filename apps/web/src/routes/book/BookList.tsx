import { useLocalStorage } from "@/hooks/useStorage";
import { RegisterAction } from "@/interface/book";
import RegisterDialog from "@/modules/book/RegisterDialog";

function BookList() {
  const { value: userId } = useLocalStorage("user_id");

  const handleRegisterSubmit = (props: {
    type: RegisterAction;
    value: string;
  }) => {
    //TODO: call api to create/restore user
    console.log(props);
  };

  return (
    <>
      <RegisterDialog open={!userId} onSubmit={handleRegisterSubmit} />
    </>
  );
}

export default BookList;
