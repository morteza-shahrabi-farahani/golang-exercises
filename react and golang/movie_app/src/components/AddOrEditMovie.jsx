import { React, useState, useEffect } from "react";
import axios from "axios";
import { useParams, Link, useNavigate } from "react-router-dom";
import Alert from "../ui-components/Alert";

function AddOrEditMovie(props) {
  const navigate = useNavigate();
  const [loaded, setloaded] = useState(false);
  const { id } = useParams();
  const [movie, setMovie] = useState({});
  const [alert, setAlert] = useState({
    type: "d-none",
    message: "",
  });
  useEffect(() => {
    if (props.do == "edit") {
      axios.get(`http://localhost:8080/movie/${id}`).then((res) => {
        console.log(res);
        console.log(res.data.movie);
        setMovie(res.data.movie);
        console.log("hellohellohello");
        id = null;
        setloaded(true);
        setAlert({
          type: "d-none",
          message: "",
        });
      });
    } else {
      setMovie({
        title: "",
        release_date: "",
        runtime: "",
        mpaa_rating: "",
        rating: "",
        description: "",
      });
      setAlert({
        type: "d-none",
        message: "",
      });
    }
  }, [props.do]);

  function handleChangeInput(event) {
    let value = event.target.value;
    let name = event.target.name;
    setMovie((prevState) => ({
      ...prevState,
      [name]: value,
    }));
    console.log(movie);
    console.log(name);
    console.log(value);
  }

  function handleSubmit(event) {
    event.preventDefault();
    console.log(movie);
    const data = new FormData(event.target);
    const sendData = Object.fromEntries(data.entries());
    let requestOptions;
    let url;

    if (props.do == "add") {
      requestOptions = {
        method: "POST",
        body: JSON.stringify(sendData),
      };
      url = "http://localhost:8080/admin/add-movie";
    } else if (props.do == "edit") {
      sendData.id = id;
      requestOptions = {
        method: "PUT",
        body: JSON.stringify(sendData),
      };
      url = `http://localhost:8080/admin/movie/${id}/edit`;
    }

    console.log(sendData);

    fetch(url, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        // console.log(data);
        if (data.error) {
          setAlert({ type: "alert-danger", message: data.error.message });
        } else {
          console.log(requestOptions.body);
          setAlert({
            type: "alert-success",
            message: "movie added to database successfully.",
          });
          navigate("/movies");
        }
      });
  }

  function handleDelete() {
    const requestOptions = {
      method: "DELETE",
    };
    fetch(`http://localhost:8080/admin/movie/${id}/delete`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          setAlert({ type: "alert-danger", message: data.error.message });
        } else {
          setAlert({
            type: "alert-success",
            message: "movie added to database successfully.",
          });
        }
      });
  }

  return (
    <>
      <h2>Add or Edit Movie</h2>
      {}
      <Alert alertType={alert.type} alertMessage={alert.message} />
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
          type="date"
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
        <label htmlFor="mpaa_rating" className="form-label">
          MPAA RAting
        </label>
        <select
          className="form-select"
          name="mpaa_rating"
          id="mpaa_rating"
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
          Rating
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
        {props.do == "edit" && (
          <Link to="/movies" className="btn btn-danger" onClick={handleDelete}>
            Delete
          </Link>
        )}
      </form>
    </>
  );
}

export default AddOrEditMovie;
