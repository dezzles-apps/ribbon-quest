package repositories

import (
	"github.com/dezzles-apps/go-common/db"
)

type RepositoryContext struct {
	CommonData CommonDataRepository
	Shiny      ShinyRepository
}

func InitaliseRepositories(database *db.Database) *RepositoryContext {
	ctx := &RepositoryContext{}
	ctx.CommonData = NewCommonDataService()
	ctx.Shiny = NewDatabaseShinyRepository(database, ctx.CommonData)
	return ctx
}
