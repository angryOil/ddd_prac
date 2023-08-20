package main

import (
	usecase "ddd_prac/domain/use-case"
	"ddd_prac/infrastructure/gorm"
	"ddd_prac/infrastructure/user"
	"ddd_prac/port/in/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	var gormEngine = gorm.NewGorm()
	err := gormEngine.MigrateUserTable()
	if err != nil {
		log.Panicf("fail migrate users\nerr: %e", err)
	}

	var userRepository = user.NewRepositoryImpl(gormEngine)
	var userService = usecase.NewUserService(userRepository)
	user, err := userService.Login("1", "2")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("user: %v\n", user)

	e := echo.New()

	controller := controller.NewUserController(userService, e)
	controller.Login()
	controller.Register()
	e.Logger.Fatal(e.Start(":8080"))
}
