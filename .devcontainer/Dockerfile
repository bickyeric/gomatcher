FROM golang:1.23

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install go.uber.org/mock/mockgen@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest
