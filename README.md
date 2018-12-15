# Fizzbuzz Rest API

## What this project do
This project is a REST API endpoint that accepts five parameters : two strings (say, string1 and string2), and three integers (say, int1, int2 and limit), and returns a JSON.

It must return a list of strings with numbers from 1 to limit, where:
all multiples of int1 are replaced by string1,
all multiples of int2 are replaced by string2,
all multiples of int1 and int2 are replaced by string1string2

## Build the APP
There are makefile to manage the build of this app. 
If you want to test the app use : 
```shell
make
./fizzbuzz
```
The App start by default on 8080 port.

After that you can call the API : 
```shell 
$ curl http://localhost:8080/v1/fizzbuzz?string1=fizz&string2=buzz&int1=3&int2=5&limit=15
```
return :
```json
{
   "status":200,
   "value":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
}
```

## Deploy the APP
The best way to deploy this app is with docker you can use this command to build the docker image and test it.
```shell
make docker-build
```

It use the distroless build of the ```golang:1.11``` version of the image to be lightweight and secure.

See more about distroless : https://github.com/GoogleContainerTools/distroless

## Other make command
 - Download dependencies : ```make deps```
 - Execute unit tests : ```make test```
 - Lint your code : ```make lint```
 - See code coverage : ```make coverage```
 - Clean project : ```make clean```