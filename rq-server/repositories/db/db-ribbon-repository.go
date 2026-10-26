package db

import (
	"dezzles-apps/rq-server/repositories/interfaces"
	_ "embed"

	"github.com/dezzles-apps/go-common/db"
)

type DBRibbonInitialiser interface {
}

type DBRibbonRepository struct {
	connection *db.Database
	common     interfaces.CommonDataRepository
}

func NewDatabaseRibbonRepository(
	connection *db.Database,
	common interfaces.CommonDataRepository,
) interfaces.RibbonRepository {
	return &DBRibbonRepository{
		connection: connection,
		common:     common,
	}
}
