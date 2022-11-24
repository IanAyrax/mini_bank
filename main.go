package main

import (
	"fmt"
	"log"
	"mini-bank/src/config"
	"mini-bank/src/models"
	"mini-bank/src/routes"
	"net/http"
	"os"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
)

var (
	validatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	ValidatorInit()

	//Migrate Database
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})

	app := fiber.New()
	app.Get("", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Service is working")
	})
	routers := routes.Routers{
		App:       app,
		DB:        db,
		Validator: validatorDriver,
	}
	routers.Routers()

	fmt.Println("Server is running....")
	log.Fatal(app.Listen(":3000"))
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	validatorDriver = validator.New()

	err := enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(validatorDriver, transID)
	if err != nil {
		fmt.Println(err)
	}
	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
