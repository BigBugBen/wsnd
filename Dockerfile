# 编译阶段
FROM golang:alpine

WORKDIR /go

COPY server.go .

RUN go build -o server .
#RUN go build ./src/server.go
#RUN go run server.go
RUN ./server


#FROM ubuntu

#COPY --from=0 /go/server .

#RUN apt-get update
#CMD [ "top" ]
#CMD [ "./server" ]
#RUN ./server
#ENTRYPOINT [ "./server" ]

# git remote add dice http://dice.dev.terminus.io/wb/dice-test/java-demo
# git push -u dice --all
# git push -u dice --tags
