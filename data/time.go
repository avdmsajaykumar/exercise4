//package data is actual package where request and response
// is processed as per needs
package data

import (
	"encoding/json"
	"io"
	"log"
	"regexp"
	"time"
)

//Request is used to bind the json request to a struct
//and the request elements are then used for further processing
type Request struct {
	Date string `json:"Date"` //Date string is gets populated from Date string of JSON request
}

//FromJSON method is used to Parse JSON request to Request struct
func (req *Request) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(req)

}

//GetResponse is the method which process the Request struct
// and returns the Response struct
func (req *Request) GetResponse() *Response {
	//defines a local variable reqdate to
	reqdate := req.Date

	//create a regexp funtion to valiate the input date
	reg := regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}$")

	//retrive the boolean value if date request matches with input
	// date
	flag := reg.MatchString(req.Date)

	//If request satifies regexp. response is being created
	if flag {

		reqdate = reqdate + ":00+00:00"

		//Converts the date string from RFC3339 format and it returns the
		//time struct
		timeFormat, err := time.Parse(
			time.RFC3339,
			reqdate)
		//if errors are observed while convierting string to time struct
		//error is being thrown
		if err != nil {
			log.Printf("Error while parsing request to time format : %v \n", req.Date)
			return nil
		}

		// fetch the month and year from the time struct
		year, month, _ := timeFormat.Date()
		// gets the first day of the corresponding month
		firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, timeFormat.UTC().Location())
		//returns the last day of the same month
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		//Return the final response object
		return &Response{Date: req.Date, LastDay: lastOfMonth.Day()}

	} else {
		return nil
	}

}

//Response is used to create response struct which
//is then passed back to client in json format
type Response struct {
	Date    string `json:"Date"`           //Date string is gets populated back to Date string of JSON response
	LastDay int    `json:"LastDayOfMonth"` //LastDay int type gets populated back to LastDayOfMonth JSON response
}

//ToJSON method is used to convert response struct to JSON
func (res *Response) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(res)

}
