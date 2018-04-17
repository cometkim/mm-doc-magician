package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

type App struct {
	UserId   string
	Client   *model.Client4
	WsClient *model.WebSocketClient
}

func NewApp(config *ClientConfig) (*App, error) {
	app := &App{}

	if err := app.initClient(config); err != nil {
		return nil, err
	}
	fmt.Println("APIv4 Client is successfully initialized")

	if err := app.initWsClient(config); err != nil {
		return nil, err
	}
	fmt.Println("APIv4 WebSocket Client is successfully initialized")

	return app, nil
}

func (app *App) ListenWebSocketEvent() {
	fmt.Println("Start listening on WebSocket events")

	app.WsClient.Listen()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			select {
			case event := <-app.WsClient.EventChannel:
				handleWebSocketEvent(event)
			case <-c:
				app.Close()
			}
		}
	}()
}

func (app *App) Close() {
	app.WsClient.Close()
}

func (app *App) initClient(config *ClientConfig) error {
	app.Client = model.NewAPIv4Client(config.MattermostBaseURL)
	if len(config.PersonalAccessToken) > 0 {
		app.Client.AuthToken = config.PersonalAccessToken
		user, res := app.Client.GetMe("")
		if res.Error != nil {
			return res.Error
		}
		app.UserId = user.Id
	} else {
		user, res := app.Client.Login(config.Username, config.Password)
		if res.Error != nil {
			return res.Error
		}
		app.UserId = user.Id
	}
	return nil
}

func (app *App) initWsClient(config *ClientConfig) error {
	wsURL := strings.Replace(config.MattermostBaseURL, "http", "ws", -1)
	wsClient, err := model.NewWebSocketClient4(wsURL, app.Client.AuthToken)
	if err != nil {
		return err
	}
	app.WsClient = wsClient
	return nil
}

func handleWebSocketEvent(event *model.WebSocketEvent) {
	switch event.Event {
	case model.WEBSOCKET_EVENT_POSTED:
	}
}
