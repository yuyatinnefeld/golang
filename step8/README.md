## Source:
https://www.mitrais.com/news-updates/how-to-dockerize-a-restful-api-with-golang-and-postgres/

```bash
# create initial fiels

# build image
docker build -t my-golang-image .

# run container
docker run -it -p 8080:8080 my-golang-image
```