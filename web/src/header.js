import React from "react";
import "./styles/head.css";
import searcher from "./svgs/search.svg";
import logIn from "./svgs/book-user.svg";
import fav from "./svgs/book-bookmark.svg";
import treeL from "./svgs/linesthreeL.svg";
import { useState } from "react";
import { Outlet, Link, Routes, Route, BrowserRouter } from "react-router-dom";

import { useNavigate } from "react-router-dom";

function Head() {
  const navigate = useNavigate();
  const [textSearch, setText] = useState("");

  const handleInputChange = (event) => {
    setText(event.target.value);
  };

  const handleClick = (event) => {
    event.preventDefault();
    navigate("/search/?q=" + textSearch);
    console.log("Click", textSearch);
  };

  const handleEnter = (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
      navigate("/search/?q=" + textSearch);
    }
  }

  return (
    <>
      <header className="headBoard">
        <div className="searcher">
          <a href="/web/public" style={{ textDecoration: "none" }} className="logo">
            Библиотека
            <span className="redPoint">.</span>ru
          </a>
          <input
            className="search"
            type="text"
            placeholder="Книга, автор, жанр"
            onChange={handleInputChange}
            onKeyDown={handleEnter}
          ></input>
          <button className="findBtn" onClick={handleClick}>
            <img className="searchSVG" src={searcher} alt="searchSVG"></img>
            <a className="findText">Поиск</a>{" "}
          </button>
          <div className="blockLogin">
            <img className="logInSVG" src={logIn} alt="loginSVG"></img>
            <a className="logInText">Войти</a>
          </div>
          <div className="fav">
            <img className="favSVG" src={fav} alt="favSVG"></img>
            <a className="favText">Мои книги</a>
          </div>
        </div>
        <div className="rectHeader">
          <div className="podHeader">
            <a className="genres">
              <img className="threeLSVG" src={treeL} alt="threeLSVG"></img>
              Жанры
            </a>
            <a className="new">Новинки</a>
            <a className="popular">Популярное</a>
            <a className="readRand">Что почитать?</a>
          </div>
        </div>
      </header>
    </>
  );
}

export default Head;
