FROM golang as builder
WORKDIR /work
COPY . .
RUN go build -o main


FROM ubuntu
WORKDIR /work
COPY --from=builder /work/main /work/main
RUN chmod +x main
CMD /work/main