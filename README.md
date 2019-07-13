# go-slack-app
A basic slack app.

----
#### Run

To run with go add your Slack API token to the `.env` file and then run:
```
$ go run testbot.go
```

----
#### As an executable

To run as an executable first build the project:
```
$ go build
```
then run the executable:
```
$ go-slack-app
```

----
#### Dockerised

To run dockerised first build the project:
```
$ GOOS=linux GOARCH=amd64 go build
```
then create the container image:
```
$ docker build -t testbot .
```
and run with the docker daemon:
```
$ docker run -d testbot
```
