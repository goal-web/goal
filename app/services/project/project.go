package project

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/requests/project"
	project0 "github.com/goal-web/goal/app/results/project"
)

func init() {

	ProjectServiceDefine.GetProject = func(req *project.GetProjectReq, ctx contracts.Context) (*project0.GetProjectResult, error) {
		return &project0.GetProjectResult{
			// 查询项目
		}, nil

	}

}
