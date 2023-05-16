import React from "react";
import { useSearchParams } from "react-router-dom";
import cssParam from "./styles/searchTab.module.css";
import Head from "./header";
import BookSheet from "./components/BookSheet";

import img1 from "./image/img1.png";

function SearchTab() {
  const [searchParams] = useSearchParams();
  const question = searchParams.get("q");
  return (
    <>
      <Head />
      <a className={cssParam.resultSearch}>Результаты поиска «{question}»</a>

      <div className={cssParam.bookOrder}>
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
        <BookSheet
          info={{
            id: 1,
            srcImg: img1,
            author: "ДЖ.К Роулинг",
            name: "Гарри Поттер и орден феникса",
          }}
        />
      </div>
    </>
  );
}

export default SearchTab;
