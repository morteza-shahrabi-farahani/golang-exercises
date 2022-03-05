import { react, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

function MoviePage() {
  const { id } = useParams();
  const [genres, setGenres] = useState([]);
  const [title, setTitle] = useState("");
  const [rating, setRating] = useState("");
  const [description, setDescription] = useState("");
  const [runtime, setRuntime] = useState("");
  const [mppaRating, setMppaRating] = useState("");
  const [isLoaded, setIsLoaded] = useState(false);
  useEffect(() => {
    axios.get(`http://localhost:8080/movie/${id}`).then((res) => {
      // console.log(res);
      setTitle(res.data.movie.title);
      setRating(res.data.movie.rating);
      setDescription(res.data.movie.description);
      setRuntime(res.data.movie.runtime);
      setMppaRating(res.data.movie.mppaRating);
      // setGenres(res.data.movie.movie_genres);
      if (res.data.movie.movie_genres) {
        setGenres(Object.values(res.data.movie.movie_genres));
      } else {
        setGenres([]);
      }
      setIsLoaded(true);
      // console.log(res.data.movie);
    });
  }, []);

  return (
    <>
      {/* {console.log(genres)} */}
      <h2>{title}</h2>
      <h4>rating: {rating}</h4>
      <h4>runtime: {runtime}</h4>
      <h4>{mppaRating}</h4>
      <p>{description}</p>
      {genres.map((element, index) => (
        <p>{element}</p>
      ))}
    </>
  );
}

export default MoviePage;
