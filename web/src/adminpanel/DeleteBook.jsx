import React, {useState} from 'react';
import Modal from "../components/Modal";
import CatalogAPI from "../http/catalog";

const DeleteBook = () => {
    const [id, setId] = useState();
    const [bookAddActive, setBookAddActive] = useState(false);
    return (
        <div>
            <button className={"admin__btn"} onClick={(e) => { setBookAddActive(true); }} > Удалить книгу</button>
            <Modal active={bookAddActive} setActive={setBookAddActive}>
                <h1>Добавление новой книги</h1>
                <input value={id} onChange={(e) =>{setId(e.target.value)}} placeholder="Введите id"/>
                <button onClick={() => {CatalogAPI.deleteBook(id).catch(e => console.error(e))}}>Удалить книгу</button>
            </Modal>
        </div>
    );
};

export default DeleteBook;
