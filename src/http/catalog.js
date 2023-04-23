import axios from 'axios';

export const CATALOG_URL = 'http://localhost:5000';

const $catalog = axios.create({
    withCredentials: true,
    baseURL: CATALOG_URL,
});

export default class CatalogAPI{
    static async postBook(title, author, isbn, count){
        try{
            return $catalog.post("/book?" + new URLSearchParams({
                title:title,
                author:author,
                isbn:isbn,
                count:count
            }));
        }catch (e){
            console.error(e);
        }

    }
}
