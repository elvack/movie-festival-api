package main

import (
	"fmt"
	"log"

	"github.com/elvack/movie-festival-api/common"
	"github.com/elvack/movie-festival-api/database"
	adminModel "github.com/elvack/movie-festival-api/model/admin"
	adminRepo "github.com/elvack/movie-festival-api/repository/admin"
	"github.com/elvack/movie-festival-api/service/admin"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var email, password string
	fmt.Print("email: ")
	if _, err := fmt.Scanln(&email); err != nil && err.Error() != common.UnexpectedNewline {
		log.Fatal(err)
	}
	fmt.Print("password: ")
	if _, err := fmt.Scanln(&password); err != nil && err.Error() != common.UnexpectedNewline {
		log.Fatal(err)
	}
	req := adminModel.SeedReq{
		Email: email,
		Password: password,
	}
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		log.Fatal(err)
	}
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	adminService := admin.NewService(adminRepo.NewRepo(db.GormDb))
	if err = adminService.Seed(&req); err != nil {
		log.Fatal(err)
	}
	log.Print(common.CreatedSuccessfully)
}
