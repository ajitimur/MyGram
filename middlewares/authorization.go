package middleware

import (
	"MyGram/internal/domain"
	"encoding/json"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_context "github.com/gorilla/context"
)

func AuthorizePhoto(u domain.PhotoUseCase) beego.HandleFunc {
	return func(ctx *context.Context) {
		//Get UserId
		currentUser := _context.Get(ctx.Request, "currentUser")

		var marshalled map[string]interface{}

		data, err := json.Marshal(currentUser)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		err = json.Unmarshal(data, &marshalled)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		// fmt.Println(marshalled)
		IdMarshal := marshalled["Id"]
		Idfloat := IdMarshal.(float64)
		Idint := int(Idfloat)

		//Get photoId from param
		photoId := ctx.Input.Param(":photoId")
		photoIdint, err := strconv.Atoi(photoId)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}

		//Get UserId from photo
		photo, err := u.GetPhotoById(ctx.Request.Context(), photoIdint)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 404
			message.Message = "Photo Not Found"
			ctx.Output.SetStatus(404)
			ctx.Output.JSON(message, true, true)
		}

		// fmt.Println(Idint, photo.Id)
		//compare
		if Idint != photo.UserId {
			var message domain.ErrorMessage
			message.Code = 403
			message.Message = "Access Forbidden"
			ctx.Output.SetStatus(403)
			ctx.Output.JSON(message, true, true)
			return
		}

	}
}

func AuthorizeComment(u domain.CommentUseCase) beego.HandleFunc {
	return func(ctx *context.Context) {
		//Get UserId
		// fmt.Println("masuk 2")
		currentUser := _context.Get(ctx.Request, "currentUser")

		var marshalled map[string]interface{}

		data, err := json.Marshal(currentUser)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		err = json.Unmarshal(data, &marshalled)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		// fmt.Println(marshalled)
		IdMarshal := marshalled["Id"]
		Idfloat := IdMarshal.(float64)
		Idint := int(Idfloat)

		//Get commentId from param
		commentId := ctx.Input.Param(":commentId")
		commentIdint, err := strconv.Atoi(commentId)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}

		//Get UserId from comment
		comment, err := u.GetCommentById(ctx.Request.Context(), commentIdint)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 404
			message.Message = "Comment Not Found"
			ctx.Output.SetStatus(404)
			ctx.Output.JSON(message, true, true)
		}

		// fmt.Println(Idint, comment.Id)
		//compare
		if Idint != comment.UserId {
			var message domain.ErrorMessage
			message.Code = 403
			message.Message = "Access Forbidden"
			ctx.Output.SetStatus(403)
			ctx.Output.JSON(message, true, true)
			return
		}

	}
}

func AuthorizeSocialMedia(u domain.SocialMediaUseCase) beego.HandleFunc {
	return func(ctx *context.Context) {
		//Get UserId
		// fmt.Println("masuk 2")
		currentUser := _context.Get(ctx.Request, "currentUser")

		var marshalled map[string]interface{}

		data, err := json.Marshal(currentUser)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		err = json.Unmarshal(data, &marshalled)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		// fmt.Println(marshalled)
		IdMarshal := marshalled["Id"]
		Idfloat := IdMarshal.(float64)
		Idint := int(Idfloat)

		//Get socialmediaId from param
		socialmediaId := ctx.Input.Param(":socialmediaId")
		socialmediaIdint, err := strconv.Atoi(socialmediaId)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}

		//Get UserId from socmed
		socMed, err := u.GetSocialMediaById(ctx.Request.Context(), socialmediaIdint)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 404
			message.Message = "Social Media Not Found"
			ctx.Output.SetStatus(404)
			ctx.Output.JSON(message, true, true)
		}

		// fmt.Println(Idint, socMed.Id)
		//compare
		if Idint != socMed.UserId {
			var message domain.ErrorMessage
			message.Code = 403
			message.Message = "Access Forbidden"
			ctx.Output.SetStatus(403)
			ctx.Output.JSON(message, true, true)
			return
		}

	}
}
