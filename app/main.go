package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := LoadFromEnv()
	if err != nil {
		exitWithError(err)
	}

	app, err := NewApp(config)
	if err != nil {
		exitWithError(err)
	}

	app.ListenWebSocketEvent()
	defer app.Close()

	select {}
}

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
