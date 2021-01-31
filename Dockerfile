FROM golang:1.15.7-buster
 ARG AWS_ACCESS_KEY_ID
 ARG AWS_SECRET_ACCESS_KEY
 RUN go get -u github.com/beego/bee
 ENV GO111MODULE=on
 #ENV GO111MODULE=off
 ENV GOFLAGS=-mod=vendor
 ENV APP_USER app
 ENV APP_HOME /go/src/main
 ARG GROUP_ID
 ARG USER_ID
 RUN groupadd --gid 500 app && useradd -m -l --uid $USER_ID --gid 500 $APP_USER
 RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
 USER $APP_USER
 WORKDIR $APP_HOME
 EXPOSE 8010
 CMD ["bee", "run"]
