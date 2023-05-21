import React, {useState} from 'react';
import Modal from "../components/Modal";
import CatalogSerivce from "../http/catalog";

const MyComponent = () => {
    const [bookAddActive, setBookAddActive] = useState(false);
    const [title,setTitle] = useState("");
    const [author,setAuthor] = useState("");
    const [isbn,setIsbn] = useState("");
    const [count,setCount] = useState("");
    const [genre, setGenre] = useState("");

    function postBook(title,author,isbn,count, genre){
      CatalogSerivce.postBook(title,author,isbn,count, genre)
           .catch(e => {console.error(e.response.data.message)})
    }

    return (
        <div>
            <button className={"admin__btn"} onClick={(e) => { setBookAddActive(true); }} > Добавить книгу </button>
            <Modal active={bookAddActive} setActive={setBookAddActive}>
                <h1>Добавление новой книги</h1>
                <input value={title} onChange={(e) =>{setTitle(e.target.value)}} placeholder="Название"/>
                <input  value={author} onChange={(e) =>{setAuthor(e.target.value)}} placeholder="Автор"/>
                <input  value={isbn} onChange={(e) =>{setIsbn(e.target.value)}} placeholder="ISBN"/>
                <input  value={count} onChange={(e) =>{setCount(e.target.value)}} placeholder="Количество"/>
                <input  value={genre} onChange={(e) =>{setGenre(e.target.value)}} placeholder="Жанры"/>
                <button onClick={() => {postBook(title,author,isbn,count, genre)}}>Добавить книгу</button>
                <button onClick={() => {setTitle("");setAuthor("");setIsbn("");setCount(""); setGenre("");}}>Отчистить все поля</button>
            </Modal>
        </div>
    );
};

export default MyComponent;
