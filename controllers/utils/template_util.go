package utils

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"text/template"
)

func RenderTemplate(name string, obj interface{}) (string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read template %q", name)
	}

	t, err := template.New("resources").Parse(string(data))
	if err != nil {
		return "", errors.Wrapf(err, "failed to parse template %q", name)
	}

	var rendered bytes.Buffer
	err = t.Execute(&rendered, obj)
	if err != nil {
		return "", errors.Wrapf(err, "failed to execute template %q", name)
	}
	return rendered.String(), nil
}