FROM golang:alpine

# set env
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# move：/build
WORKDIR /build

# copy
COPY . .

# build
RUN go build -o app .

#  move to dist 
WORKDIR /dist

# copy 
RUN cp /build/app .

# expose
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/dist/app"]