package	controller
import (
    "fmt"
    "log"
    "database/sql"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
   

)
type Car struct {
	car_id int  `json:"car_id"`
	brand string `json:"brand"`
	model string  `json:"model"`
}


func Select(c *gin.Context) {


	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    res, err := db.Query("SELECT * FROM car")

     defer res.Close()

    if err != nil {
        log.Fatal(err)
    }
    cars := []Car{}
	for res.Next() {

	   car:=Car{}
		err := res.Scan(&car.car_id,&car.brand,&car.model)
		
		if err != nil {
			log.Fatal("fetch error:",err)
		}
		 fmt.Printf("%v\n",car);

        cars = append(cars,car)
        c.JSON(http.StatusOK, gin.H{"car_id":car.car_id,"brand":car.brand,"model":car.model}) 
       
    }
   
   
    
     
        //fmt.Printf("%v\n", city)
    }


func Delete(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
   

    if err != nil {
        log.Fatal(err)
    }
    car_id := c.Query("car_id")

    res, err := db.Query("Delete FROM car where car_id ='"+car_id+"'")

    // defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next() {

        var  car Car
        err := res.Scan(&car.car_id,&car.brand,&car.model)

        if err != nil {
            log.Fatal(err)
        }else{
            c.JSON(200, gin.H{
                "message": "data Deleted Successfully",
            }) 
        }
		
		c.JSON(http.StatusOK, gin.H{
			"message": car,
		})

       

        //fmt.Printf("%v\n", city)
    }
    defer db.Close()
    defer res.Close();
}

func Update(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
   

    if err != nil {
        log.Fatal(err)
    }
    car_id := c.Query("car_id")
    brand := c.Query("brand")
    model :=c.Query("model")
    res, err := db.Query("UPDATE car SET brand='"+brand+"',model='"+model+"' WHERE  car_id='"+car_id+"'");



 

    if err != nil {
        log.Fatal(err)
    }else{
        c.JSON(http.StatusOK, gin.H{
            "message": "data Updated Successfully",
        }) 
    }

    for res.Next() {

        var car Car
        err := res.Scan(&car.car_id,&car.brand,&car.model)

        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%v\n",car)
        c.JSON(http.StatusOK, gin.H{"car_id":car.car_id,"brand":car.brand,"model":car.model})
    }
    defer db.Close()
    defer res.Close()
}



func Create(c *gin.Context) {


	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
  

    if err != nil {
        log.Fatal(err)
    }
    car_id := c.Query("car_id")
    brand :=c.Query("brand")
    model :=c.Query("model")
    res, err := db.Query("INSERT INTO car(car_id,brand,model) VALUES ('"+car_id+"', '"+brand+"','"+model+"');")

   
    if err != nil {
        log.Fatal(err)
    }else{
        c.JSON(http.StatusOK, gin.H{
			"message": "data Inserted Successfully",
		}) 
    }
    defer db.Close()
     defer res.Close()

}    