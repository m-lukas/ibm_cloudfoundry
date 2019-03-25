FROM golang:latest

WORKDIR $GOPATH/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV DB_HOST manny.db.elephantsql.com
ENV DB_PORT 5432
ENV DB_NAME jofcwcfq
ENV DB_USER jofcwcfq
ENV DB_PASS h7ufiWtJ-_diOHpyupzY8dO48Y8aV2vT

EXPOSE 5000

CMD ["app"]