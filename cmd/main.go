package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	_ "net/http"
	"usersRestAPI/internal/app/config"
	"usersRestAPI/internal/app/repository"
	"usersRestAPI/internal/app/rest"
	"usersRestAPI/internal/app/service"
	"usersRestAPI/internal/pkg/pg"
)

func main() {
	cfg := config.ReadConfig()
	if err := runPgMigrations(cfg.PostgresAddress, cfg.MigrationsPath); err != nil {
		log.Fatal(err)
	}
	db, err := pg.Dial(cfg.PostgresAddress)
	if err != nil {
		// Handle error
	}
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	userRepository := repository.UserRepositoryImpl{Db: db}
	userService := service.UserServiceImpl{
		UserRepository: userRepository,
	}
	userController := &rest.UserController{
		UserService: userService,
	}

	router := gin.Default()
	router.GET("/usersAll", userController.GetAllUsers)
	router.GET("/users", userController.GetUserById)
	router.POST("/users", userController.CreateUser)
	router.PATCH("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	routerErr := router.Run(cfg.HTTPAddr)
	if routerErr != nil {
		return
	}

}

func runPgMigrations(dsn, path string) error {
	log.Println("Running database migrations...")
	if path == "" {
		return errors.New("no migrations path provided")
	}
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	m, err := migrate.New(
		path,
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
