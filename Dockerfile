FROM --platform=linux/x86_64 golang:1.18

ADD . /workspace
WORKDIR /workspace

RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct"  \
&& go mod tidy \
&& go build main.go

CMD ["./main"]