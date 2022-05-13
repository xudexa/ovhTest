package app

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"ovhTest/app/functions"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID          string    `gorm:"column:id" json:"id,omitempty"`
	Title       string    `gorm:"column:title" json:"title"`
	Completed   bool      `gorm:"column:completed" json:"completed"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt,omitempty"`
	CompletedAt time.Time `gorm:"column:completedAt" json:"completedAt,omitempty"`
}

// TableName : Database table or collection
func (t Todo) TableName() string {
	return "todos"
}

// TableName : Database table or collection
func (t Todo) Select() string {
	return "id, title, completed, createdAt, completedAt"
}

// Controls : control fields requires
func (t Todo) Controls() error {
	var lstErr string
	rc := "\n"
	lstErr = ""

	if t.Title == "" {
		lstErr += "Todo error : must not be empty" + rc
	}

	if lstErr == "" {
		return nil
	}

	return errors.New(lstErr)

}

//********************************************************
//*                   Query Models                       *
//********************************************************

// QueryParams : HTTP requests parameters
type QueryParams struct {
	ID           string
	Path         string
	TableName    string
	Body         []byte
	SortClause   []string
	Offset       int
	Count        int
	Columns      []string
	SearchClause []string
}

// HTTPParser : Parser pour QueryParams
type HTTPParser interface {
	Parse(*http.Request)
}

// Parse : QueryParams parser
func (q QueryParams) Parse(httpReq *http.Request) QueryParams {
	query := httpReq.URL.Query()
	vars := mux.Vars(httpReq)
	count, _ := strconv.Atoi(query.Get("count"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	path := httpReq.URL.Path
	// Retrieval of items to search
	search := []string{}
	if len(query.Get("search")) > 0 {
		search = strings.Split(query.Get("search"), "+")
		value := search[0]
		tabkeyword := strings.Split(value, " ")
		search = search[:0]
		for i := 0; i < len(tabkeyword); i++ {
			// double the aspostrophes
			keyword := strings.Replace(tabkeyword[i], "'", "''", -1)
			search = append(search, keyword)
		}
	}

	// Sorting recovery
	sorting := []string{}
	if len(query.Get("sort")) > 0 {
		sorting = strings.Split(query.Get("sort"), ",")
	}

	body, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		log.Fatal(err)
	}

	functions.RemoveDuplicate(&sorting)
	functions.RemoveDuplicate(&search)

	var params = QueryParams{
		ID:           vars["id"],
		SortClause:   sorting,
		Body:         body,
		Count:        count,
		Offset:       offset,
		SearchClause: search,
		Path:         path,
	}

	return params
}

//********************************************************
//*                 Response Models                      *
//********************************************************

// WSResponse is the standardized format of the response.
//	+ Meta         : Pre-formatted header of a response returning data.
//	+ Data         : Data or list of data returned.
type WSResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

// MetaResponse is a header of a valid response
//	+ ObjectName  : Information returned to the front allowing him to know which format he receives.
//	+ TotalCount  : Total number of records the application can return.
//	+ Offset      : Starting position of the list of records returned to the Front.
//	+ Count       : Number of records returned to the Front.
type MetaResponse struct {
	ObjectName string `json:"object_name"`
	TotalCount int    `json:"total_count"`
	Offset     int    `json:"offSet"`
	Count      int    `json:"count"`
}
