package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// stateCodes slice to seed US State code data.
var stateCodes = []StateCode{
	{State: "alabama", Abbreviation: "AL"},
	{State: "alaska", Abbreviation: "AK"},
	{State: "arizona", Abbreviation: "AZ"},
	{State: "arkansas", Abbreviation: "AR"},
	{State: "california", Abbreviation: "CA"},
	{State: "colorado", Abbreviation: "CO"},
	{State: "connecticut", Abbreviation: "CT"},
	{State: "delaware", Abbreviation: "DE"},
	{State: "district of columbia", Abbreviation: "DC"},
	{State: "florida", Abbreviation: "FL"},
	{State: "georgia", Abbreviation: "GA"},
	{State: "hawaii", Abbreviation: "HI"},
	{State: "idaho", Abbreviation: "ID"},
	{State: "illinois", Abbreviation: "IL"},
	{State: "indiana", Abbreviation: "IN"},
	{State: "iowa", Abbreviation: "IA"},
	{State: "kansas", Abbreviation: "KS"},
	{State: "kentucky", Abbreviation: "KY"},
	{State: "lousiana", Abbreviation: "LA"},
	{State: "maine", Abbreviation: "ME"},
	{State: "montana", Abbreviation: "MT"},
	{State: "nebraska", Abbreviation: "NE"},
	{State: "nevada", Abbreviation: "NV"},
	{State: "new hampshire", Abbreviation: "NH"},
	{State: "new jersey", Abbreviation: "NJ"},
	{State: "new mexico", Abbreviation: "NM"},
	{State: "new york", Abbreviation: "NY"},
	{State: "north carolina", Abbreviation: "NC"},
	{State: "north dakota", Abbreviation: "ND"},
	{State: "ohio", Abbreviation: "OH"},
	{State: "oklahoma", Abbreviation: "OK"},
	{State: "oregon", Abbreviation: "OR"},
	{State: "maryland", Abbreviation: "MD"},
	{State: "massachusetts", Abbreviation: "MA"},
	{State: "michigan", Abbreviation: "MI"},
	{State: "minnesota", Abbreviation: "MN"},
	{State: "missouri", Abbreviation: "MO"},
	{State: "pennsylvania", Abbreviation: "PA"},
	{State: "rhode island", Abbreviation: "RI"},
	{State: "south carolina", Abbreviation: "SC"},
	{State: "south dakota", Abbreviation: "SD"},
	{State: "tennessee", Abbreviation: "TN"},
	{State: "texas", Abbreviation: "TX"},
	{State: "utah", Abbreviation: "UT"},
	{State: "vermont", Abbreviation: "VT"},
	{State: "virginia", Abbreviation: "VA"},
	{State: "washington", Abbreviation: "WA"},
	{State: "west virginia", Abbreviation: "WV"},
	{State: "wisconsin", Abbreviation: "WI"},
	{State: "wyoming", Abbreviation: "WY"},
}

// StateCode represents data about a statecode.
type StateCode struct {
	State        string `json:"state"`
	Abbreviation string `json:"abbreviation"`
}

func main() {
	router := gin.Default()

	router.GET("/api/v1/statecode", getStateCodes)
	router.GET("/api/v1/statecode/:state", getStateCodeByID)
	router.Run(":8080")
}

// getStateCodes responds with the list of all statecodes as JSON.
func getStateCodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stateCodes)
}

// getStateCodeByID locates the state code whose ID value matches the id
// parameter sent by the client, then returns that statecode as a response.
func getStateCodeByID(c *gin.Context) {
	state := c.Param("state")
	stateQuery := strings.ToLower(state)

	// Loop through the list of state codes, looking for
	// a state code whose ID value matches the parameter.
	for _, a := range stateCodes {
		if a.State == stateQuery {
			c.IndentedJSON(http.StatusOK, a.Abbreviation)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "statecode not found"})
}
