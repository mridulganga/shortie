package main

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

const defaultHost = "shortie.mridulganga.dev"

func ListRedirects(db *leveldb.DB) []map[string]string {
	collection := []map[string]string{}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		collection = append(collection, map[string]string{
			"code": string(iter.Key()),
			"url":  string(iter.Value()),
		})
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		logrus.Error(err)
	}
	return collection
}

func main() {
	app := fiber.New()

	app.Static("/ui", "./ui/public")

	db, err := leveldb.OpenFile("urls.db", nil)
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		if c.Query("list") == "true" {
			return c.JSON(ListRedirects(db))
		}
		return c.SendString("Hello World!")
	})

	app.Get("/domain", func(c *fiber.Ctx) error {
		host := os.Getenv("HOST")
		if host == "" {
			host = defaultHost
		}
		host += "/"
		return c.SendString(host)
	})

	app.Get("/:code", func(c *fiber.Ctx) error {
		url, err := db.Get([]byte(c.Params("code")), nil)
		if err != nil {
			return c.Redirect("/")
		}
		return c.Redirect(string(url))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		req := map[string]string{}
		c.BodyParser(&req)
		if req["code"] != "" && req["url"] != "" {
			if err := db.Put([]byte(req["code"]), []byte(req["url"]), nil); err != nil {
				return err
			}
			return nil
		}
		return errors.New("invalid code or url")
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		req := map[string]string{}
		c.BodyParser(&req)
		if err := db.Delete([]byte(req["code"]), nil); err != nil {
			return err
		}
		return c.SendString("Deleted")
	})

	app.Listen(":8000")
}
