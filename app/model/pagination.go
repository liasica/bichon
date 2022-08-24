// Copyright (C) liasica. 2021-present.
//
// Created at 2022/3/1
// Based on aurservd by liasica, magicrolan@qq.com.

package model

import (
    "math"
)

type PaginationReq struct {
    Current  int `json:"current,omitempty" query:"current"`   // current page, start from 1 (default: 1)
    PageSize int `json:"pageSize,omitempty" query:"pageSize"` // count peer page, default is 50
}

func PaginationReqFromPointer(req *PaginationReq) PaginationReq {
    var pq PaginationReq
    if req != nil {
        pq = *req
    }
    return pq
}

type Pagination struct {
    Current int `json:"current"`
    Pages   int `json:"pages"`
    Total   int `json:"total"`
}

type PaginationRes struct {
    Pagination Pagination  `json:"pagination"`
    Items      interface{} `json:"items"`
}

type PaginationQuery interface {
    PaginationItemsX(req PaginationReq) any
    PaginationResult(req PaginationReq) Pagination
}

func ParsePaginationResponse[T any, K any](pq PaginationQuery, req PaginationReq, itemFunc func(item *K) T) (res *PaginationRes) {
    res = new(PaginationRes)
    res.Pagination = pq.PaginationResult(req)

    qr := pq.PaginationItemsX(req)
    items := make([]*K, 0)
    if qr != nil {
        items = qr.([]*K)
    }
    out := make([]T, len(items))
    for i, item := range items {
        out[i] = itemFunc(item)
    }
    res.Items = out
    return
}

func (p PaginationReq) GetCurrent() int {
    c := p.Current
    if c < 1 {
        c = 1
    }
    return c
}

func (p PaginationReq) GetLimit() int {
    limit := p.PageSize
    if p.PageSize < 1 {
        limit = 50
    }
    return limit
}

func (p PaginationReq) GetOffset() int {
    return (p.GetCurrent() - 1) * p.GetLimit()
}

func (p PaginationReq) GetPages(total int) int {
    return int(math.Ceil(float64(total) / float64(p.GetLimit())))
}
