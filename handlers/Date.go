//package Handler is for creating handlers and bind them to
// a router. It provides an interface between our GO Server
//to our methods
package handlers

import (
	"log"
	"net/http"

	"github.com/avdmsajaykumar/exercise4/data"
)

//Date is a struct on which handler interface is created and
//to invoke methods
type Date struct {
	logger *log.Logger
}

//New Date returns a new Date struct
func NewDate(l *log.Logger) *Date {
	return &Date{l}
}

//ConvertData is method defined on Date struct which is bind to a
//Route "/time" by Mux Router and it returns the response back
//to client
func (d *Date) ConvertDate(rw http.ResponseWriter, r *http.Request) {

	//initial point of actual method invokation to process requests
	d.logger.Println("Processing request")

	//initilize a request struct from data package and use it further
	request := new(data.Request)
	//request body gets parsed to request struct
	err := request.FromJSON(r.Body)

	//returns error if unable to parse json request body
	if err != nil {
		log.Printf("%s \n", err)
		http.Error(rw, "JSON Request is in Wrong format", http.StatusBadRequest)
		return
	}
	//Prints the request struct
	d.logger.Printf("Request := %v", request)

	//Invokes the business logic method and retruns the response
	response := request.GetResponse()

	// Converting reponse struct to Json and retrun it back to Client
	response.ToJSON(rw)
	d.logger.Println("Sent response")
	//end of method invocation
}
