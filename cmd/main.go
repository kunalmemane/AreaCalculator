package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kunalmemane9150/AreaCalculator/internal/handler"
)

func init() {
	banner := `
	
 _______  ______    _______  _______    _______  _______  ___      _______  __   __  ___      _______  _______  _______  ______   
|   _   ||    _ |  |       ||   _   |  |       ||   _   ||   |    |       ||  | |  ||   |    |   _   ||       ||       ||    _ |  
|  |_|  ||   | ||  |    ___||  |_|  |  |       ||  |_|  ||   |    |       ||  | |  ||   |    |  |_|  ||_     _||   _   ||   | ||  
|       ||   |_||_ |   |___ |       |  |       ||       ||   |    |       ||  |_|  ||   |    |       |  |   |  |  | |  ||   |_||_ 
|       ||    __  ||    ___||       |  |      _||       ||   |___ |      _||       ||   |___ |       |  |   |  |  |_|  ||    __  |
|   _   ||   |  | ||   |___ |   _   |  |     |_ |   _   ||       ||     |_ |       ||       ||   _   |  |   |  |       ||   |  | |
|__| |__||___|  |_||_______||__| |__|  |_______||__| |__||_______||_______||_______||_______||__| |__|  |___|  |_______||___|  |_|

	`
	fmt.Println(banner)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/getArea", handler.GetAreaHandler)
	http.HandleFunc("/getPerimeter", handler.GetAreaHandler)

	fmt.Printf("Server started at port 8080\n")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// curl --header "Content-Type: application/json"   --request POST   --data '{
// 	"shape":"Square",
// 	"side":20
//   }'   http://localhost:8080/getArea
//   {"area":"400.000","perimeter":"80.000"}
