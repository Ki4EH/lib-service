import React from "react";
import "./styles/head.css";
import searcher from "./svgs/search.svg";
import logIn from "./svgs/book-user.svg";
import fav from "./svgs/book-bookmark.svg";
import treeL from "./svgs/linesthreeL.svg";

// When the user scrolls the page, execute myFunction
window.onscroll = function () {
  myFunction();
};

// Get the header
var header = document.getElementById("headBoard");

// Get the offset position of the navbar
var sticky = header.offsetTop;

// Add the sticky class to the header when you reach its scroll position. Remove "sticky" when you leave the scroll position
function myFunction() {
  if (window.pageYOffset > sticky) {
    header.classList.add("sticky");
  } else {
    header.classList.remove("sticky");
  }
}

function Head() {
  return (
    <>
      <div className="searcher">
        <a className="logo">
          Библиотека
          <span className="redPoint">.</span>ru
        </a>
        <input
          className="search"
          type="text"
          placeholder="Книга, автор, жанр"
        ></input>
        <button className="findBtn">
          <img className="searchSVG" src={searcher} alt="searchSVG"></img>
          <a className="findText">Поиск</a>
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
    </>
  );
}

export default Head;
