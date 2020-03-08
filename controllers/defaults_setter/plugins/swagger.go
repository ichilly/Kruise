package plugins

import (
	"github.com/ichilly/Kruise/models"
)

func AddSwaggerPlugin(app *models.Application) {
	if !needsAddSwaggerPlugin(app) {
		return
	}
}

func needsAddSwaggerPlugin(app *models.Application) bool {
	return *app.Metadata.Plugins.Swagger
}
