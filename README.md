# Auxilium Backend Service

`go 1.20`

## Get Dependencies

```shell
go mod vendor
```

## Running

```shell
go run main.go
```

## Build

```shell
go build
```

# API Spec

---

## Users

### Login

POST `api/v1/users`

Request Body:

```json
{
  "email": "example@mail.com",
  "password": "contohpassword"
}
```

Response Body:

Status OK 200

```json
{
  "message": "login success",
  "data": {
    "access_token": "fjda;lkfdjal;dkfja",
    "refresh_token": ";fkdsja;flkdajf;ldkaj"
  }
}
```

### Register

POST `api/v1/users/register`

Request Body:

```json
{
  "username": "username",
  "first_name": "nama depan",
  "last_name": "nama belakang",
  "email": "example@mail.com",
  "password": "contohpassword",
  "confirm_password": "contohpassword",
  "phone_number": "0812345678"
}
```

Response Body:

Status Created 201

```json
{
  "message": "register success"
}
```

### Detail User

GET `api/v1/users/:username`

Response Body:

Status OK 200

```json
{
  "message": "detail get success",
  "data": {
    "username": "test123",
    "first_name": "Coba",
    "last_name": "Mencoba",
    "email": "coba@gmail.com",
    "phone_number": "62111122222",
    "avatar_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png",
    "posts": [
      {
        "id": 1,
        "username": "test123",
        "avatar_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png",
        "body": "lorem ipsum",
        "image_url": "",
        "comments_count": 0,
        "likes_count": 0
      }
    ]
  }
}
```

### Update User

POST `api/v1/users/update`

Header

- `Authorization: Bearer <token>`

Request Body:

```json
{
    "first_name": "Coba",
    "last_name": "Mencoba",
    "email": "coba@gmail.com",
    "phone_number": "62111122222",
    "avatar_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png",
    "bio": "orang ganteng"
}
```

Response Body:

Status OK 200

```json
{
    "message": "update success"
}
```

---

## Posts

### Get Posts

GET `api/v1/posts`

Params

- `page` (default = 1)
- `size` (default = 10)

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "timeline get success",
  "data": [
    {
      "id": 1,
      "username": "username",
      "avatar_url": "lorem ipsum",
      "body": "lorem ipsum",
      "image_url": "lorem ipsum",
      "comments_count": 12, 
      "likes_count": 12
    }
  ]
}
```

### Post Detail

GET `api/v1/posts/:id`

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "detail get success",
  "data": {
    "id": 8,
    "username": "test123",
    "avatar_url": "",
    "body": "lorem ipsum",
    "image_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png",
    "comments_count": 1,
    "likes_count": 1,
    "comments": [
      {
        "id": 3,
        "username": "test123",
        "avatar_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png",
        "body": "lorem ipsum",
        "likes_count": 0
      }
    ]
  }
}
```

### Create Post

POST `api/v1/posts`

Header

- `Authorization: Bearer <token>`

Request Body:

```json
{
    "anonymouos": false,
    "body": "lorem ipsum",
    "image_url": "https://storage.googleapis.com/auxilium-bucket/test-image/cando54.png"
}
```

Response Body:

Status Created 201

```json
{
  "message": "post created successfully"
}
```

### Like Post

POST `api/v1/posts/:postID/like`

Header

- `Authorization: Bearer <token>`
  Response Body:

Status OK 200

```json
{
  "message": "like post success"
}
```

### Dislike Post

POST `api/v1/posts/:postID/dislike`

Header

- `Authorization: Bearer <token>`
  Response Body:

Status OK 200

```json
{
  "message": "dislike post success"
}
```

### Create Comment

POST `api/v1/posts/:postID/comment`

Header

- `Authorization: Bearer <token>`

Request Body:

```json
{
    "anonymouos": false,
    "body": "lorem ipsum"
}
```

Response Body:

Status Created 201

```json
{
  "message": "comment created successfully"
}
```

### Like Comment

POST `api/v1/posts/:postID/comment/:commentID/like`

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "like comment success"
}
```

### Dislike Comment

POST `api/v1/posts/:postID/comment/:commentID/dislike`

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "dislike comment success"
}
```

---

## Helper

### Create Helper

POST `api/v1/helper`

Header

- `Authorization: Bearer <token>`

Request Body:

```json
{
    "lat": 0.5,
    "lon": 1.0
}
```

Response Body:

Status Created 201

```json
{
  "message": "helper created successfully"
}
```

### Remove Helper

POST `api/v1/helper/remove`

Header

- `Authorization: Bearer <token>`

Response Body:

Status Created 200

```json
{
  "message": "helper removed successfully"
}
```

### Get Helpers

GET `api/v1/helper`

Params

- `lat` float64
- `lon` float64
- `radius` float64

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
    "message": "list helper success",
    "data": [
        {
            "id": 1,
            "username": "test123",
            "avatar_url": "",
            "lat": 0.5,
            "lon": 1
        }
    ]
}
```

---

### Upload Image

POST `api/v1/upload`

Header

- `Authorization: Bearer <token>`

Request Body:

- `image` as `file` max size 10MB

Response Body:

Status OK 200

```json
{
  "message": "upload image success",
  "data": "url_to_image"
}
```