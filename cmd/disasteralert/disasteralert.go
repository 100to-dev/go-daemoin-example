package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	defer conn.Close()

	obj := conn.Object("com.github.centodev.gonnaraind", "/com/github/centodev/gonnaraind")

	call := obj.Call("com.github.centodev.gonnaraind.DisasterAlert", 0, "disaster comming!!!")

	if call.Err != nil {
		fmt.Fprintf(os.Stderr, "failed to call DisasterAlert function: %s", err)
		os.Exit(1)
	}
}
