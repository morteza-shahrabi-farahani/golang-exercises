import { react, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import "./Movies.css";

function Movies() {
  const [movieList, setMovieList] = useState([]);
  const [isLoaded, setIsLoaded] = useState(false);
  useEffect(() => {
    axios.get("http://localhost:8080/movies").then((res) => {
      // console.log(res);
      setMovieList(res.data.movies);
      setIsLoaded(true);
    });
  });

  if (!isLoaded) {
    return <p>Loading...</p>;
  } else {
    return (
      <>
        {/* {console.log(isLoaded)} */}
        {/* {console.log(movieList.movies)} */}
        <h3>List of movies</h3>
        <table className="movies-table">
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Runtime</th>
          </tr>

          {movieList.map((element, index) => (
            <tr>
              <td>{element.id}</td>
              <td>
                <Link to={`/movies/${element.id}`}>{element.title}</Link>
                {/* {element.name} */}
              </td>
              <td>{element.runtime}</td>
            </tr>
          ))}
        </table>
      </>
    );
  }
}

export default Movies;
