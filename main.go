package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/goccy/go-json"
	"github.com/teleyos/go_ex3_routerank/api"
	"github.com/teleyos/go_ex3_routerank/types"
	"sort"
	"log"
)

func main() {
	engine := html.New("./views",".html")
	engine.Reload(true)
	engine.Debug(true)
	engine.Layout("embed")

    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
		Views: engine,
    })

	app.Get("/",func(c *fiber.Ctx) error {
		return c.Render("index",fiber.Map{
			"Title": "Hello, World!",
		},"layouts/main")
	})

	app.Get("/documentation", func(c *fiber.Ctx) error {
		return c.Render("documentation", fiber.Map{
			"Title": "Documentation",
		},"layouts/main")
	})

	app.Get("/copy-paste", func(c *fiber.Ctx) error {
		return c.Render("copy-paste", fiber.Map{
			"Title": "Copy Paste",
		},"layouts/main")
	})

	app.Get("/file", func(c *fiber.Ctx) error {
		return c.Render("file", fiber.Map{
			"Title": "File",
		}, "layouts/main")
	})

	app.Post("/api",func(c *fiber.Ctx) error {
		var trips []types.Trip

		if err := c.BodyParser(&trips); err != nil {
			return c.SendString("didn't work")
		}

		api.ScoreTrips(&trips)
		sort.Sort(types.ByScore(trips))

		bytes, err := json.Marshal(trips)

		if err!=nil{
			return c.SendString("didn't work 2")
		}

		c.Type("json","utf-8")
		return c.Send(bytes)
	})

	app.Post("/api/form",func(c *fiber.Ctx) error {
		var trips []types.Trip

		data := []byte(c.FormValue("data"))

		if err := json.Unmarshal(data,&trips); err != nil {
			return c.SendString("didn't work")
		}

		api.ScoreTrips(&trips)
		sort.Sort(types.ByScore(trips))

		bytes, err := json.MarshalIndent(trips,"","  ")

		if err!=nil{
			return c.SendString("didn't work 2")
		}

		c.Type("json","utf-8")
		return c.Send(bytes)
	})

	log.Fatal(app.Listen(":3000"))
}
