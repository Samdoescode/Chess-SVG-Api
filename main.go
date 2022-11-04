package main

import (
	"os"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

type Fen struct {
	Fen string `query:"fen"`
}

func main() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/api", createSvg)

	router.Listen(":8080")
}

func createSvg(c *fiber.Ctx) error {

	fen := new(Fen)

	// check if fun is in call
	if err := c.QueryParser(fen); err != nil {
		return c.SendString("Oh no!")
	}

	name := strings.ReplaceAll(fen.Fen, "/", "")

	// check if in system + if in system send it
	if _, err := os.Stat("imaegs/"+name); err == nil {
		return c.SendFile("imaegs/"+name)
	}

	// create a file
	f, err := os.Create("images/"+ name + ".svg")
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
	if err := image.SVG(f, pos.Board()); err != nil {
		return c.SendString("We could not complete the SVG create Process" + err.Error())
	}

	return c.SendFile("images/"+ name + ".svg")

}
