package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/distribution"
    "github.com/labstack/echo/v4"
)

type node struct{}

var Node = new(node)

// Nodes
// @ID           NodeNodes
// @Router       /nodes [GET]
// @Summary      Getting current online nodes.
// @Tags         Node
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response{data=[]string}  "Node api urls"
func (*node) Nodes(c echo.Context) (err error) {
    ctx := app.Context(c)
    return ctx.SendResponse(distribution.Hub.OnlineNodes())
}
