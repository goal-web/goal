package models

func init() {
	ProjectDefine.Appends = map[string]func(model *ProjectModel) any{
		"id-plus-1": func(model *ProjectModel) any {
			return model.Id + 0
		},
	}
}
