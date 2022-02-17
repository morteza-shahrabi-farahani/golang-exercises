import react from "react";
import { Link } from "react-router-dom";

const categoryList = ["Drama", "Action", "Adventure"];

function Categories() {
  return (
    <>
      <h3>List of Categories</h3>
      <table className="movies-table">
        <tr>
          <th>ID</th>
          <th>name</th>
        </tr>

        {categoryList.map((element, index) => (
          <tr>
            <td>{index}</td>
            <td>
              <Link to={`/categories/${element}`}>{element}</Link>
              {/* {element.name} */}
            </td>
          </tr>
        ))}
      </table>
    </>
  );
}

export default Categories;
