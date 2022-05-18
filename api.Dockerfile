FROM --platform=linux/x86_64 douyin-12306-golang

ADD . /workspace
WORKDIR /workspace/cmd/api

RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct"  \
&& go mod tidy

CMD ["./run.sh"]