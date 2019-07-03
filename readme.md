A Go application created for the purpose of writing a Go application.

#Run

```
docker build -t goapp .
docker run -d -p 8080:8080 --name goapp goapp
```

Go to http://localhost:8080

#Stop

```
docker stop goapp
docker rm goapp
```