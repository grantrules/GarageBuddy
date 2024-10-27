package internal

import (
	"database/sql"
	"log"
)

func CreateCar(db *sql.DB, c Car) (int64, error) {
	query := `INSERT INTO Cars (name, make, model, year, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id int64
	err := db.QueryRow(query, c.Name, c.Make, c.Model, c.Year, c.UserID).Scan(&id)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	return id, err
}

func CreateServiceRecord(db *sql.DB, sr ServiceRecord) (int64, error) {
	query := `INSERT INTO service_records (car_id, service_date, mileage, service_type, service_description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id int64
	err := db.QueryRow(query, sr.CarID, sr.ServiceDate, sr.Mileage, sr.ServiceType, sr.ServiceDescription).Scan(&id)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	return id, err
}

func CreateMaintenanceItem(db *sql.DB, mi MaintenanceItem) (int64, error) {
	query := `INSERT INTO maintenance_items (service_record_id, item, price)
		VALUES ($1, $2, $3)
		RETURNING id`
	var id int64
	err := db.QueryRow(query, mi.ServiceRecordID, mi.Item, mi.Price).Scan(&id)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	return id, err
}

func RemoveCar(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM Cars WHERE id = $1", id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func ListCarsByUserId(db *sql.DB, userId int) ([]Car, error) {
	rows, err := db.Query("SELECT id, name, make, model, year FROM Cars WHERE user_id = $1", userId)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	cars := []Car{}
	for rows.Next() {
		c := Car{}
		err := rows.Scan(&c.ID, &c.Name, &c.Make, &c.Model, &c.Year)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

type CarWithServiceRecords struct {
	Car            Car
	ServiceRecords []ServiceRecord
}

func GetCarWithServiceRecord(db *sql.DB, car_id string, userId int) (CarWithServiceRecords, error) {

	query := `SELECT c.id, c.name, c.make, c.model, c.year,
		s.id, s.car_id, s.service_date, s.mileage, s.service_type, s.service_description,
		m.id, m.item, m.price

		FROM Cars c
		LEFT JOIN service_records s ON c.id = s.car_id
		LEFT JOIN maintenance_items m ON s.id = m.service_record_id
		WHERE c.id = $1 AND c.user_id = $2`

	var cwr CarWithServiceRecords

	rows, err := db.Query(query, car_id, userId)
	if err != nil {
		log.Print(err)
		return cwr, err
	}

	defer rows.Close()

	c := Car{}

	var last int
	var s ServiceRecord

	for rows.Next() {
		var m MaintenanceItem

		var newSR ServiceRecord

		err := rows.Scan(&c.ID, &c.Name, &c.Make, &c.Model, &c.Year, &newSR.ID, &newSR.CarID, &newSR.ServiceDate, &newSR.Mileage, &newSR.ServiceType, &newSR.ServiceDescription, &m.ID, &m.Item, &m.Price)
		if err != nil {
			log.Print(err)
			return CarWithServiceRecords{}, err
		}

		if last != newSR.ID {
			if last != 0 {
				cwr.ServiceRecords = append(cwr.ServiceRecords, s)
				last = newSR.ID
				s = newSR
			}

			s.MaintenanceItems = append(s.MaintenanceItems, m)
		}
	}

	cwr.ServiceRecords = append(cwr.ServiceRecords, s)
	cwr.Car = c

	return cwr, nil
}
