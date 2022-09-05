package main

import (
	"er-api/database"
	"er-api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	database.ConnectApi()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	Routes(app)
	app.Listen(":3000") //port
}

// Parity Functions
func GetRates(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetUsd(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "USD").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetEur(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "EUR").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetTry(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "TRY").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetUsdTry(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "USD").Where("target_code=?", "TRY").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetUsdUsd(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "USD").Where("target_code=?", "USD").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetUsdEur(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "USD").Where("target_code=?", "EUR").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetEurTry(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "EUR").Where("target_code=?", "TRY").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetEurUsd(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "EUR").Where("target_code=?", "USD").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetEurEur(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "EUR").Where("target_code=?", "EUR").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetTryEur(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "TRY").Where("target_code=?", "EUR").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetTryUsd(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "TRY").Where("target_code=?", "USD").Find(&rates)
	return c.Status(200).JSON(rates)
}
func GetTryTry(c *fiber.Ctx) error {
	db := database.DB
	rates := []model.ExchangeRate{}
	db.Where("base_code=?", "TRY").Where("target_code=?", "TRY").Find(&rates)
	return c.Status(200).JSON(rates)
}
func Routes(app *fiber.App) { //endpoints
	api := app.Group("/api/rates")
	api.Get("/ALL", GetRates)
	api.Get("/USD", GetUsd)
	api.Get("/EUR", GetEur)
	api.Get("/TRY", GetTry)
	api.Get("/USD/TRY", GetUsdTry)
	api.Get("/USD/USD", GetUsdUsd)
	api.Get("/USD/EUR", GetUsdEur)
	api.Get("/EUR/TRY", GetEurTry)
	api.Get("/EUR/USD", GetEurUsd)
	api.Get("/EUR/EUR", GetEurEur)
	api.Get("/TRY/EUR", GetTryEur)
	api.Get("/TRY/USD", GetTryUsd)
	api.Get("/TRY/TRY", GetTryTry)
}
