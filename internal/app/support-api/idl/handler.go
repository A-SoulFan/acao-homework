package idl

import (
	"github.com/gin-gonic/gin"
)

type SupportAPIhandler interface {
	MilestoneServiceNextGroup() gin.HandlerFunc
}
