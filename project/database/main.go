package main

import (
	"api/repositories"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()
}

func initDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	dial := mysql.Open(dsn)

	return gorm.Open(dial, nil)
}

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}

	customerRepository, err := repositories.NewCustomerRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := viper.GetString("app.name")
		return c.SendString("Hello, " + name)
	})

	app.Get("/customers", func(c *fiber.Ctx) error {
		customers, err := customerRepository.FindAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(customers)
	})

	app.Listen(":3000")
}
