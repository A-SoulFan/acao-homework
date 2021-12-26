package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support-api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *defaultSupportAPIhandler) MemberServiceGetAllMembers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.memberService.SetDBwithCtx(ctx)

		if resp, err := h.memberService.GetAllMembers(); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *defaultSupportAPIhandler) MemberServiceGetMemberExperience() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.MemberExperienceReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		h.memberService.SetDBwithCtx(ctx)

		if resp, err := h.memberService.GetMemberExperience(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *defaultSupportAPIhandler) MemberServiceGetMemberVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.MemberVideoReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		h.memberService.SetDBwithCtx(ctx)

		if resp, err := h.memberService.GetMemberVideos(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
