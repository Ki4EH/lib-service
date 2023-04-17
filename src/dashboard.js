import React from "react";
import "./styles/dashboard.css";
import Container from "./components/Container.jsx";

function DashBoard() {
  let tg = ["Рекомендации для вас", "Популярное"];
  return (
    <>
      <Container info={{ text: tg[0] }} />
      <Container info={{ text: tg[1] }} />
    </>
  );
}

export default DashBoard;
