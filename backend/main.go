package main

// The code begins with importing necessary packages: net/http for HTTP server functionality,
// time for time-related operations,
// github.com/gin-contrib/cors for CORS (Cross-Origin Resource Sharing) support,
// and github.com/gin-gonic/gin for the Gin web framework.
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Define a struct type to represent transaction data
type skeleton struct {
	LenderName   string  `json:"lender"`
	ReceiverName string  `json:"receiver"`
	Amount       float64 `json:"amount"`
	Date         string  `json:"date"`
	//Date is a field of type string that will store a date value.
	// The struct tag json:"date" indicates that this field should be represented as "date" in the JSON object.
}

// var tempdata = []skeleton{
// 	{LenderName: "John", ReceiverName: "Jane", Amount: 100.0, Date: "2023-06-20"},
// }

// //A variable named tempdata is declared, which is a slice of skeleton structs.
// // It initializes the slice with a single entry containing sample transaction data.

func getdata(c *gin.Context) {
	db, err := sql.Open("mysql", "meet115:pass@123@tcp(localhost:3306)/meetdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var data []skeleton

	//The rows.Scan() method is then called, which scans the values from the current row of the result set
	//and assigns them to the corresponding fields of the d struct (LenderName, ReceiverName, Amount, and Date).
	for rows.Next() {
		var d skeleton
		if err := rows.Scan(&d.LenderName, &d.ReceiverName, &d.Amount, &d.Date); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError) // if error occurs code aborts with a HTTP 500
			return
		}
		data = append(data, d)
	}

	c.IndentedJSON(http.StatusOK, data) //sends the data slice as a JSON response with a HTTP 200 OK status code.

}

//The getdata function is a request handler for the GET method.
//It uses the Gin context (gin.Context) to send a JSON response with the tempdata slice as the payload.
//The http.StatusOK code indicates a successful HTTP response.

func adddata(c *gin.Context) {
	var newdata skeleton

	if err := c.BindJSON(&newdata); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db, err := sql.Open("mysql", "meet115:pass@123@tcp(localhost:3306)/meetdb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database successfully connected")
	defer db.Close()

	_, err = db.Exec("INSERT INTO transaction (LenderName, ReceiverName, Amount, Date) VALUES (?, ?, ?, ?)",
		newdata.LenderName, newdata.ReceiverName, newdata.Amount, newdata.Date)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println("database successfully updated")

	c.IndentedJSON(http.StatusCreated, newdata)
}

//The adddata function is a request handler for the POST method.
//It reads the JSON payload from the request body and attempts to bind it to a new main_data struct named newdata.
//If the binding fails, indicating invalid JSON or incompatible types, it responds with a HTTP 400 Bad Request status.
//If the binding is successful, the new transaction data is appended to the tempdata slice, and the updated slice is returned as the response payload with a HTTP 201 Created status.

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:19006"} // Update with your client's origin
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.GET("/transaction_GET", getdata)
	router.POST("/transaction_POST", adddata)
	router.Run("localhost:8080")

}
