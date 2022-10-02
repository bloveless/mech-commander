package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type MechPosition struct {
	X int
	Y int
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	println("Client got:", string(buf[0:n]))
}

func main() {
	sock, err := net.Dial("unix", os.TempDir()+"/mech-commander.sock")
	if err != nil {
		panic(err)
	}
	defer sock.Close()

	mp := MechPosition{X: 10, Y: 10}

	s := bufio.NewScanner(sock)
	for s.Scan() {
		tickState := s.Text()
		fmt.Printf("tickState: %s\n", tickState)

		mp.X += 5
		mp.Y += 5
		_, err = sock.Write([]byte(fmt.Sprintf("%d %d\n", mp.X, mp.Y)))
		if err != nil {
			log.Fatal("write error:", err)
		}
	}
}
