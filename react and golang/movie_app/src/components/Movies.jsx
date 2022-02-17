import react from "react";
import { Link } from "react-router-dom";
import "./Movies.css";

const list = [
  { id: 1, name: "Batman darknight", runtime: "120", category: "Action" },
  { id: 2, name: "Godfather2", runtime: "121", category: "Adventure" },
  { id: 3, name: "ForrestGump", runtime: "119", category: "Drama" },
];

function Movies() {
  return (
    <>
      <h3>List of movies</h3>
      <table className="movies-table">
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Runtime</th>
          <th>Category</th>
        </tr>

        {list.map((element, index) => (
          <tr>
            <td>{element.id}</td>
            <td>
              <Link to={`/movies/${element.id}`}>{element.name}</Link>
              {/* {element.name} */}
            </td>
            <td>{element.runtime}</td>
            <td>
              <Link to={`/categories/${element.category}`}>
                {element.category}
              </Link>
            </td>
          </tr>
        ))}
      </table>
    </>
  );
}

export default Movies;
