FROM golang:1.18
WORKDIR /build

COPY main.go .
COPY go.mod .
COPY go.sum .


RUN mkdir rego
COPY ./rego/get_pic_cat.go /build/rego
COPY ./rego/get_pic_dog.go /build/rego
COPY ./rego/homepage.go /build/rego
COPY ./rego/random.go /build/rego

RUN mkdir /build/rego/resources


RUN mkdir /build/rego/resources/html
COPY ./rego/resources/html/client.html  /build/rego/resources/html


RUN mkdir /build/rego/resources/pictures

RUN mkdir /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/1.jpg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/2.jpg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/3.jpg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/4.jpeg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/5.jpeg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/6.jpeg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/7.jpg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/8.jpeg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/9.jpg /build/rego/resources/pictures/cats
COPY ./rego/resources/pictures/cats/10.jpg /build/rego/resources/pictures/cats


RUN mkdir /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/1.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/2.jpeg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/3.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/4.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/5.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/6.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/7.jpeg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/8.jpg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/9.jpeg /build/rego/resources/pictures/dogs
COPY ./rego/resources/pictures/dogs/10.jpg /build/rego/resources/pictures/dogs


RUN go mod download
RUN go mod tidy
RUN GOOS=linux go build -pkgdir /build/rest_go ServerTestDocker
EXPOSE 8080
CMD [ "./ServerTestDocker" ]