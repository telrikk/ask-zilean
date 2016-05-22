package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {

	p.Task("development", do.S{"build", do.P{"webpack", "server"}}, nil)

	p.Task("build", do.S{}, func(c *do.Context) {
		c.Run("go fmt ./...", do.M{"$in": "./"})
		c.Run("go fix ./...", do.M{"$in": "./"})
		c.Run("go vet ./...", do.M{"$in": "./"})
		c.Run("go build", do.M{"$in": "./app"})
	}).Src("**/*.go")

	p.Task("server", do.S{}, func(c *do.Context) {
		c.Run("revel run github.com/telrikk/ask-zilean", do.M{"$in": "./"})
	})

	p.Task("webpack", do.S{}, func(c *do.Context) {
		c.Run("npm run webpack", do.M{"$in": "./"})
	})
}

func main() {
	do.Godo(tasks)
}
