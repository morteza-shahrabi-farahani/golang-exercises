import { react, useState, useEffect } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";

function CategoryPage() {
  const { genre_name } = useParams();
  const [movies, setMovies] = useState([]);
  useEffect(() => {
    axios.get(`http://localhost:8080/genres/${genre_name}`).then((res) => {
      console.log(res);
      console.log(res.data.movies);
      setMovies(res.data.movies);
      console.log(movies);
      // setIsLoaded(true);
    });
  }, []);
  return (
    <>
      {console.log(movies)}
      <h2>{`category is: ${genre_name}`}</h2>
      <table className="movies-table">
        
        {movies.map((element, index) => (
          <tr>
            <td>{index}</td>
            <td>
              <Link to={`/movies/${element.id}`}>{element.title}</Link>
            </td>
          </tr>
        ))}
      </table>
    </>
  );
}

export default CategoryPage;
