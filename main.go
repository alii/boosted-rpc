package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hugolgst/rich-go/client"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/connected", func(ctx *fiber.Ctx) error {
		println("received connected")

		now := time.Now()

		// Safe to call login every time as we check
		// if we are already logged in
		if err := client.Login("932730208820265011"); err != nil {
			println("Could not login to boosted RPC")
			return err
		}

		if err := client.SetActivity(client.Activity{
			State:      "Riding my Boosted Board v3",
			Details:    "Out and about...",
			LargeImage: "v3",
			LargeText:  "Picture of the v3 board",
			SmallImage: "logo_black",
			SmallText:  "Boosted Boards Logo",
			Timestamps: &client.Timestamps{Start: &now},
		}); err != nil {
			println("Couldn't start boosted rpc! " + err.Error())
			return err
		}

		return ctx.SendStatus(200)
	})

	app.Get("/disconnected", func(ctx *fiber.Ctx) error {
		println("received disconnected")

		client.Logout()
		return ctx.SendStatus(200)
	})

	println("ight we lit on " + port)

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
