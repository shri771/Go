FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ENV DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable \
    PLATFORM="dev" \
    SECRET="uPOO/GpzWa9D5poJFBqTKbVOoDr4lKthPWVd2NZJQsGmHu9R3xSRtx2YX7f30lKXTHT5DwTpqMIMUp+zC/4WEw==" \
    POLKA_KEY="f271c81ff7084ee5b99a5091b42d486e"

CMD ["/bin/chirps"]
