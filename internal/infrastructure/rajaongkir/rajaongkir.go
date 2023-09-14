package rajaongkir

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RajaOngkir struct {
	credentialCode string
}

func InitRajaOngkit(credentialCode string) *RajaOngkir {
	return &RajaOngkir{
		credentialCode: credentialCode,
	}
}

func (r *RajaOngkir) GetProvince(idProvince string) (interface{}, error) {
	url := fmt.Sprintf("https://api.rajaongkir.com/starter/province?id=%s", idProvince)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("key", r.credentialCode)
	res, err := http.DefaultClient.Do(req)
	res.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", res.StatusCode)
	}

	// Read the response body into a byte slice
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	formatted := strings.ReplaceAll(string(responseBody), "\"", "")

	return formatted, nil
}

func (r *RajaOngkir) GetCity(idProvince, idCity string) (interface{}, error) {
	url := fmt.Sprintf("https://api.rajaongkir.com/starter/city?id=%s&province=%s", idCity, idProvince)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("key", r.credentialCode)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res.Header.Set("Content-Type", "application/json")
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return string(body), nil
}

func (r *RajaOngkir) GetCost(cityOrigin, cityDestination, weight, courier string) (interface{}, error) {
	url := "https://api.rajaongkir.com/starter/cost"

	payload := strings.NewReader(fmt.Sprintf("origin=%s&destination=%s&weight=%s&courier=%s", cityOrigin, cityDestination, weight, courier))
	fmt.Println(payload)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("key", r.credentialCode)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return string(body), nil
}
