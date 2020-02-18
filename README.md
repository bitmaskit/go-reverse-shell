# go-reverse-shell
Simple reverse shell written in GO

### Usage
```shell script
$ go run remote.go
```
or
```shell script
$ go run remote.go -addr=:1337
```
or
```shell script
$ go run remote.go -addr=0.0.0.0:1337
```

Then you can use netcat to connect remotely
```shell script
$ nc 0.0.0.0 1337 
echo hi
```
should return back 'hi'