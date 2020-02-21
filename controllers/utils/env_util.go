package utils

import "github.com/ichilly/Kruise/models"

func HasEnv(values []*models.Value, envName string) bool {
	for _, env := range values {
		if env.Name == envName {
			return true
		}
	}
	return false
}

func HasEnvFrom(values []*models.ValueFrom, envName string) bool {
	for _, env := range values {
		if env.Name == envName {
			return true
		}
	}
	return false
}