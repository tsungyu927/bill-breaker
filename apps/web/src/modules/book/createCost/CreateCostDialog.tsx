import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTrigger,
  DialogDescription,
  DialogFooter,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormLabel,
  FormItem,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { BookMember, CreateCostForm } from "@/interface/book";
import { convertToCreateCostPayload } from "@/server/axios/book/converter";
import { FormProvider, SubmitHandler, useForm } from "react-hook-form";
import MemberSelectList from "./MemberSelectList";
import { ScrollArea } from "@/components/ui/scroll-area";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createCost } from "@/server/axios/book";

interface Props {
  bookId: string;
  members: BookMember[];
}

function CreateCostDialog(props: Props) {
  const { bookId, members } = props;
  const [open, setOpen] = useState(false);

  const queryClient = useQueryClient();
  const { mutate } = useMutation({
    mutationFn: createCost,
    onSuccess: () => {
      setOpen(false);
      // Re-fetch the book list
      queryClient.invalidateQueries({
        queryKey: ["bookDetail", bookId],
      });
    },
  });

  const DEFAULT_FORM_VALUES: CreateCostForm = {
    bookId,
    title: "",
    amount: 0,
    currency: "TWD",
    payers: members.map((member) => ({ userId: member.userId, amount: 0 })),
    sharers: members.map((member) => ({ userId: member.userId, amount: 0 })),
  };
  const form = useForm<CreateCostForm>({
    defaultValues: DEFAULT_FORM_VALUES,
    mode: "onChange",
    delayError: 500,
  });

  const handleCreateCost: SubmitHandler<CreateCostForm> = (
    data: CreateCostForm
  ) => {
    const payload = convertToCreateCostPayload(data);
    mutate(payload);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Create new cost</Button>
      </DialogTrigger>
      <FormProvider {...form}>
        <Form {...form}>
          <form
            id="create-cost-form"
            onSubmit={form.handleSubmit(handleCreateCost)}
          >
            <DialogContent className="max-h-full flex flex-col">
              <DialogHeader>
                <DialogTitle>Create new cost</DialogTitle>
                <DialogDescription>
                  Please input below information for you new cost
                </DialogDescription>
              </DialogHeader>
              <ScrollArea className="flex-grow overflow-y-auto p-0">
                <div className="flex flex-col gap-4">
                  <FormField
                    control={form.control}
                    name="title"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel required>Title</FormLabel>
                        <FormControl>
                          <Input {...field} />
                        </FormControl>
                      </FormItem>
                    )}
                  />
                  <div className="w-full flex gap-4">
                    <FormField
                      control={form.control}
                      name="currency"
                      render={({ field }) => (
                        <FormItem>
                          <FormLabel required>Currency</FormLabel>
                          <Select
                            onValueChange={field.onChange}
                            defaultValue={field.value}
                          >
                            <FormControl>
                              <SelectTrigger>
                                <SelectValue placeholder="$" />
                              </SelectTrigger>
                            </FormControl>
                            <SelectContent>
                              <SelectItem value="TWD">TWD</SelectItem>
                            </SelectContent>
                          </Select>
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name="amount"
                      render={({ field }) => (
                        <FormItem className="flex-grow">
                          <FormLabel required>Amount</FormLabel>
                          <FormControl>
                            <Input
                              type="number"
                              value={field.value}
                              onChange={(e) => {
                                if (e.target.value === "") {
                                  return field.onChange("");
                                }
                                field.onChange(Number(e.target.value));
                              }}
                            />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                  </div>
                  <div>
                    <FormLabel required>Payers</FormLabel>
                    <MemberSelectList name="payers" members={members} />
                  </div>
                  <div>
                    <FormLabel required>Sharers</FormLabel>
                    <MemberSelectList name="sharers" members={members} />
                  </div>
                  <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Description</FormLabel>
                        <FormControl>
                          <Textarea {...field} />
                        </FormControl>
                      </FormItem>
                    )}
                  />
                </div>
              </ScrollArea>
              <DialogFooter>
                <Button type="submit" form="create-cost-form">
                  Create
                </Button>
              </DialogFooter>
            </DialogContent>
          </form>
        </Form>
      </FormProvider>
    </Dialog>
  );
}

export default CreateCostDialog;
