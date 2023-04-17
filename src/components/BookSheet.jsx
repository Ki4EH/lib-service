import React from "react";
import "../styles/bookSheet.css";

const BookSheet = (props) => {
  return (
    <>
      <div className="bookPrev">
        <img src={props.info.srcImg} alt="bookSheet" className="bookImg"></img>
        <a className="bookInfo">{props.info.author + ". " + props.info.name}</a>
      </div>
    </>
  );
};

export default BookSheet;
