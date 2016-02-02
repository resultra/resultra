package webui

import (
	"appengine"
	"appengine/datastore"
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Parse the templates once
var templates = template.Must(template.ParseFiles("template/main.html"))
var addRowTemplates = template.Must(template.ParseFiles("template/addRow.html"))

type PageInfo struct {
	Title string `json:"title"`
}

type ColumnInfoMap map[string]string

type RowData struct {
	Symbol string `json:"symbol"`
	Qty    string `json:"qty"`
	Price  string `json:"price"`
}

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

type AddRowResult struct {
	AddSucceeded bool
	ErrorMsg     string
}

func dataTable(w http.ResponseWriter, r *http.Request) {
	numCols := 3
	tblCols := make([]ColumnInfoMap, numCols)
	tblCols[0] = ColumnInfoMap{"title": "Symbol"}
	tblCols[1] = ColumnInfoMap{"title": "Qty"}
	tblCols[2] = ColumnInfoMap{"title": "Price"}

	dsCntxt := appengine.NewContext(r)
	query := datastore.NewQuery("Trades")
	var rowData []RowData
	if _, err := query.GetAll(dsCntxt, &rowData); err != nil {
		log.Println("Error retrieving data:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Retrieved Row Data: ", rowData)
	log.Println("Retrieved Row Data length: ", len(rowData))
	numRows := len(rowData)

	// The data comes out of the datastore in a map, with keys and
	// values for each and every row. However, the DataStore object
	// expects the results in a array/slice per row, in the same order
	// as the columns.
	tblData := make([][]string, numRows)
	for tblIndex, tblRow := range rowData {
		tblData[tblIndex] = []string{tblRow.Symbol, tblRow.Qty, tblRow.Price}
	}

	data := DataTablesData{"Test Data", 1, numRows, numRows, tblData, tblCols}

	jsonBuf := new(bytes.Buffer)
	json.NewEncoder(jsonBuf).Encode(data)
	log.Println("Table JSON data: ", jsonBuf.String())

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

func addRow(w http.ResponseWriter, r *http.Request) {

	log.Println("addRow method:", r.Method) //get request method

	if r.Method == "GET" {
		p := PageInfo{"Add Row Page"}
		err := addRowTemplates.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {

		var rawRowData map[string]string
		json.NewDecoder(r.Body).Decode(&rawRowData)
		log.Println("addRow: Raw submitted row info:", rawRowData)

		var rowData RowData
		rowData.Symbol = rawRowData["symbol"]
		rowData.Qty = rawRowData["qty"]
		rowData.Price = rawRowData["price"]
		log.Println("addRow: Structured row info:", rowData)

		appEngCntxt := appengine.NewContext(r)
		rowKey := datastore.NewIncompleteKey(appEngCntxt, "Trades", nil)
		_, err := datastore.Put(appEngCntxt, rowKey, &rowData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(AddRowResult{true, "No Error"})
	}
}
