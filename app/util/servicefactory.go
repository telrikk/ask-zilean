package util

import (
	"os"
	"sync"

	"github.com/telrikk/lol-go-api"
	"github.com/telrikk/lol-go-api/util"
)

var serviceFactory league.ServiceFactory
var once sync.Once

// GetServiceFactory creates or gets a singleton thread-safe service factory
func GetServiceFactory() league.ServiceFactory {
	once.Do(func() {
		config := new(util.ServiceConfiguration)
		config.RetryOnNonFatal = true
		config.RetryOnRateLimit = true
		config.MaxRetries = 6
		serviceFactory = league.NewServiceFactory(league.NA, os.Getenv("LOL_API_KEY"), *config)
	})
	return serviceFactory
}
