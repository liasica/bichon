package distribution

import (
    "encoding/binary"
    "errors"
    "github.com/panjf2000/gnet/v2"
    "io"
)

const (
    bodySize = 4
)

var errIncompletePacket = errors.New("incomplete packet")

func (h *hub) decode(c gnet.Conn) ([]byte, error) {
    buf, _ := c.Peek(bodySize)
    if len(buf) < bodySize {
        return nil, errIncompletePacket
    }

    bodyLen := binary.BigEndian.Uint32(buf[:bodySize])
    msgLen := bodySize + int(bodyLen)
    if c.InboundBuffered() < msgLen {
        return nil, errIncompletePacket
    }
    buf, _ = c.Peek(msgLen)
    _, _ = c.Discard(msgLen)

    return buf[bodySize:msgLen], nil
}

func (c *client) decode() (message []byte, err error) {
    buf := make([]byte, bodySize)
    _, err = io.ReadFull(c, buf)
    if err != nil {
        return
    }

    bodyLen := binary.BigEndian.Uint32(buf)

    message = make([]byte, int(bodyLen))
    _, err = io.ReadFull(c, message)

    return message, err
}

func pack(message []byte) []byte {
    msgLen := bodySize + len(message)

    data := make([]byte, msgLen)

    binary.BigEndian.PutUint32(data[:bodySize], uint32(len(message)))
    copy(data[bodySize:msgLen], message)

    return data
}
