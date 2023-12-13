FROM golang:1.19.2-bullseye
 
WORKDIR /app
 
COPY . .
 
RUN go mod download
 
RUN go build -o /eulabs ./framework/cmd
 
EXPOSE 8080
 
CMD [ "/eulabs" ]