import axios from 'axios';

export const CATALOG_URL = 'http://localhost:8080';

const $catalog = axios.create({
    withCredentials: true,
    baseURL: CATALOG_URL,
});
//    ${localStorage.getItem('token')}
$catalog.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoxfQ.tMnR8SWXirK3X0_MLJA-nHNBkcnVEgOdN0pMkmGcb4U`;
    return config;
});

function genreParse(genres){
    const sep = ",";
    return genres.replace(/\s+/g, sep);
}

export default class CatalogAPI{

    static async postBook(title, author, isbn, count, genres ){
        genres = genreParse(genres);
        return $catalog.post("/book?" + new URLSearchParams({
                title:title,
                author:author,
                isbn:isbn,
                count:count,
                genres:genres,
            }));
    }

    static async deleteBook(id){
        return $catalog.delete("/book?"+ new URLSearchParams({
            id:id,
        }))
    }
}
