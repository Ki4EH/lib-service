import React from "react";
import "./styles/App.css";
import searcher from "./svgs/search.svg";
import logIn from "./svgs/book-user.svg";
import fav from "./svgs/book-bookmark.svg";

window.onscroll = function () {
  myFunction();
};

// Get the header
var header = document.getElementById("headBoard");

var sticky = header.offsetTop;

function myFunction() {
  if (window.pageYOffset > sticky) {
    header.classList.add("sticky");
  } else {
    header.classList.remove("sticky");
  }
}

function App() {
  return (
    <header>
      <a className="logo">
        Библиотека
        <span className="redPoint">.</span>ru
      </a>
      <input
        className="search"
        placeholder="Книга, автор, жанр"
        type="text"
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
      <div className="podHeader"></div>
    </header>
  );
}

export default App;
