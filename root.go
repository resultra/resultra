package datasheet

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/pageinfo", pageinfo)
	http.HandleFunc("/dataTable", dataTable)
}

// Parse the templates once
var templates = template.Must(template.ParseFiles("root.html"))

type PageInfo struct {
	Title string `json:"title"`
}

type ColumnInfoMap map[string]string

type DataTablesData struct {
	Title string
	// Fields must be exposed as public with a capital letter
	// to be visible for JSON serialization. The special
	// json:"tag" overrides the field name used for serialization.
	Draw            int             `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	TableData       [][]string      `json:"data"`
	TableColumns    []ColumnInfoMap `json:"columns"`
}

func dataTable(w http.ResponseWriter, r *http.Request) {
	numCols := 2
	tblCols := make([]ColumnInfoMap, numCols)
	tblCols[0] = ColumnInfoMap{"title": "rowInfo"}
	tblCols[1] = ColumnInfoMap{"title": "rowNum"}

	numRows := 3
	tableData := make([][]string, numRows)
	tableData[0] = []string{"r1", "1"}
	tableData[1] = []string{"r2", "2"}
	tableData[2] = []string{"r3", "3"}

	data := DataTablesData{"Test Data", 1, 3, 3, tableData, tblCols}

	json.NewEncoder(w).Encode(data)
}

func pageinfo(w http.ResponseWriter, r *http.Request) {

	info := PageInfo{"Server generated JSON Page Info"}

	json.NewEncoder(w).Encode(info)
}

func root(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	p := PageInfo{"Main Page"}
	err := templates.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
