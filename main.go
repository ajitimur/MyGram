package main

import (
	"MyGram/internal/domain"
	middleware "MyGram/middlewares"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_userHandler "MyGram/internal/user/delivery/http"
	_userRepo "MyGram/internal/user/repository"
	_userUcase "MyGram/internal/user/usecase"

	_photoHandler "MyGram/internal/photo/delivery/http"
	_photoRepo "MyGram/internal/photo/repository"
	_photoUcase "MyGram/internal/photo/usecase"

	_commentHandler "MyGram/internal/comment/delivery/http"
	_commentRepo "MyGram/internal/comment/repository"
	_commentUcase "MyGram/internal/comment/usecase"

	_socialMediaHandler "MyGram/internal/socialMedia/delivery/http"
	_socialMediaRepo "MyGram/internal/socialMedia/repository"
	_socialMediaUcase "MyGram/internal/socialMedia/usecase"
)

func main() {
	dbstring := "postgres://postgres:postgres@localhost:5432/MyGram?sslmode=disable"
	//Gorm
	db, err := gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		fmt.Println(err, "++++")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err, "++++")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Photo{})
	db.AutoMigrate(&domain.Comment{})
	db.AutoMigrate(&domain.SocialMedia{})

	timeoutContext := time.Duration(10) * time.Second

	//User
	userRepository := _userRepo.NewUserRepository(db)
	userUseCase := _userUcase.NewUserUseCase(timeoutContext, userRepository)
	_userHandler.NewUserHandler(userUseCase)

	//Photo
	photoRepository := _photoRepo.NewPhotoRepository(db)
	photoUseCase := _photoUcase.NewPhotoUseCase(timeoutContext, photoRepository)
	_photoHandler.NewPhotoHandler(photoUseCase)

	//Comment
	commentRepository := _commentRepo.NewCommentRepository(db)
	commentUseCase := _commentUcase.NewCommentUseCase(timeoutContext, commentRepository)
	_commentHandler.NewCommentHandler(commentUseCase)

	//SocialMedia
	socialMediaRepository := _socialMediaRepo.NewSocialMediaRepository(db)
	socialMediaUseCase := _socialMediaUcase.NewSocialMediaUseCase(timeoutContext, socialMediaRepository)
	_socialMediaHandler.NewSocialMediaHandler(socialMediaUseCase)

	beego.InsertFilter("/users", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/photos", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/photos/:photoId", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/photos/:photoId", 2, middleware.AuthorizePhoto(photoUseCase))
	beego.InsertFilter("/comments", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/comments/:commentId", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/comments/:commentId", 2, middleware.AuthorizeComment(commentUseCase))
	beego.InsertFilter("/socialmedias", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/socialmedias/:socialmediaId", 1, middleware.AuthMiddleware(userUseCase))
	beego.InsertFilter("/socialmedias/:socialmediaId", 2, middleware.AuthorizeSocialMedia(socialMediaUseCase))

	beego.Run()
}
