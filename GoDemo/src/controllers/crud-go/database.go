package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbName   = "testdb"
)
var db *sql.DB
type Car struct {
	car_id int `field:"id"`
	brand string `field:"brand"`
	model string  `field:"model"`
}
func database() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName))
	if err != nil {
		log.Printf("[MYSQL] Error %s when opening DB\n", err)
	} else {
		fmt.Printf("Connected successfully!");
	}
	res, err2 := db.Query("SELECT * FROM car");
	fmt.Println(reflect.TypeOf(res))
	
    // defer res.Close()
	for res.Next() {

		var car Car
		err := res.Scan(&car.car_id,&car.brand,&car.model)
		
		if err != nil {
			log.Fatal("fetch error:",err)
		}
		fmt.Printf("%v\n",car);

	}
	
	if err2 != nil {
		log.Fatal("query error",err2)
	}
	defer res.Close()
	

}
