// generated code - do not edit
package probe

import (
	"github.com/gin-gonic/gin"

	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

type Playground struct {
	r                  *gin.Engine
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	formStage          *form.StageStruct
	tableStage         *form.StageStruct
}