package repository

import (
	"fmt"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type PlaceBD struct {
	db *sqlx.DB
}

func NewPlaceBD(db *sqlx.DB) *PlaceBD {
	return &PlaceBD{
		db: db,
	}
}

func (r PlaceBD) GetPlaceByID(id int) (interface{}, error) {
	// take place types
	query := fmt.Sprintf("SELECT type_id FROM \"%s\" WHERE place_id = $1", placeTypeTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	// looking for housing or events or others
	ind := 3
	for rows.Next() {
		tp := -1
		if err := rows.Scan(&tp); err != nil {
			return nil, err
		}
		if tp == 1 {
			ind = 1
			break
		}
		if tp == 2 {
			ind = 2
			break
		}
	}
	// parse in struct
	// return struct
	switch ind {
	case 1:
		var house ent.Housing
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,house_price,house_type_id,count_room,square FROM \"%s\" WHERE id = $1", placeTable)
		row := r.db.QueryRow(query, id)
		if err := row.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, &house.PlaceInfo.Description, &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, &house.PlaceInfo.Number, &house.HousePrice, &house.HouseTypeId, &house.CountRoom, &house.Square); err != nil {
			return ent.Housing{}, err
		}
		return house, nil
	case 2:
		var event ent.Event
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price FROM \"%s\" WHERE id = $1", placeTable)

		row := r.db.QueryRow(query,id)
		if err := row.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name, &event.PlaceInfo.Description, &event.PlaceInfo.Longitude, &event.PlaceInfo.Latitude, &event.PlaceInfo.Adress, &event.PlaceInfo.Number, &event.Pushkin,&event.Price); err != nil {
			return ent.Event{}, err
		}
		return event, nil
	default:
		var location ent.Location
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price FROM \"%s\" WHERE id = $1", placeTable)
		row := r.db.QueryRow(query,id)
		if err := row.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name, &location.PlaceInfo.Description, &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.PlaceInfo.Number, &location.Pushkin,&location.MinPrice); err != nil {
			return ent.Location{}, err
		}
		return location, nil
	}
}
func (r PlaceBD) GetAllPlaces(placeInd int) (interface{}, error) {
	switch placeInd{
	case 1:
		houses,err := r.getAllHousing()
		if err != nil{
			return nil,err
		}
		return houses,nil
	case 2:
		events,err:=  r.getAllEvents()
		if err != nil{
			return nil,err
		}
		return events,nil
	default:
		locals,err := r.getAllLocations()
		if err != nil{
			return nil,err
		}
		return locals,nil
	}
	
}

func (r PlaceBD) getAllHousing() (*[]ent.Housing,error){
	houses := make([]ent.Housing,0)
	query := fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,house_price,house_type_id,count_room,square FROM \"%s\"", placeTable)
	rows,err := r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var house ent.Housing

		if err := rows.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, &house.PlaceInfo.Description, &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, &house.PlaceInfo.Number, &house.HousePrice, &house.HouseTypeId, &house.CountRoom, &house.Square);err!=nil{
			return nil, err
		}

		houses = append(houses, house)
	}
	return &houses,nil
}


func (r PlaceBD) getAllEvents() (*[]ent.Event,error){
	events := make([]ent.Event,0)
	query := fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price FROM \"%s\"", placeTable)
	rows,err := r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var event ent.Event

		if err := rows.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name, &event.PlaceInfo.Description, &event.PlaceInfo.Longitude, &event.PlaceInfo.Latitude, &event.PlaceInfo.Adress, &event.PlaceInfo.Number, &event.Pushkin,&event.Price);err!=nil{
			return nil, err
		}

		events = append(events, event)
	}
	return &events,nil
}

func (r PlaceBD) getAllLocations() (*[]ent.Location,error){
	locations := make([]ent.Location,0)
	query := fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price\"%s\"", placeTable)
	rows,err := r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var location ent.Location

		if err := rows.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name, &location.PlaceInfo.Description, &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.PlaceInfo.Number, &location.Pushkin,&location.MinPrice);err!=nil{
			return nil, err
		}

		locations = append(locations, location)
	}
	return &locations,nil
}