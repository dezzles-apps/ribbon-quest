package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model/dto"
	"errors"
	"log"
	"time"

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

func (es *EventService) readEvents(rows *sql.Rows) ([]dto.Event, error) {
	eventList := []dto.Event{}
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

func (es *EventService) GetLatestEvents(date string) ([]dto.Event, error) {
	var rows *sql.Rows
	var err error

	if date != "" {
		parsedDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			log.Printf("Date issue: %s", err.Error())
			return nil, errors.New("Invalid key")
		}
		log.Printf("Key: %s", parsedDate)
		rows, err = es.connection.GetDB().Query("SELECT * FROM events_v1 WHERE timestamp < ? LIMIT 10", parsedDate)
	} else {
		rows, err = es.connection.GetDB().Query("SELECT * FROM events_v1 LIMIT 10")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return es.readEvents(rows)
}
