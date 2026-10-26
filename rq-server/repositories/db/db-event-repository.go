package db

import (
	"database/sql"
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"
	"strings"
	"time"

	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"

	"github.com/dezzles-apps/go-common/db"
)

type DBEventRepository struct {
	db *db.Database
}

func NewDBEventRepository(
	database *db.Database,
) interfaces.EventRepository {
	return &DBEventRepository{
		db: database,
	}
}

func (es *DBEventRepository) GetEvents(
	ctx *cmodel.DAContext,
	date *time.Time,
	category string,
	pokemon string,
) ([]dto.Event, errors.Error) {
	var query strings.Builder
	var params []any
	query.WriteString("SELECT * FROM events_v5 WHERE 1")

	if date != nil {
		params = append(params, date)
		query.WriteString(" AND timestamp < ?")
	}
	if category != "" {
		params = append(params, category)
		query.WriteString(" AND event_category = ?")
	}
	if pokemon != "" {
		params = append(params, pokemon)
		query.WriteString(" AND pokemon = ?")
	}
	query.WriteString(" LIMIT 10")
	rows, err := es.db.GetDB().Query(query.String(), params...)
	if err != nil {
		ctx.Logger.Error("Error getting events", zap.Error(err))
		return nil, errors.InternalServerError
	}
	defer rows.Close()
	ctx.Logger.Info("Latest events retrieved")
	return es.readEvents(ctx, rows)
}

func (es *DBEventRepository) readEvents(ctx *cmodel.DAContext, rows *sql.Rows) ([]dto.Event, errors.Error) {
	eventList := []dto.Event{}
	for rows.Next() {
		var ribbonKey sql.NullString
		var pokemon sql.NullString
		var nickname sql.NullString
		var ribbonName sql.NullString
		var ribbonCategory sql.NullString
		var ribbonType sql.NullString
		var pokedex sql.NullString
		event := dto.Event{}
		err := rows.Scan(
			&event.Category,
			&event.Type,
			&event.EventTime,
			&pokemon,
			&nickname,
			&ribbonKey,
			&ribbonName,
			&ribbonCategory,
			&ribbonType,
			&pokedex,
		)
		if err != nil {
			ctx.Logger.Error("Error loading event row", zap.Error(err))
			return nil, errors.InternalServerError
		}
		event.MetaData = make(map[string]string)
		if ribbonKey.Valid {
			event.MetaData["ribbonKey"] = ribbonKey.String
		}
		if pokemon.Valid {
			event.MetaData["pokemon"] = pokemon.String
		}
		if nickname.Valid {
			event.MetaData["nickname"] = nickname.String
		}
		if ribbonName.Valid {
			event.MetaData["ribbonName"] = ribbonName.String
		}
		if ribbonCategory.Valid {
			event.MetaData["ribbonCategory"] = ribbonCategory.String
		}
		if ribbonType.Valid {
			event.MetaData["ribbonType"] = ribbonType.String
		}
		if pokedex.Valid {
			event.MetaData["pokedex"] = pokedex.String
		}
		eventList = append(eventList, event)
	}
	return eventList, nil
}
