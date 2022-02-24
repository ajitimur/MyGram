package http

import (
	"MyGram/internal/domain"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type commentHandler struct {
	beego.Controller
	CommentUseCase domain.CommentUseCase
}

func NewCommentHandler(usecase domain.CommentUseCase) {
	pHandler := &commentHandler{
		CommentUseCase: usecase,
	}

	beego.Router("/comments", pHandler, "get:GetComments")
	beego.Router("/comments", pHandler, "post:StoreComment")
	beego.Router("/comments/:commentId", pHandler, "put:UpdateComment")
	beego.Router("/comments/:commentId", pHandler, "delete:DeleteComment")
}

func (uh commentHandler) StoreComment() {
	//Get UserId
	currentUser := _context.Get(uh.Ctx.Request, "currentUser")

	var marshalled map[string]interface{}

	data, err := json.Marshal(currentUser)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}
	err = json.Unmarshal(data, &marshalled)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	IdMarshal := marshalled["Id"]
	Idfloat := IdMarshal.(float64)
	Idint := int(Idfloat)

	//Get Body
	var body domain.CommentRequest

	err = json.NewDecoder(uh.Ctx.Request.Body).Decode(&body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	//Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&body)
	if err != nil {
		// handle error
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}
	if !b {
		// validation does not pass
		var message domain.ErrorMessage
		message.Code = 400
		message.Message = ""
		for i, err := range valid.Errors {

			if i == 0 {
				message.Message = err.Message
			} else {
				message.Message += ", " + err.Message
			}
		}
		uh.Ctx.Output.SetStatus(400)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	body.UserId = Idint

	//store comment
	res, err := uh.CommentUseCase.SaveComment(uh.Ctx.Request.Context(), body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	uh.Ctx.Output.SetStatus(201)
	uh.Ctx.Output.JSON(res, true, true)
	return
}

func (uh commentHandler) GetComments() {
	//Get UserId
	currentUser := _context.Get(uh.Ctx.Request, "currentUser")

	var marshalled map[string]interface{}

	data, err := json.Marshal(currentUser)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}
	err = json.Unmarshal(data, &marshalled)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	IdMarshal := marshalled["Id"]
	Idfloat := IdMarshal.(float64)
	Idint := int(Idfloat)

	//Get Comments
	res, err := uh.CommentUseCase.GetComments(uh.Ctx.Request.Context(), Idint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(res, true, true)
	return
}

func (uh commentHandler) UpdateComment() {
	//req body
	var body domain.CommentRequest
	err := json.NewDecoder(uh.Ctx.Request.Body).Decode(&body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	//Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&body)
	if err != nil {
		// handle error
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}
	if !b {
		// validation does not pass
		var message domain.ErrorMessage
		message.Code = 400
		message.Message = ""
		for i, err := range valid.Errors {

			if i == 0 {
				message.Message = err.Message
			} else {
				message.Message += ", " + err.Message
			}
		}
		uh.Ctx.Output.SetStatus(400)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	//Get commentId from param
	commentId := uh.Ctx.Input.Param(":commentId")
	commentIdint, err := strconv.Atoi(commentId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	res, err := uh.CommentUseCase.UpdateComment(uh.Ctx.Request.Context(), body, commentIdint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(res, true, true)
	return
}

func (uh commentHandler) DeleteComment() {
	//Get commentId from param
	commentId := uh.Ctx.Input.Param(":commentId")
	commentIdint, err := strconv.Atoi(commentId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	var message domain.Message

	err = uh.CommentUseCase.DeleteComment(uh.Ctx.Request.Context(), commentIdint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	message.Message = "Your comment has been successfully deleted"

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(message, true, true)
	return
}
