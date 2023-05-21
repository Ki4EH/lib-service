import React from 'react';
import "./style/style.css"
import AddBook from "./AddBook";
import DeleteBook from "./DeleteBook";
function Adminpage(props) {

    return (
        <div className="admin__panel">
            <h1>ADMIN PANEL</h1>
            <AddBook/>
            <DeleteBook/>
        </div>
    );
}

export default Adminpage;