FROM golang:1.19

COPY main.go /main.go
ENTRYPOINT ["go", "run", "/main.go"]
