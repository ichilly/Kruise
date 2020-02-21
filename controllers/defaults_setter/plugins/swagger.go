package plugins

import (
	"github.com/ichilly/Kruise/models"
)

func AddSwaggerPlugin(app *models.Application) {
	if !needsAddSwaggerPlugin(app) {
		return
	}

	swaggerContainer := getSwaggerContainer(app)
	if swaggerContainer == nil {
		swaggerContainer = &models.Container{
			Image:    "swaggerapi/swagger-ui",
			ImageTag: "latest",
			Name:     "swagger",
			PortNames: []string{"swagger"},
			Env:      &models.Env{},
		}
	}

	if !hasSwaggerAPIUrl(swaggerContainer) {
		swaggerContainer.Env.ValueFroms = []*models.ValueFrom{
			{
				Name: "SWAGGER_ROOT",
				ValueRef: &models.ValueRef{
					Key:  "swagger.json",
					Name: "swagger-config",
					Type: "configMapKeyRef",
				},
			},
		}
	}

	if !hasSwaggerPortConfig(app) {
		swaggerPort := &models.ServicePort{
			Name:       "swagger",
			Port:       8080,
			TargetPort: 8080,
		}
		app.Component.Service.Ports = append(app.Component.Service.Ports, swaggerPort)
	}

	app.Component.Containers = append(app.Component.Containers, swaggerContainer)
}

func needsAddSwaggerPlugin(app *models.Application) bool {
	return *app.Metadata.Plugins.Swagger
}

func getSwaggerContainer(app *models.Application) *models.Container {
	for _, container := range app.Component.Containers {
		if container.Name == "swagger" {
			return container
		}
	}
	return nil
}

func hasSwaggerPortConfig(app *models.Application) bool {
	for _, port := range app.Component.Service.Ports {
		if port.Name == "swagger" {
			return true
		}
	}
	return false
}

func hasSwaggerAPIUrl(container *models.Container) bool {
	for _, env := range container.Env.Values {
		if env.Name == "API_URL" {
			return true
		}
	}
	return false
}
