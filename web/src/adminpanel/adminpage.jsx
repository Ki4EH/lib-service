import React,{useState} from 'react';
import Modal from "../components/Modal";
import CatalogSerivce from "../http/catalog";

function Adminpage(props) {
    const [bookAddActive, setBookAddActive] = useState(false);
    const [title,setTitle] = useState("");
    const [author,setAuthor] = useState("");
    const [isbn,setIsbn] = useState("");
    const [count,setCount] = useState("");

    function postBook(title,author,isbn,count){
        try{
            const resp = CatalogSerivce.postBook(title,author,isbn,count)
        }catch (e){
            console.error(e)
        }
    }

    return (
        <div>
            <h1>ADMINKA</h1>
            <button onClick={(e) => { setBookAddActive(true); }} > Добавить книгу </button>
            <Modal active={bookAddActive} setActive={setBookAddActive}>
                <h1>Добавление новой книги</h1>
                <input value={title} onChange={(e) =>{setTitle(e.target.value)}} placeholder="Название"/>
                <input  value={author} onChange={(e) =>{setAuthor(e.target.value)}} placeholder="Автор"/>
                <input  value={isbn} onChange={(e) =>{setIsbn(e.target.value)}} placeholder="ISBN"/>
                <input  value={count} onChange={(e) =>{setCount(e.target.value)}} placeholder="Количество"/>
                <button onClick={() => {postBook(title,author,isbn,count)}}>Добавить книгу</button>
                <button onClick={() => {setTitle("");setAuthor("");setIsbn("");setCount("");}}>Отчистить все поля</button>
            </Modal>
        </div>
    );
}

export default Adminpage;