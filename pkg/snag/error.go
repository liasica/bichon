package snag

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data"`
}

func (e Error) Error() string {
    return e.Message
}
