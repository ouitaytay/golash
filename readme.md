# Golash

A golang interpreter script


Client Example
```
Welcome to the GOLASH Interpreter. We made this to replace SSH because someone got paranoid after CVE-2024-6387. SO WHY NOT JUST MAKE IT OUR OWN. This should be password protected, but I have given up.

Use ### to end the script
package main
import "fmt"
func main(){
fmt.Println("Test")
}
###
Code executed successfully
```

Server Output Example:
```
Listening on  0.0.0.0:8080
Connection from 127.0.0.1:35172
running catcher
Received string: package main
import "fmt"
func main(){
fmt.Println("Test")
}

Test
```


## Config
To change port, change line 12
```
const LISTEN_PORT = 8080
```


## Building

For windows
```
env GOOS=windows GOARCH=amd64 go build -o golash.exe main.go
```

For Linux
```
go build -o golash main.go
```

## Installing Notes
This service will crash alot while users are testing. 

On linux, this should be set as a systemctl service and have restart enabled

On windows, this is not a service binary, so you may want to schedule this as a task under task scheduler.

*NOTE*: This binary is pretty big uncompressed (windows 26mb roughly)
