package rajaongkir

import (
	"encoding/json"
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

type RajaOngkirResponseProvince struct {
	RajaOngkir struct {
		Query struct {
			ID string `json:"id"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func (r *RajaOngkir) GetProvince(idProvince string) (*RajaOngkirResponseProvince, error) {
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

	var response RajaOngkirResponseProvince
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	return &response, nil
}

type RajaOngkirResponseCity struct {
	RajaOngkir struct {
		Query struct {
			ID       string `json:"id"`
			Province string `json:"province"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func (r *RajaOngkir) GetCity(idProvince, idCity string) (*RajaOngkirResponseCity, error) {
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
	var response RajaOngkirResponseCity
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	return &response, nil
}

type RajaOngkirResponseCost struct {
	RajaOngkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func (r *RajaOngkir) GetCost(cityOrigin, cityDestination string, weight int) ([]*RajaOngkirResponseCost, error) {
	var costs []*RajaOngkirResponseCost

	JNE, err := getCostByCourier(cityOrigin, cityDestination, "jne", r.credentialCode, weight)
	if err != nil {
		return nil, err
	}
	POS, err := getCostByCourier(cityOrigin, cityDestination, "pos", r.credentialCode, weight)
	if err != nil {
		return nil, err
	}
	TIKI, err := getCostByCourier(cityOrigin, cityDestination, "tiki", r.credentialCode, weight)
	if err != nil {
		return nil, err
	}

	costs = append(costs, JNE, POS, TIKI)
	return costs, nil
}

func getCostByCourier(cityOrigin, cityDestination, courier, key string, weight int) (*RajaOngkirResponseCost, error) {
	url := "https://api.rajaongkir.com/starter/cost"

	payload := strings.NewReader(fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", cityOrigin, cityDestination, weight, courier))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("key", key)
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

	var response RajaOngkirResponseCost
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
