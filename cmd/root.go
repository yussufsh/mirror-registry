package cmd

import (
	_ "embed" // embed package is used to embed service files
	"log"
	"os"

	"github.com/spf13/cobra"
)

//go:embed "assets/quay.service"
var quayServiceBytes []byte

//go:embed "assets/postgres.service"
var postgresServiceBytes []byte

//go:embed "assets/redis.service"
var redisServiceBytes []byte

type service struct {
	name     string
	location string
	bytes    []byte
}

var services = []service{
	{
		"quay-app", "/etc/systemd/system/quay-app.service", quayServiceBytes,
	},
	{
		"quay-postgres", "/etc/systemd/system/quay-postgres.service", postgresServiceBytes,
	},
	{
		"quay-redis", "/etc/systemd/system/quay-redis.service", redisServiceBytes,
	},
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

func check(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "quay-installer",
		Short: "A generator for Cobra based Applications",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
