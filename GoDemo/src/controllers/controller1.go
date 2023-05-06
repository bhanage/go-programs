package	controller
type City struct {
    Id         int
    Name       string
    Population int
}


func Select(c *gin.Context) {


	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    res, err := db.Query("SELECT * FROM cities")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next() {

        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }
		
		c.JSON(200, gin.H{
			"message": city,
		})


        //fmt.Printf("%v\n", city)
    }
}


func Delete(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }
    id := c.Query("id")

    res, err := db.Query("Delete FROM cities where Id ="+id)

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next() {

        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }else{
            c.JSON(200, gin.H{
                "message": "data Deleted Successfully",
            }) 
        }
		
		c.JSON(200, gin.H{
			"message": city,
		})


        //fmt.Printf("%v\n", city)
    }
}

func Update(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }
    id := c.Query("id")
    name := c.Query("name")
    population :=c.Query("population")
    res, err := db.Query("UPDATE cities set name="+name+",population="+population+" WHERE  id ="+id)



    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }else{
        c.JSON(200, gin.H{
            "message": "data Updated Successfully",
        }) 
    }

    for res.Next() {

        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)

        if err != nil {
            log.Fatal(err)
        }
    }
}



func create(c *gin.Context) {


	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }
    id := c.Query("id")
    name :=c.Query("name")
    population :=c.Query("population")
    res, err := db.Query("INSERT INTO cities(id,name, population) VALUES ('"+id+"', "+name+","+population+")")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }else{
        c.JSON(200, gin.H{
			"message": "data Inserted Successfully",
		}) 
    }

}    