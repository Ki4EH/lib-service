import React from "react";
import "../styles/bookSheet.css";
import { Link, Route, Routes } from "react-router-dom";

let props_exp;

const BookSheet = (props) => {
  props_exp = props;
  let id = props.info.id;
  return (
    <>
      <Link to={"/book/?" + id} style={{ textDecoration: "none" }}>
        <div className="bookPrev">
          <img
            src={props.info.srcImg}
            alt="bookSheet"
            className="bookImg"
          ></img>
          <a className="bookInfo">
            {props.info.author + ". " + props.info.name}
          </a>
        </div>
      </Link>
    </>
  );
};

export { props_exp };
export default BookSheet;
