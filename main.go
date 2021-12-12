package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// stateCodes slice to seed US State code data.
var stateCodes = []StateCode{
	{State: "Alabama", Abbreviation: "AL"},
	{State: "Alaska", Abbreviation: "AK"},
	{State: "Arizona", Abbreviation: "AZ"},
	{State: "Arkansas", Abbreviation: "AR"},
	{State: "California", Abbreviation: "CA"},
	{State: "Colorado", Abbreviation: "CO"},
	{State: "Connecticut", Abbreviation: "CT"},
	{State: "Delaware", Abbreviation: "DE"},
	{State: "District of Columbia", Abbreviation: "DC"},
	{State: "Florida", Abbreviation: "FL"},
	{State: "Georgia", Abbreviation: "GA"},
	{State: "Hawaii", Abbreviation: "HI"},
	{State: "Idaho", Abbreviation: "ID"},
	{State: "Illinois", Abbreviation: "IL"},
	{State: "Indiana", Abbreviation: "IN"},
	{State: "Iowa", Abbreviation: "IA"},
	{State: "Kansas", Abbreviation: "KS"},
	{State: "Kentucky", Abbreviation: "KY"},
	{State: "Lousiana", Abbreviation: "LA"},
	{State: "Maine", Abbreviation: "ME"},
	{State: "Montana", Abbreviation: "MT"},
	{State: "Nebraska", Abbreviation: "NE"},
	{State: "Nevada", Abbreviation: "NV"},
	{State: "New Hampshire", Abbreviation: "NH"},
	{State: "New Jersey", Abbreviation: "NJ"},
	{State: "New Mexico", Abbreviation: "NM"},
	{State: "New York", Abbreviation: "NY"},
	{State: "North Carolina", Abbreviation: "NC"},
	{State: "North Dakota", Abbreviation: "ND"},
	{State: "Ohio", Abbreviation: "OH"},
	{State: "Oklahoma", Abbreviation: "OK"},
	{State: "Oregon", Abbreviation: "OR"},
	{State: "Maryland", Abbreviation: "MD"},
	{State: "Massachusetts", Abbreviation: "MA"},
	{State: "Michigan", Abbreviation: "MI"},
	{State: "Minnesota", Abbreviation: "MN"},
	{State: "Missouri", Abbreviation: "MO"},
	{State: "Pennsylvania", Abbreviation: "PA"},
	{State: "Rhode Island", Abbreviation: "RI"},
	{State: "South Carolina", Abbreviation: "SC"},
	{State: "South Dakota", Abbreviation: "SD"},
	{State: "Tennessee", Abbreviation: "TN"},
	{State: "Texas", Abbreviation: "TX"},
	{State: "Utah", Abbreviation: "UT"},
	{State: "Vermont", Abbreviation: "VT"},
	{State: "Virginia", Abbreviation: "VA"},
	{State: "Washington", Abbreviation: "WA"},
	{State: "West Virginia", Abbreviation: "WV"},
	{State: "Wisconsin", Abbreviation: "WI"},
	{State: "Wyoming", Abbreviation: "WY"},
}

// StateCode represents data about a statecode.
type StateCode struct {
	State        string `json:"state"`
	Abbreviation string `json:"abbreviation"`
}

func main() {
	router := gin.Default()

	router.GET("/statecode", getStateCodes)
	router.GET("/statecode/:state", getStateCodeByID)
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

	// Loop through the list of state codes, looking for
	// a state code whose ID value matches the parameter.
	for _, a := range stateCodes {
		if a.State == state {
			c.IndentedJSON(http.StatusOK, a.Abbreviation)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "state code not found"})
}
