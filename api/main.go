package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
	"image/color"
	
	"os"
)



type Fen struct {
	Fen string `query:"fen"`
	
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/api", func(c *fiber.Ctx) error {


		fen := new(Fen)
		
		// check if fun is in call
		if err := c.QueryParser(fen); err != nil {
			
			return c.SendString("Oh no!")
		}
		
		// create a file
		f, err := os.Create("2.svg")
		if err != nil {
			c.SendString("Couldn't Create the File" + err.Error())
		}
		defer f.Close()

		// create board position
		pos := &chess.Position{}
		if err := pos.UnmarshalText([]byte(fen.Fen)); err != nil {
			return c.SendString("This position is not valid" + err.Error())
		}

		// write board SVG to file
		yellow := color.RGBA{255, 255, 0, 1}
		mark := image.MarkSquares(yellow, chess.D2, chess.D4)
		if err := image.SVG(f, pos.Board(), mark); err != nil {
			return c.SendString("We could not complete the SVG create Process" + err.Error())
		}

		return c.SendFile("2.svg")
	})

	router.Listen(":3000")
}
