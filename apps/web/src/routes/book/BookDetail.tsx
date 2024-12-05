import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { fetchBookByBookId } from "@/server/axios/book";
import { PersonIcon, StarFilledIcon } from "@radix-ui/react-icons";
import { useQuery } from "@tanstack/react-query";
import React from "react";
import { Link, useParams } from "react-router-dom";
import { type BookDetail } from "@/interface/book";

function BookDetail() {
  const { bookId } = useParams();
  const { data } = useQuery({
    queryKey: ["bookDetail", bookId],
    queryFn: () => {
      if (bookId) {
        return fetchBookByBookId(bookId);
      }
    },
    enabled: !!bookId,
    initialData: null,
  });

  const {
    name = "",
    description = "",
    creatorId = "",
    members = [],
    costs = [],
  } = data || {};

  const memberList = members.reduce((acc, member) => {
    if (member.userId === creatorId) {
      // Unshift member to head of the array
      acc.unshift(member);
    } else {
      // Push member to the tail of the array
      acc.push(member);
    }
    return acc;
  }, [] as BookDetail["members"]);

  return (
    <div className="w-full h-full flex flex-col p-4 gap-4">
      <Card className="w-full break-words">
        <CardHeader>
          <CardTitle>{name}</CardTitle>
          <CardDescription>{description}</CardDescription>
        </CardHeader>
        <Separator className="mb-4" />
        <CardContent className="flex flex-col gap-2">
          <div className="flex flex-col">
            <div className="flex flex-col gap-4">
              {memberList.map((member, index) => (
                <div key={member.userId} className="flex gap-4 items-center">
                  <span>
                    {index === 0 ? <StarFilledIcon /> : <PersonIcon />}
                  </span>
                  {member.userName}
                </div>
              ))}
            </div>
          </div>
        </CardContent>
      </Card>
      <ScrollArea className="w-full flex-grow rounded-md border">
        {costs.map((cost) => {
          const creator = members.find(
            (member) => member.userId === cost.creatorId
          );

          return (
            <React.Fragment key={cost.id}>
              <Button
                variant="ghost"
                className="w-full h-full p-6 flex flex-col gap-2"
                asChild
              >
                <Link to={`/book/${bookId}/cost/${cost.id}`}>
                  <div className="w-full flex justify-between items-center">
                    <p>
                      $<span className="font-bold">{cost.amount}</span>{" "}
                      {cost.currency}
                    </p>
                    <p>{creator?.userName}</p>
                  </div>
                  <div className="w-full text-sm text-muted-foreground whitespace-pre-wrap">
                    {cost.description && <p>{cost.description}</p>}
                  </div>
                </Link>
              </Button>
              <Separator />
            </React.Fragment>
          );
        })}
      </ScrollArea>
    </div>
  );
}

export default BookDetail;
