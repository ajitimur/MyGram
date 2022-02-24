package middleware

import (
	"MyGram/internal/domain"
	"MyGram/pkg/helpers"
	"encoding/json"
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_context "github.com/gorilla/context"
)

func AuthMiddleware(us domain.UserUseCase) beego.HandleFunc {
	return func(ctx *context.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		// fmt.Println("masuk 1")

		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("1111111")
			var message domain.ErrorMessage
			message.Code = 401
			message.Message = "Unauthorized"
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(message, true, true)
			// ctx.Abort(400, "Bad Request")
			return
		}

		token, err := helpers.VerifyToken(ctx)
		if err != nil {
			fmt.Println("22222")
			var message domain.ErrorMessage
			message.Code = 401
			message.Message = "Unauthorized"
			ctx.Output.SetStatus(401)
			ctx.Output.JSON(message, true, true)
			return
		}

		var res map[string]interface{}

		data, err := json.Marshal(token)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}
		err = json.Unmarshal(data, &res)
		if err != nil {
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
		}

		userID := res["id"]
		Idfloat := userID.(float64)
		Idint := int(Idfloat)

		user, err := us.GetUserById(ctx.Request.Context(), Idint)
		if err != nil {
			fmt.Println("444444")
			var message domain.ErrorMessage
			message.Code = 500
			message.Message = "Internal Server Error"
			ctx.Output.SetStatus(500)
			ctx.Output.JSON(message, true, true)
			return
		}

		_context.Set(ctx.Request, "currentUser", user)

	}
}
