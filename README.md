![go 1.19](https://img.shields.io/badge/golang-1.19-blue)

# 📚 book

북치기 박치기 팀 back-end repository.

책을 읽고 쉽게 기록할 수 있는 서비스를 만듭니다.
https://www.notion.so/b5645e865d6b45e1afe01a465c86d722

## Uses

- uber/fx for DI
- gorm
- release-please (by google)

## How to Use

Uses go 1.19

### Install Packages

```
make install
```

Add `.env` file to the root with the following:

```
PORT=
DATABASE_USER=r
DATABASE_URL=
DATABASE_PASSWORD=
DATABASE_PORT=
DATABASE_NAME=
NAVER_SEARCH_API=
NAVER_CLIENT_ID=
NAVER_CLIENT_SECRET=
API_HOST=
ENV=
LOGGER=
```

### Run

```
make run
```


### Update Open API
(Swagger)
```
make docs
```