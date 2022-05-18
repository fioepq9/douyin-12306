FROM --platform=linux/x86_64 douyin-12306-golang

ADD . /workspace
WORKDIR /workspace/cmd/user

RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct"  \
&& go mod tidy

RUN sh build.sh
CMD ["sh", "output/bootstrap.sh"]