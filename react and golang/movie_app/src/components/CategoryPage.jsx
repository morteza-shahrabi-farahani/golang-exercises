import react from "react";
import { useParams } from "react-router-dom";

function CategoryPage() {
  const { category_name } = useParams();

  return (
    <>
      <h2>Category characteristic:</h2>
      <h2>{`category is: ${category_name}`}</h2>
    </>
  );
}

export default CategoryPage;
