FROM golang:latest
COPY ./tahor /
CMD ["/tahor"]
EXPOSE 8000