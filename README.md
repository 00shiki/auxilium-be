## API Spec

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

Status OK 201

```json
{
  "message": "register success"
}
```

### Get Posts

GET `api/v1/posts`

Params

- `page` (default = 1)
- `size` (default = 5)

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "timeline get success",
  "data": {
    "list": [
      {
        "id": 1,
        "username": "username",
        "avatar_url": "lorem ipsum",
        "body": "lorem ipsum",
        "image_url": "lorem ipsum",
        "comment_count": 12,
        "like_count": 12
      }
    ]
  }
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
  "message": "timeline get success",
  "data": {
    "detail": {
      "id": 1,
      "username": "username",
      "avatar_url": "lorem ipsum",
      "body": "lorem ipsum",
      "image_url": "lorem ipsum",
      "comment_count": 12,
      "like_count": 12,
      "comment": [
        {
          "id": 2,
          "username": "username",
          "avatar_url": "lorem ipsum",
          "body": "lorem ipsum",
          "image_url": "lorem ipsum",
          "comment_count": 12,
          "like_count": 12
        }
      ]
    }
  }
}
```

### Create Post

POST `api/v1/posts`

Header

- `Authorization: Bearer <token>`

Content-Type: `multipart/form-data`

Request Body:

- `anonymous` as `boolean`
- `body` as `string`
- `image` as `file`

Response Body:

Status OK 201

```json
{
  "message": "post created successfully"
}
```