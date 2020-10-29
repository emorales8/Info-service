FROM amd64/golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv
COPY get .
EXPOSE 8082

CMD ["/app/get"]
