package application

import "github.com/ichilly/Kruise/models"

func SetMetadataDefaults(app *models.Application) {
	if app.Metadata.Namespace == "" {
		app.Metadata.Namespace = "default"
	}
	if app.Metadata.Output == "" {
		app.Metadata.Output = "yaml"
	}
}