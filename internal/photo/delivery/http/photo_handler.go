package http

import (
	"MyGram/internal/domain"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type photoHandler struct {
	beego.Controller
	PhotoUseCase domain.PhotoUseCase
}

func NewPhotoHandler(usecase domain.PhotoUseCase) {
	pHandler := &photoHandler{
		PhotoUseCase: usecase,
	}

	beego.Router("/photos", pHandler, "get:GetPhotos")
	beego.Router("/photos", pHandler, "post:StorePhoto")
	beego.Router("/photos/:photoId", pHandler, "put:UpdatePhoto")
	beego.Router("/photos/:photoId", pHandler, "delete:DeletePhoto")
}

func (uh photoHandler) StorePhoto() {
	var body domain.PhotoRequest

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

	body.UserId = Idint

	//Store Photo
	res, err := uh.PhotoUseCase.SavePhoto(uh.Ctx.Request.Context(), body)
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

func (uh photoHandler) GetPhotos() {
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

	//Get Photo
	res, err := uh.PhotoUseCase.GetPhotos(uh.Ctx.Request.Context(), Idint)
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

func (uh photoHandler) UpdatePhoto() {

	//req body
	var body domain.PhotoUpdateRequest
	err := json.NewDecoder(uh.Ctx.Request.Body).Decode(&body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
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

	//Get photoId from param
	photoId := uh.Ctx.Input.Param(":photoId")
	photoIdint, err := strconv.Atoi(photoId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	res, err := uh.PhotoUseCase.UpdatePhoto(uh.Ctx.Request.Context(), body, photoIdint)
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

func (uh photoHandler) DeletePhoto() {
	//Get photoId from param
	photoId := uh.Ctx.Input.Param(":photoId")
	photoIdint, err := strconv.Atoi(photoId)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	var message domain.Message

	err = uh.PhotoUseCase.DeletePhoto(uh.Ctx.Request.Context(), photoIdint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	message.Message = "Your Photo has been successfully deleted"

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(message, true, true)
	return
}
