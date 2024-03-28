package logging

import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gitlab.com/revolutionize1/foward-api/internal/api/middleware"
	"gitlab.com/revolutionize1/foward-api/internal/app"
	"gitlab.com/revolutionize1/foward-api/internal/file"
)

func setOutput(file *os.File) {
	writer := io.MultiWriter(os.Stdout, file)
	log.SetOutput(writer)
}

func Setup() error {
	if app.Instance.Config.Api.Environment == "dev" {
		logFile, err := file.RetrieveOrCreateFile("log/logs.txt")
		if err != nil {
			return err
		}
		setOutput(logFile)

		logFile, err = file.RetrieveOrCreateFile("log/request-logs.txt")
		if err != nil {
			return err
		}

		middleware.UseLogger(app.Instance.FiberApp, logFile)
	}
	return nil
}
