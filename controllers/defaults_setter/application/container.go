package application

import (
	"github.com/ichilly/Kruise/controllers/defaults_setter/dependencies"
	"github.com/ichilly/Kruise/controllers/utils"
	"github.com/ichilly/Kruise/models"
)

// TODO: Load secrets in initContainer and write to env

func SetContainerDefaults(app *models.Application) {
	// Preload data
	if len(app.Component.Containers) == 1 {
		app.Component.Containers[0].Name = app.Metadata.App
	}
	if app.Component.Replicas == 0 {
		app.Component.Replicas = 1
	}

	// Setup service container
	container := getServiceContainer(app)
	if len(container.PortNames) == 0 {
		container.PortNames = []string{app.Metadata.App}
	}
	if container.ImageTag == "" {
		container.ImageTag = app.Metadata.Version
	}

	// Fill env
	envs := container.Env.Values
	if !utils.HasEnv(envs, "HOST") {
		envs = append(envs, &models.Value{
			Name:  "HOST",
			Value: "0.0.0.0",
		})
	}
	if !utils.HasEnv(envs, "PORT") {
		envs = append(envs, &models.Value{
			Name:  "PORT",
			Value: "80",
		})
	}
	container.Env.Values = envs

	// Fill dependency env
	dependencies.AddAwsConfig(app, container)
	dependencies.AddMysqlConfig(app, container)
}

func getServiceContainer(app *models.Application) *models.Container {
	for _, container := range app.Component.Containers {
		if container.Name == app.Metadata.App {
			return container
		}
	}
	return nil
}
