import { react, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const categoryList = ["Drama", "Action", "Adventure"];

function Categories() {
  const [genres, setGenres] = useState([]);
  useEffect(() => {
    axios.get("http://localhost:8080/genres").then((res) => {
      console.log(res);
      setGenres(res.data.genres);
      console.log(genres);
      // setIsLoaded(true);
    });
  }, []);

  return (
    <>
      <h3>List of Genres</h3>
      <table className="movies-table">
        <tr>
          <th>ID</th>
          <th>name</th>
        </tr>

        {genres.map((element, index) => (
          <tr>
            <td>{index}</td>
            <td>
              <Link to={`/genres/${element.genre_name}`}>
                {element.genre_name}
              </Link>
            </td>
          </tr>
        ))}
      </table>
    </>
  );
}

export default Categories;
