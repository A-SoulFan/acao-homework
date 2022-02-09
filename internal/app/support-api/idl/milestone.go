package idl

import (
	"context"
)

type MilestoneTask interface {
	InitTask(ctx context.Context)
}

type MilestoneService interface {
	NextGroup(ctx context.Context, req NextGroupReq) (*PaginationList, error)
}

type NextGroupReq struct {
	NextKey uint `form:"next_key,default=0" binding:"omitempty,numeric,gte=0"`
	Size    uint `form:"size,default=50" binding:"omitempty,numeric,gt=0,lt=100"`
}

type NextGroupReply struct {
	Title     string `json:"title"`
	Subtitled string `json:"subtitled"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	TargetUrl string `json:"target_url"`
	Timestamp uint   `json:"timestamp"`
}

type PaginationList struct {
	List    interface{} `json:"list"`
	NextKey interface{} `json:"next_key"`
}
