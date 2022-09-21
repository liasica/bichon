package ws

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Serve(hub *hub, c echo.Context) (err error) {
    address := c.Param("address")
    if address == "" {
        return model.ErrWebsocketPath
    }

    var upgrader = websocket.Upgrader{}
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    var conn *websocket.Conn
    conn, err = upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return
    }

    client := NewClient(hub, conn, address)

    go client.readPump()
    go client.writePump()

    return nil
}
