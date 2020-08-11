FROM golang:1.13.4
COPY ./tahor /
CMD ["/tahor"]
EXPOSE 8000