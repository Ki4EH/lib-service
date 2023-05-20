import React from "react";
import "../styles/dashboard.css";
import rigthArrow from "../svgs/arrow_right.svg";
import BookSheet from "./BookSheet.jsx";
import imgBook1 from "../image/img1.png";

let getComponents = {
  id: 6,
  srcImg: imgBook1,
  author: "ДЖ.К Роулинг",
  name: "Гарри Поттер и орден феникса",
};

const Container = (prop) => {
  return (
    <>
      <div className="container">
        <div className="textContainer">
          <a className="tagText">{prop.info.text}</a>
          <img src={rigthArrow} alt="rightArrow" className="rightArrow1"></img>
        </div>
        <div className="scroller">
          <BookSheet
            info={{
              id: 1,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              id: 2,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              id: 3,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              id: 4,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              id: 5,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              id: 6,
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
        </div>
      </div>
    </>
  );
};

export default Container;
