## API Spec

### Login

POST `/v1/login`

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
  "token": "fjda;lkfdjal;dkfja",
  "refresh_token": ";fkdsja;flkdajf;ldkaj"
}
```

### Register

POST `/v1/register`

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

GET `/v1/list`

Params

- `page` (default = 1)
- `size` (default = 5)

Header

- `Authorization: Bearer <token>`

Response Body:

```json
{
  "message": "timeline get success",
  "list": {
    "username": "username",
    "description": "lorem ipsum",
    "comment_count": 12,
    "like_count": 12,
    "created_at": "2022-01-08T06:34:18.598Z"
  }
}
```