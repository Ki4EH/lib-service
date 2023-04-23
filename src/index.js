import React from "react";
import ReactDOM from "react-dom/client";
import Head from "./header";
import DashBoard from "./dashboard";
import Adminpage from "./adminpanel/adminpage";

const headeBoard = ReactDOM.createRoot(document.getElementById("headBoard"));
headeBoard.render(<Head />);
const dashBoard = ReactDOM.createRoot(document.getElementById("dashboard"));
dashBoard.render(<Adminpage />);
