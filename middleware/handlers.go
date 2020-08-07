package middleware

import (
	"database/sql"
	"encoding/json"
	"f1api/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func connect() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//login := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=f1api sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// check the conn
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connection successful")
	return db
}

func GetAllCircuits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	circuits, err := getAllCircuits()
	if err != nil {
		log.Fatalf("Error in GetAllCircuits %v", err)
	}

	json.NewEncoder(w).Encode(circuits)

}

func GetCircuit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Error converting %v", err)
	}

	circuit, err := getCircuit(int64(id))
	if err != nil {
		log.Fatalf("Error when getting circuit %v", err)
	}

	json.NewEncoder(w).Encode(circuit)

}

func GetDrivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	drivers, err := getAllDrivers()
	if err != nil {
		log.Fatalf("Cannot get drivers %v", err)
	}
	json.NewEncoder(w).Encode(drivers)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Cannot convert string %v", err)
	}

	driver, err := getDriver(int64(id))
	if err != nil {
		log.Fatalf("cannot find driver %v", err)
	}

	json.NewEncoder(w).Encode(driver)

}

func GetConstructors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	constructor, err := getConstructors()
	if err != nil {
		log.Fatalf("cannot get constructors %v", err)
	}

	json.NewEncoder(w).Encode(constructor)
}

func GetConstructor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Cannot pars string id %v", err)
	}

	constructor, err := getConstructor(int64(id))
	if err != nil {
		log.Fatalf("cannot get constructor %v", err)
	}

	json.NewEncoder(w).Encode(constructor)
}

// ********** HANDLER FUNCTIONS ************** //
func getAllCircuits() ([]models.Circuit, error) {
	db := connect()
	defer db.Close()

	var circuits []models.Circuit

	sqlStatement := `select * from circuits`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Cannot execute query %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var circuit models.Circuit
		err = rows.Scan(&circuit.ID, &circuit.Ref, &circuit.Name, &circuit.Location, &circuit.Country, &circuit.Url)

		if err != nil {
			log.Fatalf("Unable to scan row %v", err)
		}

		circuits = append(circuits, circuit)
	}
	return circuits, err

}

func getCircuit(id int64) (models.Circuit, error) {
	db := connect()
	defer db.Close()

	var circuit models.Circuit

	sqlStatement := `select * from circuits where id=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&circuit.ID, &circuit.Ref, &circuit.Name, &circuit.Location, &circuit.Country, &circuit.Url)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return circuit, nil
	case nil:
		return circuit, nil
	default:
		log.Fatalf("Error: %v", err)
	}
	return circuit, err
}

func getAllDrivers() ([]models.Driver, error) {
	db := connect()
	defer db.Close()

	// the return  val
	var drivers []models.Driver

	sqlStatement := `select * from drivers`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Cannot execute drivers query %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var driver models.Driver
		err := rows.Scan(&driver.ID, &driver.Ref, &driver.Firstname, &driver.Lastname, &driver.Dob, &driver.Nationality, &driver.Url)
		if err != nil {
			log.Fatalf("Cannot get rows for all drivers %v", err)
		}
		drivers = append(drivers, driver)
	}
	return drivers, err
}

func getDriver(id int64) (models.Driver, error) {
	db := connect()
	defer db.Close()

	var driver models.Driver

	sqlStatement := `select * from drivers where id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&driver.ID, &driver.Ref, &driver.Firstname, &driver.Lastname, &driver.Dob, &driver.Nationality, &driver.Url)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return driver, nil
	case nil:
		return driver, nil
	default:
		log.Fatalf("Error: %v", err)
	}

	return driver, err
}

func getConstructors() ([]models.Constructor, error) {
	db := connect()
	defer db.Close()

	sqlStatement := `select * from constructors`

	var constructors []models.Constructor

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("error getting const %v", err)
	}

	for rows.Next() {
		var constructor models.Constructor
		err := rows.Scan(&constructor.ID, &constructor.Ref, &constructor.Name, &constructor.Nationality, &constructor.Url)
		if err != nil {
			log.Fatalf("Cannot get constructor rows %v", err)
		}

		constructors = append(constructors, constructor)
	}

	return constructors, err
}

func getConstructor(id int64) (models.Constructor, error) {
	db := connect()
	defer db.Close()

	sqlStatement := `select * from constructors where id=$1`

	var constructor models.Constructor

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&constructor.ID, &constructor.Ref, &constructor.Name, &constructor.Nationality, &constructor.Url)
	if err != nil {
		log.Fatalf("cannot scan %v", err)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
		return constructor, nil
	case nil:
		return constructor, nil
	default:
		log.Fatalf("Error %v", err)
	}

	return constructor, err
}
