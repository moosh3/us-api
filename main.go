package main

import (
	"net/http"
)

// StateCode represents data about a record statecode.
type StateCode struct {
	State        string `json:"state"`
	Abbreviation string `json:"abbreviation"`
}

func main() {
	router := gin.Default()

	router.GET("/statecode", getStateCodes)
	router.GET("/statecode/:id", getStateCodeByID)
	router.Run("localhost:8080")
}

// getStateCodes responds with the list of all statecodes as JSON.
func getStateCodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stateCodes)
}

// getAlbumByID locates the state code whose ID value matches the id
// parameter sent by the client, then returns that statecode as a response.
func getStateCodeByID(c *gin.Context) {
	state := c.Param("state")

	// Loop through the list of state codes, looking for
	// a state code whose ID value matches the parameter.
	for _, a := range stateCodes {
		if a.State == state {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "state code not found"})
}
