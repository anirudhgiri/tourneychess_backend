package tourneychess_backend

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func Start() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env failed to load")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{"status": 200, "message": "Hello, World!"})
	})

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Invalid port value %d", PORT)
	}

	return e.Start(fmt.Sprintf(":%d", PORT))
}
