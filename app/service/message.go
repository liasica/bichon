package service

import (
    "context"
)

type messageService struct {
    ctx context.Context
}

func NewMessage() *messageService {
    return &messageService{
        ctx: context.Background(),
    }
}
