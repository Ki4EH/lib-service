## API Reference

#### Get all items

```http
  GET /catalog/book?book_id=22
```

| Parameter | Type     | Description                                   |
|:----------|:---------|:----------------------------------------------|
| `book_id` | `string` | **Required**. The ID of the book to retrieve. |

### Response:

```json
{
  "id": 22,
  "title": "Carrie",
  "author": "Stephen King",
  "isbn": "978-5-17-088071-3",
  "count": 10,
  "genres": [
    "string",
    "string"
  ]
}
```

Create a new item

```http
  POST /catalog/book
```

| Parameter | Type     | Description                        |
|:----------|:---------|:-----------------------------------|
| `title`   | `string` | **Required**. The title of a book  |
| `author`  | `string` | **Required**. The author of a book |
| `isbn`    | `string` | **Required**. The ISBN of a book   |
| `genres`  | `array`  | **Required**. The genres of a book |
| `count`   | `int`    | **Required**. The count of a book  |

Delete an item

```http
  DELETE /catalog/book?book_id=22
```