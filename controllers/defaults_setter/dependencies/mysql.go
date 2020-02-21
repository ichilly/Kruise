package dependencies

import (
	"github.com/ichilly/Kruise/controllers/utils"
	"github.com/ichilly/Kruise/models"
)

func AddMysqlConfig(app *models.Application, container *models.Container) {
	if *app.Metadata.Dependencies.Mysql {
		envs := container.Env.Values

		if !utils.HasEnv(envs, "MYSQL_HOST") {
			envs = append(envs, &models.Value{
				Name:  "MYSQL_HOST",
				Value: "mysql-service",
			})
		}
		if !utils.HasEnv(envs, "MYSQL_PORT") {
			envs = append(envs, &models.Value{
				Name:  "MYSQL_PORT",
				Value: "3306",
			})
		}
		container.Env.Values = envs
	}
}