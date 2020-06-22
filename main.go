//Package main is used to invoke a Rest API so fetch the
//request date and return the end date in response
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avdmsajaykumar/exercise4/handlers"
	"github.com/gorilla/mux"
)

func main() {

	//create a logger point and this is used for logging the application logs
	logger := log.New(os.Stdout, "Exercise - 4:", log.LstdFlags)

	//create a Handler which later used to invoke requests
	//passed to an particular URI
	dateHandler := handlers.NewDate(logger)

	// initialize a mux router to process the request messages over HTTP
	router := mux.NewRouter()

	//listen the messages receiving on to the /time URI and pass them
	//to ConverDate method of date handler
	router.HandleFunc("/time", dateHandler.ConvertDate).Methods(http.MethodPost)

	//initilize a go server struct and bind to mux router
	//along with Read and Write timeouts
	server := http.Server{
		ReadTimeout:  5 * time.Second,  //Read timeout value set to 120 sec for any request
		Handler:      router,           //Binds the mux router to http server
		WriteTimeout: 10 * time.Second, //Write timeout value set to 120 sec for any request
		IdleTimeout:  60 * time.Second, //Idle timeout value set to 120 sec for any request
		Addr:         ":80",            // Addr on which service is listened
	}

	//starts the server and listen he overall incoming messages
	// and shuts down if there are any fatal errors
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error : %v", err)
	}
}
