package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MechPosition struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

// App struct
type App struct {
	ctx  context.Context
	sock net.Listener
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func tickServer(ctx context.Context, c net.Conn) {
	// I'm not sure if this makes sense but it seems like we should handle closing the connection eventually
	defer c.Close()

	// User has connected so we can begin sending the game state to them
	// After we have sent the game state we will wait for the user to respond
	// if the user doesn't respond within X seconds (maybe 1 second)
	// then the user will be skipped and any response they send will be ignored
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	// Listen for any responses from the user
	go func() {
		s := bufio.NewScanner(c)
		for s.Scan() {
			userCmd := s.Text()
			fmt.Printf("User CMD: %s\n", userCmd)

			cmdParts := strings.Split(userCmd, " ")
			x, err := strconv.ParseInt(cmdParts[0], 10, 0)
			if err != nil {
				panic(err)
			}

			y, err := strconv.ParseInt(cmdParts[1], 10, 0)
			if err != nil {
				panic(err)
			}

			mp := MechPosition{X: x, Y: y}
			runtime.EventsEmit(ctx, "game-tick", mp)
			fmt.Printf("Mech Position %v\n", mp)
		}
	}()

	// Start sending the user tick data every second
	// This is too rudimentary but will eventually send the user tick data as soon as they respond (or something like that)
	for {
		c.Write([]byte("tick\n"))
		time.Sleep(1 * time.Second)
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	socketPath := os.TempDir() + "/mech-commander.sock"
	if err := os.RemoveAll(socketPath); err != nil {
		panic(err)
	}

	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	// socket is saved so it can be closed when the application is shutting down
	a.sock = sock

	go func() {
		for {
			// Accept new connections
			fmt.Println("Waiting for new connection")
			conn, err := sock.Accept()
			if err != nil {
				panic(err)
			}

			go tickServer(ctx, conn)
		}
	}()
}

func (a *App) shutdown(ctx context.Context) {
	err := a.sock.Close()
	if err != nil {
		panic(err)
	}

	socketPath := os.TempDir() + "/mech-commander.sock"
	if err = os.RemoveAll(socketPath); err != nil {
		panic(err)
	}
}
