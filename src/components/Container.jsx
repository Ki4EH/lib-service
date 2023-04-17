import React from "react";
import "../styles/dashboard.css";
import rigthArrow from "../svgs/arrow_right.svg";
import BookSheet from "./BookSheet.jsx";
import imgBook1 from "../image/img1.png";

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
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
              srcImg: imgBook1,
              author: "ДЖ.К Роулинг",
              name: "Гарри Поттер и орден феникса",
            }}
          />
          <BookSheet
            info={{
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
