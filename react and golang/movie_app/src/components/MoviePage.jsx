import react from "react";
import { useParams } from "react-router-dom";

function MoviePage() {
  const { id } = useParams();

  return (
    <>
      <h2>Movie is movie number:</h2>
      <h2>{`id is: ${id}`}</h2>
    </>
  );
}

export default MoviePage;
