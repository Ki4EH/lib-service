import axios from 'axios';

export const CATALOG_URL = 'http://localhost:5000';

const $catalog = axios.create({
    withCredentials: true,
    baseURL: CATALOG_URL,
});

$catalog.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
    return config;
});


export default class CatalogAPI{
    static async postBook(title, author, isbn, count){
            return $catalog.post("/book?" + new URLSearchParams({
                title:title,
                author:author,
                isbn:isbn,
                count:count
            }));
    }
}
