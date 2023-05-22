# Library Catalog Service

The Library Catalog Service is a microservice designed to manage the catalog of books in a library system. It is
responsible for storing and retrieving information about the books, including their title, author, ISBN, genre, and
count (availability).

This service has been built using Go and follows REST principles.

## API Endpoints

The Catalog Service exposes the following RESTful endpoints:

### GET /catalog/book

Get book entries from the catalog.

Need to provide at **least one** of the following query parameters:

| Parameter | Type     | Description                                       |
|:----------|:---------|:--------------------------------------------------|
| `book_id` | `string` | **Optional**. The ID of the book to retrieve.     |
| `title`   | `string` | **Optional**. The title of the book to retrieve.  |
| `author`  | `string` | **Optional**. The author of the book to retrieve. |

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

### POST /catalog/book

Creates a new book entry in the catalog.

| Parameter | Type     | Description                        |
|:----------|:---------|:-----------------------------------|
| `title`   | `string` | **Required**. The title of a book  |
| `author`  | `string` | **Required**. The author of a book |
| `isbn`    | `string` | **Required**. The ISBN of a book   |
| `genres`  | `array`  | **Required**. The genres of a book |
| `count`   | `int`    | **Required**. The count of a book  |

Example:

```http
POST /catalog/book
Content-Type: application/json

{
  "title": "The Book Title",
  "author": "The Book Author",
  "isbn": "123-4567890123",
  "genres": ["Fiction", "Thriller"],
  "count": 5
}
```

### DELETE /catalog/book

Deletes the book entry based on the book id provided as a query parameter.

Example: `DELETE /catalog/book?book_id=22`

## Error Handling

The Catalog Service will return a 400 Bad Request HTTP response code for invalid requests and a 404 Not Found HTTP
response code for non-existing book ids.
