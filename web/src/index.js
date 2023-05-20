import React from "react";
import ReactDOM from "react-dom/client";
import Head from "./header";
import DashBoard from "./dashboard";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import SearchTab from "./serchTab.js";

import BookPage from "./bookpage.js";

import AuthPage from "./regPage.js";
import LoginPage from "./loginPage";

const App = ReactDOM.createRoot(document.getElementById("root"));
App.render(
  <>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={[<Head />, <DashBoard />]}></Route>
        <Route path="/book/" element={<BookPage />}></Route>
        <Route path="/search/" element={<SearchTab />}></Route>
        <Route path="/auth/" element={<AuthPage />}></Route>
        <Route path="/login/" element={<LoginPage />}></Route>
      </Routes>
    </BrowserRouter>
  </>
);
