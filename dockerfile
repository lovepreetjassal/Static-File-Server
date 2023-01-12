FROM golang:latest
COPY ./ /app
workdir /app
cmd ["go", "run" ,"."]
EXPOSE 80