import Head from "./header";
import React from "react";
import img1 from "./image/img1.png";
import styled from "./styles/bookPage.module.css";
import { useState, useEffect } from "react";
import { useSearchParams } from "react-router-dom";

// GET Запрос на каталог

const Page = () => {
  const [getData, setGetData] = useState([]);

  const [postData, setPostData] = useState("");

  const [searchParams] = useSearchParams();
  const bookID = searchParams.get("id");

  useEffect(() => {
    fetchBook();
  }, []);

  const fetchBook = () => {
    fetch("https://jsonplaceholder.typicode.com/posts", { method: "GET" })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setGetData(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  };

  const sendData = () => {
    fetch(`http://10.11.165.211:7000/api/queue/status/${bookID}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: 2, name: "Book1" }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setPostData(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
    console.log("DATA SEND POST");
  };

  const handlePOST = (event) => {
    event.preventDefault();
    sendData();
  };

  let inStock = "Да";
  let n = 1;
  let queue = 0;
  return (
    <div>
      <Head />
      <div className={styled.bookPage}>
        <img src={img1} alt="image1" className={styled.bookImg}></img>
        <div className={styled.infoBlock}>
          <a className={styled.nameBook}>Гарри Поттер и орден феникса</a>
          <a className={styled.author}>Автор: Дж.К Роулинг</a>
          <a className={styled.genre}>Жанр: </a>
          <a className={styled.tags}>Теги: </a>
          <a className={styled.stock}>
            Есть в наличии: <span>{inStock}</span>
          </a>
          <a className={styled.stockCount}>Количество в библиотеке: {n}</a>
          <a className={styled.queue}>В очереди: {queue}</a>
        </div>
        <button className={styled.queueLoan} onClick={handlePOST}>
          <span>Забронировать</span>
        </button>
        <div className={styled.descriptionBlock}>
          <a className={styled.descriptionHeader}>Описание:</a>
          <a className={styled.description}>
            Семья Уизли поддерживает Гарри, пока он ожидает дисциплинарного
            слушания в министерстве магии. Думбльдор возрождает Орден Феникса,
            тайное общество, противостоящее Черному Лорду. Стойкость
            гриффиндорцев непоколебима, как и преданность, что в полной мере
            демонстрирует Сириус Блэк.
          </a>
        </div>
      </div>
    </div>
  );
};

export default Page;
