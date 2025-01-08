package main

import (
	"fmt"
	"log"
	"net"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const LISTEN_PORT = 8080

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", LISTEN_PORT))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on  0.0.0.0:%d\n", LISTEN_PORT)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := "Welcome to the GOLASH Interpreter. We made this to replace SSH because someone got paranoid after CVE-2024-6387. SO WHY NOT JUST MAKE IT OUR OWN. This should be password protected, but I have given up.\n\nUse ### to end the script\n"
	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}

	var cmd []byte
	buf := make([]byte, 1024)

	for {
		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		for i := range buf {
			if buf[i] == '#' {
				if buf[i+1] == '#' && buf[i+2] == '#' {
					fmt.Println("running catcher")
					cmd = append(cmd, buf[0:i]...)
					fmt.Printf("Received string: %s\n", cmd)
					//fmt.Printf("Received hex: %d\n", cmd)
					res := RunTheCode(string(cmd))
					_, err = conn.Write([]byte(res))
					if err != nil {
						log.Fatal(err)
					}
					break
				}
			} else if buf[i] == '\n' {
				cmd = append(cmd, buf[0:i+1]...)
				break

			}

		}

	}

}

func RunTheCode(cmd string) string {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err := i.Eval(cmd)

	if err != nil {
		return fmt.Sprintf("Error: %v\n", err)
	}

	return "Code executed successfully\n"

}
