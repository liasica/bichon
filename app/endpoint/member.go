package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/labstack/echo/v4"
)

type member struct{}

var Member = new(member)

// Test
// @ID           MemberTest
// @Router       /memeber/test [GET]
// @Summary      Member Test
// @Tags         Member
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response  "Response success"
func (*member) Test(c echo.Context) (err error) {
    ctx := app.Context(c)

    return ctx.SendResponse()
}
