package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SSShipment struct {
	ShipmentID     int     `json:"shipmentId"`
	OrderID        int     `json:"orderId"`
	OrderKey       string  `json:"orderKey"`
	UserID         string  `json:"userId"`        // NOT FJ user id
	CustomerEmail  string  `json:"customerEmail"` // FJ email
	OrderNumber    string  `json:"orderNumber"`   // FJ order
	CreateDate     string  `json:"createDate"`
	ShipDate       string  `json:"shipDate"`
	ShipmentCost   float64 `json:"shipmentCost"`   // important
	TrackingNumber string  `json:"trackingNumber"` // important
	CarrierCode    string  `json:"carrierCode"`    // imporant
	ServiceCode    string  `json:"serviceCode"`
	Confirmation   string  `json:"confirmation"`
	WarehouseID    int     `json:"warehouseId"`
	// See https://www.shipstation.com/docs/api/shipments/list/ for full list
}

type SSResp struct {
	Shipments []SSShipment `json:"shipments,omitempty"`
	Total     int          `json:"total,omitempty"`
	Page      int          `json:"page,omitempty"`
	Pages     int          `json:"pages,omitempty"`
}

func main() {
	username := ""
	password := ""
	req, err := getReq("GET", "https://ssapi.shipstation.com/shipments?shipDateStart=2019-01-01&shipDateEnd=2019-01-07&pageSize=499", username, password)
	if err != nil {
		fmt.Printf("getReq error: %s", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do error: %s", err)
		return
	}

	ssresp := &SSResp{}
	err = decodeResp(resp, ssresp)
	if err != nil {
		fmt.Printf("Read error: %s", err)
		return
	}
	bytes, _ := json.Marshal(ssresp)
	fmt.Printf("%d - %s \n", resp.StatusCode, bytes)
	fmt.Printf("X-Rate-Limit-Limit: %v \n", resp.Header.Get("X-Rate-Limit-Limit"))
	fmt.Printf("X-Rate-Limit-Remaining: %v \n", resp.Header.Get("X-Rate-Limit-Remaining"))
	fmt.Printf("X-Rate-Limit-Reset: %v \n", resp.Header.Get("X-Rate-Limit-Reset"))
}

func getReq(method, url, username, password string) (*http.Request, error) {
	b := new(bytes.Buffer)
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func decodeResp(resp *http.Response, response interface{}) error {
	var err error
	if resp.StatusCode == 204 || strings.Contains(resp.Header.Get("Content-Type"), "No Content") {
		return nil
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "json") {
		var b []byte
		var body string
		b, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			body = string(b)
		}
		return fmt.Errorf("StatusCode(%d) %s", resp.StatusCode, body)
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	return err
}
