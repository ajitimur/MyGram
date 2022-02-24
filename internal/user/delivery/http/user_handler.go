package http

import (
	"MyGram/internal/domain"
	"MyGram/pkg/helpers"
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type userHandler struct {
	beego.Controller
	UserUseCase domain.UserUseCase
}

func NewUserHandler(usecase domain.UserUseCase) {
	pHandler := &userHandler{
		UserUseCase: usecase,
	}

	beego.Router("/users/register", pHandler, "post:StoreUser")
	beego.Router("/users/login", pHandler, "post:LoginUser")
	beego.Router("/users", pHandler, "put:UpdateUser")
	beego.Router("/users", pHandler, "delete:DeleteUser")
}

func (uh userHandler) StoreUser() {

	var body domain.UserRequest

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

	body.Password = helpers.HashPassword(body.Password)
	res, err := uh.UserUseCase.SaveUser(uh.Ctx.Request.Context(), body)
	if err != nil {
		errSplit := strings.Split(err.Error(), " ")
		// fmt.Println(errSplit[7])

		if errSplit[7] == `"users_username_key"` {
			var message domain.ErrorMessage
			message.Code = 400
			message.Message = "Username Already Used"
			uh.Ctx.Output.SetStatus(400)
			uh.Ctx.Output.JSON(message, true, true)
			return
		} else if errSplit[7] == `"users_email_key"` {
			var message domain.ErrorMessage
			message.Code = 400
			message.Message = "Email Already Used"
			uh.Ctx.Output.SetStatus(400)
			uh.Ctx.Output.JSON(message, true, true)
			return
		} else {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			uh.Ctx.Output.SetStatus(500)
			uh.Ctx.Output.JSON(message, true, true)
			return
		}
	}

	uh.Ctx.Output.SetStatus(201)
	uh.Ctx.Output.JSON(res, true, true)
	return
}

func (uh userHandler) UpdateUser() {

	currentUser := _context.Get(uh.Ctx.Request, "currentUser")

	var marshalled map[string]interface{}

	data, err := json.Marshal(currentUser)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
	}
	err = json.Unmarshal(data, &marshalled)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
	}

	IdMarshal := marshalled["Id"]
	Idfloat := IdMarshal.(float64)
	Idint := int(Idfloat)

	var body domain.UserUpdateRequest

	err = json.NewDecoder(uh.Ctx.Request.Body).Decode(&body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
	}

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

	res, err := uh.UserUseCase.UpdateUser(uh.Ctx.Request.Context(), body, Idint)
	if err != nil {
		errSplit := strings.Split(err.Error(), " ")
		if errSplit[7] == `"users_username_key"` {
			var message domain.ErrorMessage
			message.Code = 400
			message.Message = "Username Already Used"
			uh.Ctx.Output.SetStatus(400)
			uh.Ctx.Output.JSON(message, true, true)
			return
		} else if errSplit[7] == `"users_email_key"` {
			var message domain.ErrorMessage
			message.Code = 400
			message.Message = "Email Already Used"
			uh.Ctx.Output.SetStatus(400)
			uh.Ctx.Output.JSON(message, true, true)
			return
		} else {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			uh.Ctx.Output.SetStatus(500)
			uh.Ctx.Output.JSON(message, true, true)
			return
		}
	}

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(res, true, true)
	return
}

func (uh userHandler) DeleteUser() {

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

	var message domain.Message

	err = uh.UserUseCase.DeleteUser(uh.Ctx.Request.Context(), Idint)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	message.Message = "Your Account has been successfully deleted"

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(message, true, true)
	return
}

func (uh userHandler) LoginUser() {

	var body domain.LoginRequest
	// var message domain.Message

	err := json.NewDecoder(uh.Ctx.Request.Body).Decode(&body)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 500
		message.Message = "Internal Server Error"
		uh.Ctx.Output.SetStatus(500)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

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

	user, err := uh.UserUseCase.GetUserByEmail(uh.Ctx.Request.Context(), body.Email)
	if err != nil {
		var message domain.ErrorMessage
		message.Code = 400
		message.Message = "User does not exist"
		uh.Ctx.Output.SetStatus(400)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	comparePass := helpers.ComparePassword(user.Password, body.Password)
	if !comparePass {
		var message domain.ErrorMessage
		message.Code = 400
		message.Message = "Invalid Email/Password"
		uh.Ctx.Output.SetStatus(400)
		uh.Ctx.Output.JSON(message, true, true)
		return
	}

	// fmt.Println(user, "<<<<<<")
	token := helpers.GenerateToken(user.Id, user.Email)
	var tokenMessage domain.TokenMessage
	tokenMessage.Token = token

	uh.Ctx.Output.SetStatus(200)
	uh.Ctx.Output.JSON(tokenMessage, true, true)
	return
}
