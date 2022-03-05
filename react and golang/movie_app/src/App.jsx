import React from "react";
import { BrowserRouter as Router, Link, Route, Routes } from "react-router-dom";
import Grid from "@material-ui/core/Grid";
import "./App.css";
import Box from "@mui/material/Box";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import FormatListBulletedIcon from "@mui/icons-material/FormatListBulleted";
import HomeIcon from "@mui/icons-material/Home";
import LocalMoviesIcon from "@mui/icons-material/LocalMovies";
import CategoryIcon from "@mui/icons-material/Category";
import Divider from "@mui/material/Divider";
import Movies from "./components/Movies";
import Home from "./components/Home";
import Admin from "./components/Admin";
import MoviePage from "./components/MoviePage";
import Categories from "./components/Categories";
import CategoryPage from "./components/CategoryPage";

function App() {
  return (
    <>
      <h1 float="left" className="site-header">
        Go Movie app
      </h1>
      <Divider />
      <Router>
        <Grid container columnSpacing={1} className="page-container">
          <Grid item className="list-container" xs={3}>
            <List>
              <ListItem disablePadding component={Link} to="/">
                <ListItemButton>
                  <ListItemIcon>
                    <HomeIcon />
                  </ListItemIcon>
                  <ListItemText primary="Home" />
                </ListItemButton>
              </ListItem>
              <Divider />
              <ListItem disablePadding component={Link} to="/movies">
                <ListItemButton>
                  <ListItemIcon>
                    <FormatListBulletedIcon />
                  </ListItemIcon>
                  <ListItemText primary="Movie List" />
                </ListItemButton>
              </ListItem>
              <Divider />
              <ListItem disablePadding component={Link} to="/genres">
                <ListItemButton>
                  <ListItemIcon>
                    <CategoryIcon />
                  </ListItemIcon>
                  <ListItemText primary="Genres" />
                </ListItemButton>
              </ListItem>
              <Divider />
              <ListItem disablePadding component={Link} to="/movie-catalogue">
                <ListItemButton>
                  <ListItemIcon>
                    <LocalMoviesIcon />
                  </ListItemIcon>
                  <ListItemText primary="Movie Catalogue" />
                </ListItemButton>
              </ListItem>
            </List>
          </Grid>

          <Grid item xs={8} className="text-container">
            <Routes>
              <Route path="/movies" element={<Movies />}></Route>
              <Route path="/movie-catalogue" element={<Admin />}></Route>
              <Route path="/movies/:id" element={<MoviePage />}></Route>
              <Route path="/genres" element={<Categories />}></Route>
              <Route
                path="/genres/:genre_name"
                element={<CategoryPage />}
              ></Route>
              <Route path="/" element={<Home />}></Route>
            </Routes>
          </Grid>
        </Grid>
      </Router>
    </>
  );
}

export default App;
