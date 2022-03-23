package src

import (
	conn "covid_api/database"
	m "covid_api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/jasonwinn/geocoder"

	"log"
)

func SaveCovidData(c echo.Context) error {

	// Open our jsonFile
	jsonFile, err := os.Open("data.min.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("\n\nSuccessfully Opened users.json\n\n")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	docs := []m.MongoFields{}
	// doc := m.StateData{}
	var doc []interface{}

	for key := range result {
		n := m.MongoFields{State: key, TotalCases: result[key].(map[string]interface{})["total"].(map[string]interface{})["confirmed"].(float64)}
		docs = append(docs, n)
		doc = append(doc, n)
	}

	// doc.StateCases = docs
	fmt.Println("printintg doc", doc)

	conn.Connect(doc)
	return c.JSON(http.StatusOK, "success")
}

func GetStateName(c echo.Context) error {

	var location m.Location

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &location)
	if err != nil {
		log.Printf("Failed unmarshal  %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	fmt.Println(location)
	geocoder.SetAPIKey("Fmjtd%7Cluub256alu%2C7s%3Do5-9u82ur")
	address, err := geocoder.ReverseGeocode(location.Lat, location.Long)
	if err != nil {
		panic("THERE WAS SOME ERROR!!!!!")
	}

	fmt.Println(address.State)

	stateMapping := map[string]string{
		"Andaman and Nicobar Islands": "AN",
		"Andhra Pradesh":              "AP",
		"Arunachal Pradesh":           "AR",
		"Assam":                       "AS",
		"Bihar":                       "BR",
		"Chandigarh":                  "CH",
		"Chhattisgarh":                "CT",
		"Dadra and Nagar Haveli":      "DN",
		"Daman and Diu":               "DD",
		"Delhi":                       "DL",
		"Goa":                         "GA",
		"Gujarat":                     "GJ",
		"Haryana":                     "HR",
		"Himachal Pradesh":            "HP",
		"Jammu and Kashmir":           "JK",
		"Jharkhand":                   "JH",
		"Karnataka":                   "KA",
		"Kerala":                      "KL",
		"Lakshadweep":                 "LD",
		"Madhya Pradesh":              "MP",
		"Maharashtra":                 "MH",
		"Manipur":                     "MN",
		"Meghalaya":                   "ML",
		"Mizoram":                     "MZ",
		"Nagaland":                    "NL",
		"Odisha":                      "OR",
		"Orissa":                      "OD",
		"cherry":                      "PY",
		"Punjab":                      "PB",
		"Rajasthan":                   "RJ",
		"Sikkim":                      "SK",
		"Tamil Nadu":                  "TN",
		"Telangana":                   "TG",
		"Tripura":                     "TR",
		"Uttar Pradesh":               "UP",
		"Uttarakhand":                 "UT",
		"West Bengal":                 "WB",
	}

	if v, found := stateMapping[address.State]; found {

		fmt.Println("state code: ", v)
		fetchedData := conn.ConnectAndGet(v)

		responseData := &m.MongoFields{
			State:      fetchedData.State,
			TotalCases: fetchedData.TotalCases,
		}
		fmt.Println("response: ", *responseData)
		return c.JSON(http.StatusOK, responseData)
	} else {
		fmt.Println("Coordinates are not belongs to India")
	}

	return c.JSON(http.StatusOK, "Coordinates are not belongs to India")
}
