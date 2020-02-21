package dependencies

import (
	"github.com/ichilly/Kruise/controllers/utils"
	"github.com/ichilly/Kruise/models"
)

func AddAwsConfig(app *models.Application, container *models.Container) {
	if *app.Metadata.Dependencies.Aws {
		envFroms := container.Env.ValueFroms

		if !utils.HasEnvFrom(envFroms, "AWS_ACCESS_KEY") {
			envFroms = append(envFroms, &models.ValueFrom{
				Name:  "AWS_ACCESS_KEY",
				ValueRef: &models.ValueRef{
					Key:  "AWS_ACCESS_KEY",
					Name: "aws-secret",
					Type: "secretKeyRef",
				},
			})
		}
		if !utils.HasEnvFrom(envFroms, "AWS_SECRET_KEY") {
			envFroms = append(envFroms, &models.ValueFrom{
				Name:  "AWS_SECRET_KEY",
				ValueRef: &models.ValueRef{
					Key:  "AWS_SECRET_KEY",
					Name: "aws-secret",
					Type: "secretKeyRef",
				},
			})
		}
		container.Env.ValueFroms = envFroms
	}
}