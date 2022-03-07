import { React, useState, useEffect } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";

function AddOrEditMovie() {
  const [release_date, setRelease_date] = useState("");
  const { id } = useParams();
  const [movie, setMovie] = useState({});
  // setMovie({
  //   title: "",
  //   releaseDate: "",
  //   runtime: "",
  //   mppaRating: "",
  //   rating: "",
  //   description: "",
  // });
  // if (id !== undefined) {
  useEffect(() => {
    if (id != undefined) {
      axios.get(`http://localhost:8080/movie/${id}`).then((res) => {
        console.log(res);
        console.log(res.data.movie);
        setMovie(res.data.movie);
        // setIsLoaded(true);
      });
    }
  }, []);
  // } else {
  //   setMovie({
  //     title: "",
  //     releaseDate: "",
  //     runtime: "",
  //     mppaRating: "",
  //     rating: "",
  //     description: "",
  //   });
  // }
  // const [movie, setMovie] = useState({
  //   title: "",
  //   releaseDate: "",
  //   runtime: "",
  //   mppaRating: "",
  //   rating: "",
  //   description: "",
  // });

  function handleChangeInput(event) {
    let value = event.target.value;
    let name = event.target.name;
    setMovie((prevState) => ({
      ...prevState,
      [name]: value,
    }));
    // console.log(movie);
  }

  function handleSubmit(event) {
    event.preventDefault();
    console.log(movie);
  }

  return (
    <>
      {/* {console.log(movie.mpaa_rating)} */}
      {/* {console.log(movie.release_date)} */}
      {console.log(typeof movie.release_date)}
      <h2>Add or Edit Movie</h2>
      <form onSubmit={handleSubmit}>
        <label htmlFor="title" className="form-label">
          Title
        </label>
        <input
          type="text"
          name="title"
          id="title"
          value={movie.title}
          className="form-control"
          onChange={handleChangeInput}
        />
        <br></br>
        <label htmlFor="release_date" className="form-label">
          Release Date
        </label>
        <input
          type="text"
          name="release_date"
          id="release_date"
          value={
            movie.release_date
              ? movie.release_date.split("T")[0]
              : movie.release_date
          }
          className="form-control"
          onChange={handleChangeInput}
        />
        <br></br>
        <label htmlFor="runtime" className="form-label">
          Runtime
        </label>
        <input
          type="text"
          name="runtime"
          id="runtime"
          value={movie.runtime}
          className="form-control"
          onChange={handleChangeInput}
        />
        <br />
        <label htmlFor="mppa_rating" className="form-label">
          MPAA RAting
        </label>
        <select
          className="form-select"
          name="mppa_rating"
          id="mppa_rating"
          value={movie.mpaa_rating}
          onChange={handleChangeInput}
        >
          <option className="form-select">Choose...</option>
          <option className="form-select" value="G">
            G
          </option>
          <option className="form-select" value="PG">
            PG
          </option>
          <option className="form-select" value="PG13">
            PG13
          </option>
          <option className="form-select" value="R">
            R
          </option>
          <option className="form-select" value="NC17">
            NC17
          </option>
        </select>

        <br></br>
        <label htmlFor="rating" className="form-label">
          Runtime
        </label>
        <input
          type="text"
          name="rating"
          id="rating"
          value={movie.rating}
          className="form-control"
          onChange={handleChangeInput}
        />

        <br></br>
        <label htmlFor="description" className="form-label">
          Description
        </label>
        <textarea
          name="description"
          id="description"
          value={movie.description}
          className="form-control"
          onChange={handleChangeInput}
        />

        <br />
        <button className="btn btn-primary">Save</button>
      </form>
    </>
  );
}

export default AddOrEditMovie;
