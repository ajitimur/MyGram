package http

import (
	"MyGram/internal/domain"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type socialMediaHandler struct {
	beego.Controller
	SocialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaHandler(usecase domain.SocialMediaUseCase) {
	pHandler := &socialMediaHandler{
		SocialMediaUseCase: usecase,
	}

	beego.Router("/socialmedias", pHandler, "get:GetSocialMedias")
	beego.Router("/socialmedias", pHandler, "post:StoreSocialMedia")
	beego.Router("/socialmedias/:socialmediaId", pHandler, "put:UpdateSocialMedia")
	beego.Router("/socialmedias/:socialmediaId", pHandler, "delete:DeleteSocialMedia")
}

func (uh socialMediaHandler) StoreSocialMedia() {
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
	var body domain.SocialMediaRequest

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

	//store socmed
	res, err := uh.SocialMediaUseCase.SaveSocialMedia(uh.Ctx.Request.Context(), body)
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

func (uh socialMediaHandler) GetSocialMedias() {
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

	//Get
	res, err := uh.SocialMediaUseCase.GetSocialMedias(uh.Ctx.Request.Context(), Idint)
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

func (uh socialMediaHandler) UpdateSocialMedia() {
	//req body
	var body domain.SocialMediaRequest
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

	//Get socialmediaId from param
	socialmediaId := uh.Ctx.Input.Param(":socialmediaId")
	socialmediaIdint, err := strconv.Atoi(socialmediaId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	res, err := uh.SocialMediaUseCase.UpdateSocialMedia(uh.Ctx.Request.Context(), body, socialmediaIdint)
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

func (uh socialMediaHandler) DeleteSocialMedia() {
	//Get socialmediaId from param
	socialmediaId := uh.Ctx.Input.Param(":socialmediaId")
	socialmediaIdint, err := strconv.Atoi(socialmediaId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	var message domain.Message

	err = uh.SocialMediaUseCase.DeleteSocialMedia(uh.Ctx.Request.Context(), socialmediaIdint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	message.Message = "Your social media has been successfully deleted"

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(message, true, true)
	return
}
