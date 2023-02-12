package controller

import (
	"github.com/gin-gonic/gin"
	"project/common"
)

func GetCommentById(c *gin.Context) {

	//commentId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	//
	//dto := &dto.GetCommentReq{
	//	CommentId: commentId,
	//}
	//
	//if check := common.GetErrorMsg(dto); check != "" {
	//	common.WithError(c, common.PARAM_ERROR, nil, common.Options{
	//		ShowDetail:  true,
	//		FullMessage: check,
	//	})
	//	return
	//}
	//
	//comment, _ := service.GetCommentById(c, commentId)

	//common.Success(c, c.GetString("userId"))
	//
	//return

	token, detail, err := common.CreateToken("10209001", common.AppGuardName)

	c.SetCookie("ATK", token.AccessToken, token.ExpiresIn, "/", c.Request.Host, false, false)

	common.Success(c, map[string]any{
		"token":  token,
		"detail": detail,
		"e":      err,
	})
	return
}

func AddComment(c *gin.Context) {
	common.Success(c, c.GetString("userId"))
	return
}
