import { useParams } from "react-router-dom";

function CostDetail() {
  const { bookId, costId } = useParams();

  return (
    <div>
      <p>bookId: {bookId}</p>
      <p>costId: {costId}</p>
    </div>
  );
}

export default CostDetail;
