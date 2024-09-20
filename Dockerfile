FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
ENV GOPATH=/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN env GOOS=js GOARCH=wasm go build -o dist/.dist/web-client.wasm ./cmd/client/main.go
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" dist/.dist/wasm_exec.js
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/app ./cmd/server/main.go

FROM scratch
COPY --from=backend-builder /app/dist /server
WORKDIR /server
ENTRYPOINT ["./app"]