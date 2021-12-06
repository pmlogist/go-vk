package testutil

import (
	"log"
	"os"

	"github.com/hnktt/go-vk/internal/projectpath"
	"github.com/joho/godotenv"
)

type Envars struct {
	VK_SERVICE_TOKEN string
}

func LoadEnv() Envars {
	err := godotenv.Load(projectpath.Root() + "/.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := Envars{
		VK_SERVICE_TOKEN: os.Getenv("VK_SERVICE_TOKEN"),
	}

	return env
}
