package application

import (
	"github.com/ichilly/Kruise/controllers/defaults_setter/plugins"
	"github.com/ichilly/Kruise/models"
)

func SetApplicationDefaults(app *models.Application) {
	SetMetadataDefaults(app)
	SetServiceDefaults(app)
	SetContainerDefaults(app)
	plugins.AddSwaggerPlugin(app)
}