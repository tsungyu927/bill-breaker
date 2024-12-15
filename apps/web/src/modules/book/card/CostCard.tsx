import {
  Card,
  CardContent,
  CardDescription,
  CardTitle,
} from "@/components/ui/card";
import { BookMember, Cost } from "@/interface/book";
import { Link } from "react-router-dom";

function CostCard(props: Cost & { creator?: BookMember }) {
  const { id, bookId, title, creator, amount, currency } = props;

  return (
    <Card className="w-full max-w-lg border-none rounded-none cursor-pointer">
      <Link to={`/book/${bookId}/cost/${id}`}>
        <CardContent className="p-4 flex flex-col gap-2">
          <CardTitle className="w-full flex justify-between items-center">
            <p>{title}</p>
            <p>{creator?.userName}</p>
          </CardTitle>
          <CardDescription className="flex items-center gap-2">
            $<span className="font-bold">{amount}</span> {currency}
          </CardDescription>
        </CardContent>
      </Link>
    </Card>
  );
}

export default CostCard;
