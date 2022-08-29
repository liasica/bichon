package ws

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Wrap(hub *Hub, c echo.Context) (err error) {
    ctx, ok := c.(*app.MemberContext)
    if !ok || ctx.Member == nil {
        err = model.ErrAuthRequired
        return
    }
    var upgrader = websocket.Upgrader{}
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    var conn *websocket.Conn
    conn, err = upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
    if err != nil {
        return
    }

    client := &Client{
        mem:  ctx.Member,
        hub:  hub,
        conn: conn,
        // TODO use larger buffer or chunks message
        send: make(chan []byte, 256),
    }

    client.hub.register <- client

    // TODO read and write pump

    return nil
}
