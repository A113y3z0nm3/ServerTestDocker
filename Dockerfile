FROM golang:1.18
WORKDIR /build

COPY main.go .
COPY go.mod .
COPY go.sum .


RUN mkdir rest_go
COPY ./rest_go/get_pic_cat.go /build/rest_go
COPY ./rest_go/get_pic_dog.go /build/rest_go
COPY ./rest_go/homepage.go /build/rest_go
COPY ./rest_go/random.go /build/rest_go


RUN mkdir /build/rest_go/resources


RUN mkdir /build/rest_go/resources/html
COPY ./rest_go/resources/html/client.html  /build/rest_go/resources/html


RUN mkdir /build/rest_go/resources/pictures

RUN mkdir /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/1.jpg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/2.jpg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/3.jpg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/4.jpeg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/5.jpeg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/6.jpeg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/7.jpg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/8.jpeg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/9.jpg /build/rest_go/resources/pictures/cats
COPY ./rest_go/resources/pictures/cats/10.jpg /build/rest_go/resources/pictures/cats


RUN mkdir /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/1.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/2.jpeg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/3.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/4.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/5.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/6.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/7.jpeg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/8.jpg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/9.jpeg /build/rest_go/resources/pictures/dogs
COPY ./rest_go/resources/pictures/dogs/10.jpg /build/rest_go/resources/pictures/dogs


RUN go mod download
RUN go mod tidy
RUN GOOS=linux go build -pkgdir /build/rest_go ServerTestDocker
EXPOSE 8000
CMD [ "./server" ]