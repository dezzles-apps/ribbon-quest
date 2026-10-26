package repositories

import (
	dbrepos "dezzles-apps/rq-server/repositories/db"
	"dezzles-apps/rq-server/repositories/interfaces"
	"dezzles-apps/rq-server/repositories/static"

	"github.com/dezzles-apps/go-common/db"
)

type RepositoryContext struct {
	CommonData interfaces.CommonDataRepository
	Shiny      interfaces.ShinyRepository
	Ribbons    interfaces.RibbonRepository
}

func InitaliseRepositories(database *db.Database) *RepositoryContext {
	ctx := &RepositoryContext{}
	ctx.CommonData = static.NewCommonDataService()
	ctx.Shiny = dbrepos.NewDatabaseShinyRepository(database, ctx.CommonData)
	ctx.Ribbons = dbrepos.NewDatabaseRibbonRepository(database, ctx.CommonData)
	return ctx
}
