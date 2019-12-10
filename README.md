# go-library
A simple Library RESTful API built with Golang

## Specification
Built with Golang, served with Echo framework, and GORM ORM to connect with PostgreSQL database.

## Setup
- Make a new file `config.yml` to store your settings with the content from `config.yml.example`.
Or you can rename the `config.yml.example` file.

- Change the `dsn` parameter with the local PostgreSQL connection string on your machine. If you dont have
PostgreSQL running on your machine, you can use the containerized PostgreSQL by executing `docker-compose up`.
The PostgreSQL is exposed on port `54320`

- Run all the migrations SQL files contained in folder 

- Execute `make run` and the server will be served in the port that you have set in your `config.yml` file

## Endpoints
**Also Available in Postman Collection**
### `GET` /books?page={{page}}&size={{size}}
Returning all books contained in DB. Pass the pagination parameter as query parameter

Sample Return:
```
[
  {
    "id": 1575986917501328677,
    "title": "How to Master Javascript in 5 Hours",
    "author": "Myself",
    "publisher": "O'really? Publishing",
    "is_issued": false,
    "issue_count": 2,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null
  },
  {
    "id": 1575999467996150343,
    "title": "How to Master Haskell in 5 Hours",
    "author": "Myself",
    "publisher": "O'really? Publishing",
    "is_issued": true,
    "issue_count": 1,
    "CreatedAt": "2019-12-11T00:37:47.996169Z",
    "DeletedAt": null
  }
]
```

### `GET` /books/{{book_id}}
Find a book by passing its ID.

Sample return:
**If Issued**
```
{
  "book": {
    "id": 1575999467996150343,
    "title": "How to Master Haskell in 5 Hours",
    "author": "Myself",
    "publisher": "O'really? Publishing",
    "is_issued": true,
    "issue_count": 1,
    "CreatedAt": "2019-12-11T00:37:47.996169Z",
    "DeletedAt": null
  },
  "issued_by": {
    "id": 1575987455585862489,
    "name": "Hendrawan",
    "created_at": "2019-12-10T21:17:35.586151Z"
  }
}
```

**If Not Issued**
```
{
  "book": {
    "id": 1575999467996150343,
    "title": "How to Master Haskell in 5 Hours",
    "author": "Myself",
    "publisher": "O'really? Publishing",
    "is_issued": false,
    "issue_count": 1,
    "CreatedAt": "2019-12-11T00:37:47.996169Z",
    "DeletedAt": null
  }
}
```

### `GET` /books/popular
Returning the most issued book

### `POST` /books/create
Create a book by passing title, author, and publisher parameter
Request Body :
```
{
	"title": "{{books_title}}",
	"author": "{{authors_name}}",
	"publisher": "{{books_publisher}}"
}
```

Return:
```
{
  "id": 1575999467996150343, -> `Book's ID`
  "title": "How to Master Haskell in 5 Hours",
  "author": "Myself",
  "publisher": "O'really? Publishing",
  "is_issued": false,
  "issue_count": 0,
  "CreatedAt": "2019-12-11T00:37:47.996169+07:00",
  "DeletedAt": null
}
```

### `POST` /books/issue
Issue a book by passing the books identity, issue details, and issuer
Request body:
```
{
	"issued_book": {{book_id}},
	"start_time": "{{issue_start_time}}" -> `in ISO 8601 format`,
	"finish_time": "{{issue_start_time}} -> `in ISO 8601 format`",
	"issued_by": {{issuer_user_id}}
}
```

Return:
```
{
  "issue": {
    "id": 1576001140135740974, <- `Issue's ID`
    "issued_book": 1575999467996150343,
    "start_time": "2019-10-12T07:20:50.52Z",
    "finish_time": "2019-10-12T07:20:50.52Z",
    "issued_by": 1575987455585862489
  },
  "issued_by": {
    "id": 1575987455585862489,
    "name": "Hendrawan",
    "created_at": "2019-12-10T21:17:35.586151Z"
  }
}
```

### `GET` /books/unissue/{{issue_id}}
Unissue a book

### `DELETE` /books/{{book_id}}
Soft deleting a book without deleting related issue

### `POST` /users/create
Create new user

Request body:
```
{
	"name":"{{users_name}}"
}
```

Return :
```
{
  "id": 1575987455585862489, <- `User's ID`
  "name": "Hendrawan",
  "created_at": "2019-12-10T21:17:35.586151+07:00"
}
```