FROM alpine:latest

COPY bird.bin . 

COPY templates ./templates/
COPY static ./static/

CMD ["./bird.bin"]
