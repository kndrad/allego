package resource

import (
	"allego/request"
	"encoding/json"
	"net/http"
)

func Cart(client *http.Client, accessToken string) (*cart, error) {
	res := newResource[cart]("order/checkout-forms", accessToken)
	resp, err := request.Send(client, res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&res.Response); err != nil {
		return nil, err
	}
	return res.Response, nil
}

func OrderDetails(id string, client *http.Client, accessToken string) (*Details, error) {
	res := newResource[Details]("order/checkout-forms/"+id, accessToken)
	resp, err := request.Send(client, res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&res.Response); err != nil {
		return nil, err
	}
	return res.Response, nil
}

type cart struct {
	Orders []order `json:"checkoutForms"`
}

func (c *cart) FindOrders(name string) []order {
	var orders []order
	for _, order := range c.Orders {
		for _, item := range order.LineItems {
			if item.Offer.Name == name {
				orders = append(orders, order)
			}
		}
	}
	return orders
}

type order struct {
	ID        string `json:"id"`
	LineItems []struct {
		ID    string `json:"id"`
		Offer struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			External struct {
				ID string `json:"id"`
			} `json:"external"`
		} `json:"offer"`
	}
}

type Details struct {
	ID      string `json:"id"`
	Message string `json:"messageToSeller"`
	Buyer   struct {
		Email string `json:"email"`
	} `json:"buyer"`
}
