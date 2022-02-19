package handler

import (
	"net/http"

	"github.com/A-SoulFan/acao-homework/internal/app/support_api/idl"
	"github.com/A-SoulFan/acao-homework/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type Member struct {
	memberService idl.MemberService
}

func NewMemberApi(memberService idl.MemberService) *Member {
	return &Member{
		memberService: memberService,
	}
}

func (h *Member) MemberServiceGetAllMembers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if resp, err := h.memberService.GetAllMembers(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *Member) MemberServiceGetMemberExperience() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.MemberExperienceReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		if resp, err := h.memberService.GetMemberExperience(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *Member) MemberServiceGetMemberVideos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.MemberVideoReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		if resp, err := h.memberService.GetMemberVideos(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func (h *Member) MemberServiceGetMemberDebts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req idl.MemberDebtReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		resp := h.memberService.GetMemberDebts(ctx, req)
		ctx.JSON(http.StatusOK, SuccessResponse(resp))
	}
}
