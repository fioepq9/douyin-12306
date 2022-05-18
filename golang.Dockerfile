FROM --platform=linux/x86_64 golang:1.18

ADD . /temp
WORKDIR /temp


RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct"  \
&& go mod tidy