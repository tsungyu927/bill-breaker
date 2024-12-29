import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import Loading from "@/modules/Loading";
import { fetchCostDetail } from "@/server/axios/book";
import { useQuery } from "@tanstack/react-query";
import { format } from "date-fns";
import { useParams } from "react-router-dom";

function CostDetail() {
  const { bookId, costId } = useParams();
  const { isFetching, data } = useQuery({
    queryKey: ["costDetail", bookId, costId],
    queryFn: () => {
      if (bookId && costId) {
        return fetchCostDetail(bookId, costId);
      }
    },
    enabled: !!bookId && !!costId,
    initialData: null,
  });

  const { title, description, date, amount, currency } = data || {};

  if (isFetching || !bookId || !costId) {
    return <Loading />;
  }
  return (
    <div className="w-full h-full flex flex-col p-4 gap-4">
      <Card className="w-full break-words">
        <CardHeader>
          <CardTitle>{title}</CardTitle>
          <CardDescription>{description}</CardDescription>
        </CardHeader>
        <Separator className="mb-4" />
        <CardContent className="flex flex-col gap-2">
          <p className="font-bold text-xl">
            ${amount} {currency}
          </p>
          <p className="text-right text-sm text-muted-foreground">
            {format(date || "", "yyyy-MM-dd HH:mm")}
          </p>
        </CardContent>
      </Card>
    </div>
  );
}

export default CostDetail;
