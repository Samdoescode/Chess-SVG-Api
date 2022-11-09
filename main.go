package main


import (
	"os"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
	"github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)


type Fen struct {
	Fen string `query:"fen"`
}

type boardReq struct {

	Fen string `form:"Fen" json:"Fen" xml:"Fen"`
}

func main() {

	
	router := fiber.New()

	router.Use("/", filesystem.New(filesystem.Config{
        Root: rice.MustFindBox("build").HTTPBox(),
    }))

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.Get("/api", createSvg)
	router.Post("/api/adv", avdCreateSvg)

	

	router.Listen("0.0.0.0:$PORT")
}

func createSvg(c *fiber.Ctx) error {

	fen := new(Fen)

	// check if fun is in call
	if err := c.QueryParser(fen); err != nil {
		return c.SendString("Oh no!")
	}

	name := strings.ReplaceAll(fen.Fen, "/", "")

	// check if in system + if in system send it
	if _, err := os.Stat("build/images/" + name); err == nil {
		return c.SendFile("build/images/" + name)
	}

	// create a file
	f, err := os.Create("build/images/" + name + ".svg")
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

	return c.SendFile("build/images/" + name + ".svg")

}

func avdCreateSvg(c *fiber.Ctx)error{

	

	boardReq := new(boardReq)

	if err := c.BodyParser(boardReq); err != nil {
		return err
	}

	// check if fun is in call
	name := strings.ReplaceAll(boardReq.Fen, "/", "")

	// check if in system + if in system send it
	if _, err := os.Stat("build/images/" + name); err == nil {
		return c.SendFile("build/images/" + name)
	}

	// create a file
	f, err := os.Create("build/images/" + name + ".svg")
	if err != nil {
		c.SendString("Couldn't Create the File" + err.Error())
	}
	defer f.Close()

	// create board position
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(boardReq.Fen)); err != nil {
		return c.SendString("This position is not valid" + err.Error())
	}

	

	// write board SVG to file
	if err := image.SVG(f, pos.Board()); err != nil {
		return c.SendString("We could not complete the SVG create Process" + err.Error())
	}

	return c.SendFile("build/images/" + name + ".svg")

}