package controllers

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/models"
	"github.com/gookit/goutil/dump"
)

func HelloWorld() any {
	createProject()
	updateProject()
	dump.P(23)
	projects := getProjectList()
	return contracts.Fields{
		"data": projects,
	}
}

func createProject() {
	fields := contracts.Fields{
		"uuid":           "45",
		"name":           "your_project_name",
		"creator_id":     1,
		"group_id":       1,
		"key_id":         1,
		"repo_address":   "your_repo_address",
		"project_path":   "your_project_path",
		"default_branch": "main",
		"settings":       "{}",
	}
	model := models.NewProjectModel(fields)
	err := model.Save()
	if err != nil {
		fmt.Printf("Failed to create project: %v\n", err)
	} else {
		fmt.Println("Project created successfully111")
	}
}

func getProjectOne() any {
	query := models.ProjectQuery()
	project := query.Where("id", 1).First() // 假设查询 id 为 1 的记录
	fmt.Printf("Project: %+v\n", project)

	return project
}

func getProjectList() any {
	query := models.ProjectQuery()

	projects := query.Get().ToAnyArray() //
	dump.P(projects)

	return projects
}

func updateProject() {
	query := models.ProjectQuery()
	project := query.Where("id", 1).First() // 假设更新 id 为 1 的记录

	updateFields := contracts.Fields{
		"name": "updated_project_name11",
	}
	err := project.Update(updateFields)

	dump.P(err)
}

func deleteProject() {
	query := models.ProjectQuery()
	project := query.Where("id", 1).First() // 假设删除 id 为 1 的记录

	_ = project.Delete()

}
