# shortie
![example workflow](https://github.com/mridulganga/shortie/actions/workflows/docker-image.yml/badge.svg)
short url shortner

### About
This is a small URL shortner (60 lines of code). It uses leveldb and fiber(based on fasthttp). Why you say? I just wanted to try them out.
Anyway, there are only 3 things to do
- add a new shortened url
- get list of entries in db
- visit the shortened url to be redirected to the actual url

### To run it -
```
go run main.go
```

### To add an entry
```
curl -X POST 'https://shortie.junglesucks.com/' \
-H 'Content-Type: application/json' \
-d '{
    "code": "mg",
    "url": "https://mridulganga.dev"
}'
```
now to access this you can goto https://shortie.junglesucks.com/mg

### To list the entries
```
curl -X GET https://shortie.junglesucks.com?list=true
```
oh also, there is no delete feature, you can however overwrite previous entries.

### What more?
There is a nice Dockerfile which can be used to containerize this, use a volume for urls.db folder to make it persistent.
