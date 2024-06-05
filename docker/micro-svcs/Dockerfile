FROM golang:1.22.3 as builder

# Install protoc, protoc-gen-go, and protoc-gen-go-grpc
RUN apt update && \
    apt install -y protobuf-compiler && \
    protoc --version

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Generate pb
COPY proto ./proto
COPY Makefile .
RUN make proto

# Copy the code into the container
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

# Build the application
RUN go build -o ./out/billing ./cmd/billing
RUN go build -o ./out/notification ./cmd/notification
RUN go build -o ./out/ordering ./cmd/ordering
RUN go build -o ./out/rating ./cmd/rating
RUN go build -o ./out/usermgmt ./cmd/usermgmt

FROM gcr.io/distroless/base-debian12

WORKDIR /dist
COPY --from=builder /build/out .

# Command to run when starting the container
CMD ["echo", "Please specify what to run (billing, notification, ordering, rating, or usermgmt)"]
