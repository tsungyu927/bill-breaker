import {
  Card,
  CardContent,
  CardDescription,
  CardTitle,
} from "@/components/ui/card";
import { Book } from "@/interface/book";
import { Link } from "react-router-dom";

function BookCard(props: Book) {
  const { id, name, description } = props;

  return (
    <Card className="w-full max-w-lg cursor-pointer">
      <Link to={`/book/${id}`}>
        <CardContent className="p-4 flex flex-col gap-2">
          <CardTitle>{name}</CardTitle>
          {description && (
            <CardDescription className="flex items-center gap-2">
              {description}
            </CardDescription>
          )}
        </CardContent>
      </Link>
    </Card>
  );
}

export default BookCard;
