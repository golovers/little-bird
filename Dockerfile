FROM scratch

WORKDIR /app
COPY bird.bin /app

COPY templates /app/templates/
COPY static /app/static/

CMD ["/app/bird.bin"]