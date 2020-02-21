package renderers

import (
	"strings"
	"github.com/ichilly/Kruise/controllers/utils"
	"github.com/ichilly/Kruise/models"
)

type Renderer struct {
	templateDir string
}

type RenderedFile struct {
	Name    string
	Content string
}

func NewRenderer(templateDir string) (*Renderer, error) {
	return &Renderer{templateDir}, nil
}

// Render application into an all-in-one file
func (r *Renderer) RenderApplication(app *models.Application) (string, error) {
	fileContents := []string{}

	resourceKinds := r.getResourceKinds(app)
	resourceFiles, err := r.renderResourceKinds(resourceKinds, app)
	if err != nil {
		return "", err
	}
	fileContents = append(fileContents, utils.GetAllValues(resourceFiles)...)

	if r.needsRenderKustomize(app) {
		fileNames := utils.GetAllKeys(resourceFiles)
		kustomizationFile, err := r.renderKustomization(app, fileNames)
		if err != nil {
			return "", err
		}
		fileContents = append(fileContents, kustomizationFile.Content)
	}
	return strings.Join(fileContents, "---\n"), nil
}

func (r *Renderer) getResourceKinds(app *models.Application) []string {
	kinds := []string{}

	if app.Component.Service != nil {
		kinds = append(kinds, "service")
	}
	if len(app.Component.Containers) > 0 {
		kinds = append(kinds, "deployment")
	}
	if *app.Metadata.Plugins.Swagger {
		kinds = append(kinds, "ingress")
	}
	return kinds
}

func (r *Renderer) needsRenderKustomize(app *models.Application) bool {
	return app.Metadata.Output == "kustomize"
}

func (r *Renderer) renderResourceKind(kind string, data interface{}) (*RenderedFile, error) {
	templateFile := utils.GetFilePath(r.templateDir, kind+".yaml")
	result, err := utils.RenderTemplate(templateFile, data)
	if err != nil {
		return nil, err
	}
	return &RenderedFile{Name: kind + ".yaml", Content: result}, nil
}

func (r *Renderer) renderResourceKinds(kinds []string, app *models.Application) (map[string]string, error) {
	manifest := map[string]string{}

	for _, kind := range kinds {
		data := struct {
			App         *models.Application
			Service     *models.Service
			Containers  []*models.Container
		}{
			App:        app,
			Service:    app.Component.Service,
			Containers: app.Component.Containers,
		}
		fileInfo, err := r.renderResourceKind(kind, data)
		if err != nil {
			return map[string]string{}, err
		}
		manifest[fileInfo.Name] = fileInfo.Content
	}
	return manifest, nil
}

func (r *Renderer) renderKustomization(app *models.Application, resources []string) (*RenderedFile, error) {
	data := struct {
		App       *models.Application
		Resources []string
	}{
		App:       app,
		Resources: resources,
	}
	return r.renderResourceKind("kustomization", data)
}
