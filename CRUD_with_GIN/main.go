package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Truck struct {
	Id      int    `json:id`
	DriverName    string `json:drivername`
	TruckName   string `json:truckname`
	CleanerName string `json:cleanername`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "order_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {

		drivername := c.Query("drivername")
		truckname := c.Query("truckname")
		cleanername := c.Query("cleanername")

		c.JSON(200, gin.H{
			"drivername":  drivername,
			"truckname":   truckname,
			"cleanername": cleanername,
		})
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO truck(drivername, truckname, cleanername) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(drivername, truckname, cleanername)
		fmt.Printf("drivername: %s; truckname: %s; quality: %s", drivername, truckname, cleanername)
	})

	router.PUT("/update", func(c *gin.Context) {
		id1 := c.Query("id")
		drivername := c.Query("drivername")
		truckname := c.Query("truckname")
		cleanername := c.Query("cleanername")

		c.JSON(200, gin.H{
			"drivername":  drivername,
			"truckname":   truckname,
			"cleanername": cleanername,
		})
		db := dbConn()
		upForm, err := db.Prepare("UPDATE truck SET drivername=?, truckname=?, cleanername=? Where id=?")
		if err != nil {
			panic(err.Error())
		}
		upForm.Exec(drivername, truckname, cleanername, id1)
		fmt.Printf("drivername: %s; truckname: %s; cleanername: %s", drivername, truckname, cleanername)
	})

	router.GET("/GET", func(c *gin.Context) {
		id := c.Query("id")
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM truck WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}

		var drivername, truckname, cleanername string
		for selDB.Next() {

			err = selDB.Scan(&id, &drivername, &truckname, &cleanername)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("drivername: %s; truckname: %s; cleanername: %s", drivername, truckname, cleanername)

		c.JSON(200, gin.H{
			"id":      id,
			"drivername":  drivername,
			"truckname":   truckname,
			"cleanername": cleanername,
		})

	})

	router.DELETE("/delete", func(c *gin.Context) {
		var t truck
		if c.BindJSON(&p) == nil {
			db := dbConn()
			delForm, err := db.Prepare("DELETE FROM truck WHERE name=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(t.DriverName)
			log.Println("DELETE")
			defer db.Close()
		}

	})

	router.Run(":8080")
}
