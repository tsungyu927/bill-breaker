import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTrigger,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { JoinBookForm } from "@/interface/book";
import { joinBook } from "@/server/axios/book";
import { convertToJoinBookPayload } from "@/server/axios/book/converter";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useForm } from "react-hook-form";

function JoinBookDialog() {
  const [open, setOpen] = useState(false);
  const form = useForm<JoinBookForm>();
  const queryClient = useQueryClient();

  const { mutate } = useMutation({
    mutationFn: joinBook,
    onSuccess: () => {
      setOpen(false);
      // Re-fetch the book list
      queryClient.invalidateQueries({
        queryKey: ["bookList"],
      });
    },
  });

  const handleJoinBook = (data: JoinBookForm) => {
    mutate(convertToJoinBookPayload(data));
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="secondary">Join the book</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Join the existed book</DialogTitle>
          <DialogDescription>
            Please input Book ID to join the book
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            className="flex flex-col gap-8"
            onSubmit={form.handleSubmit(handleJoinBook)}
          >
            <div className="flex flex-col gap-4">
              <FormField
                control={form.control}
                name="id"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel required>Book ID</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="Please input the book id"
                        required
                        {...field}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />
            </div>
            <DialogFooter>
              <Button type="submit">Join</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}

export default JoinBookDialog;
