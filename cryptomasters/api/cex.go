package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"Mohcode.com/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) { //use pointer to change the var directly
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	//var wg.Waitgroup

	if res.StatusCode == http.StatusOK {
		bodybytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodybytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received : %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil

}
