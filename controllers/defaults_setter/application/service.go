package application

import "github.com/ichilly/Kruise/models"

func SetServiceDefaults(app *models.Application) {
	if app.Component.Service.Name == "" {
		app.Component.Service.Name = app.Metadata.App
	}

	if len(app.Component.Service.Ports) == 1 {
		port := app.Component.Service.Ports[0]
		port.Name = app.Metadata.App
		if port.Port == 0 {
			port.Port = port.TargetPort
		}
	}
}