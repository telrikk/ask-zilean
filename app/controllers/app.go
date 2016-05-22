package controllers

import "github.com/revel/revel"

// App is the Ask Zilean app
type App struct {
	*revel.Controller
}

// Index renders the page for Ask Zilean
func (c App) Index() revel.Result {
	return c.Render()
}
