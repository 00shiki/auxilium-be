## API Spec

### Login

POST `/v1/users`

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
  "access_token": "fjda;lkfdjal;dkfja",
  "refresh_token": ";fkdsja;flkdajf;ldkaj"
}
```

### Register

POST `/v1/users/register`

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

### Timeline

GET `/v1/posts`

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
  "list": [
    {
      "id": 1,
      "username": "username",
      "avatar_url": "lorem ipsum",
      "body": "lorem ipsum",
      "comment_count": 12,
      "like_count": 12
    }
  ]
}
```

### Timeline Detail

GET `/v1/posts/:id`

Header

- `Authorization: Bearer <token>`

Response Body:

Status OK 200

```json
{
  "message": "timeline get success",
  "detail": {
    "id": 1,
    "username": "username",
    "avatar_url": "lorem ipsum",
    "description": "lorem ipsum",
    "comment_count": 12,
    "like_count": 12
  },
  "comment": [
    {
      "id": 2,
      "username": "username",
      "avatar_url": "lorem ipsum",
      "description": "lorem ipsum",
      "comment_count": 12,
      "like_count": 12
    }
  ]
}
```