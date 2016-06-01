package util

import (
	"github.com/revel/revel"
	"github.com/telrikk/lol-go-api/util"
)

// HandleError will inspect an error object and call the appropriate controller
// function
func HandleError(c revel.Controller, err error) revel.Result {
	riotError, isRiotError := err.(util.APIError)
	if isRiotError && riotError.Status.StatusCode == 404 {
		return c.NotFound(riotError.Status.Message)
	}
	return c.RenderError(err)
}
