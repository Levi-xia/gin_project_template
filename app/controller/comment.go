package controller

import (
	"github.com/gin-gonic/gin"
	"project/app/dto"
	"project/app/model"
	"project/app/service"
	"project/common"
	"strconv"
	"time"
)

func GetCommentById(c *gin.Context) {

	commentId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	req := &dto.GetCommentReq{
		CommentId: commentId,
	}

	if check := common.GetErrorMsg(req); check != "" {
		common.WithError(c, common.PARAM_ERROR, nil, common.Options{
			ShowDetail:  true,
			FullMessage: check,
		})
		return
	}
	comment, err := service.GetCommentById(c, commentId)

	if err != nil {
		common.WithError(c, common.PARAM_ERROR, nil, common.Options{
			ShowDetail:  true,
			FullMessage: err.Error(),
		})
		return
	}

	resp := &dto.GetCommentResp{
		CommentId: comment.CommentId,
		Content:   comment.Content,
	}

	common.Success(c, resp)
}

func AddComment(c *gin.Context) {

	req := dto.AddCommentReq{}

	_ = c.ShouldBind(&req)

	if check := common.GetErrorMsg(req); check != "" {
		common.WithError(c, common.DB_QUERY_FAILED, nil, common.Options{
			ShowDetail:  true,
			FullMessage: check,
		})
		return
	}
	comment := model.Comment{
		UserId:     req.UserId,
		TopicId:    req.TopicId,
		TopicType:  req.TopicType,
		TextStruct: req.TextStruct,
		ReplyId:    req.ReplyId,
		Content:    req.Content,
		Owner:      req.Owner,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	commentId, err := model.AddComment(comment)

	if err != nil {
		common.WithError(c, common.DB_QUERY_FAILED, nil, common.Options{
			ShowDetail:  true,
			FullMessage: err.Error(),
		})
		return
	}

	resp := dto.AddCommentResp{
		CommentId: commentId,
	}
	common.Success(c, resp)
	return
}

// 测试写入token
//token, detail, err := common.CreateToken("10209001", common.AppGuardName)
//
//c.SetCookie("ATK", token.AccessToken, token.ExpiresIn, "/", c.Request.Host, false, false)
//
//common.Success(c, map[string]any{
//	"token":  token,
//	"detail": detail,
//	"e":      err,
//})
//return
