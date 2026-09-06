package services

import (
	"dezzles-apps/rq-server/model/dto"

	"database/sql"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
)

//go:embed sql/get-events.sql
var getAllEvents string

type EventService struct {
	connection *cdb.Database
}

func NewEventService(db *cdb.Database) *EventService {
	return &EventService{
		connection: db,
	}
}

func (es *EventService) GetLatestEvents() ([]dto.Event, error) {
	eventList := []dto.Event{}
	rows, err := es.connection.GetDB().Query(getAllEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ribbonKey sql.NullString
	var pokemon sql.NullString
	var nickname sql.NullString
	var ribbonName sql.NullString
	var ribbonCategory sql.NullString
	var ribbonType sql.NullString
	for rows.Next() {
		event := dto.Event{}
		err := rows.Scan(
			&event.Type,
			&event.EventTime,
			&pokemon,
			&nickname,
			&ribbonKey,
			&ribbonName,
			&ribbonCategory,
			&ribbonType,
		)
		if err != nil {
			return nil, err
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
		eventList = append(eventList, event)
	}
	return eventList, nil
}
