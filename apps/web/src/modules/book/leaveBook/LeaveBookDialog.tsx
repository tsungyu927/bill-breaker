import { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useMutation } from "@tanstack/react-query";
import { leaveBook } from "@/server/axios/book";
import { useNavigate, useParams } from "react-router-dom";

function LeaveBookDialog() {
  const { bookId } = useParams();
  const [open, setOpen] = useState(false);
  const navigate = useNavigate();

  const { mutate } = useMutation({
    mutationFn: leaveBook,
    onSuccess: () => {
      setOpen(false);
      navigate("/book");
    },
  });

  const handleLeave = () => {
    if (bookId) {
      mutate(bookId);
    }
  };

  const handleCloseDialog = () => {
    setOpen(false);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="destructive">Leave</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Leave Book</DialogTitle>
          <DialogDescription>
            Are you sure you want to leave this book?
          </DialogDescription>
        </DialogHeader>
        <DialogFooter className="flex gap-2">
          <Button variant="outline" onClick={handleCloseDialog}>
            Cancel
          </Button>
          <Button variant="destructive" onClick={handleLeave}>
            Leave
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default LeaveBookDialog;
