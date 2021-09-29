package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"sara-platform-order-gateway-service/internal/core/domain"
	v1 "sara-platform-order-gateway-service/internal/http/platform"
	"sara-platform-order-gateway-service/pkg/apprequest"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type PlatformOrderHTTP struct {
	serviceURL string
	apprequest apprequest.HTTPRequest
}

func NewPlatformOrderHTTP(serviceURL string, apprequest apprequest.HTTPRequest) *PlatformOrderHTTP {
	return &PlatformOrderHTTP{
		serviceURL: serviceURL,
		apprequest: apprequest,
	}
}

func (repo *PlatformOrderHTTP) GetByCustomer(status *string) (*domain.Result, error) {
	response := domain.Result{}
	url := fmt.Sprintf("%v/v1/api/order", repo.serviceURL)
	switch {
	case status != nil:
		url += fmt.Sprintf("?status=%v", *status)
	}
	req, resp := repo.apprequest.NewRequest(nil, apprequest.GET, url)

	// making HTTP request ...
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			logrus.Error(err.Error())
			return &response, err
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	// data serializing ...
	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return &response, errors.New(string(body))
		}
		serverResponse, err := repo.serverResponseSuccess(body)
		if err != nil {
			return &response, err
		}
		var data v1.GetOrderResponse
		err = serverResponse.BodyParser(&data)
		if err != nil {
			return &response, err
		}

		for i := range data {
			response.OrderDisplay = append(response.OrderDisplay, repo.mapToOrder(data[i]))
		}

		response.CurrentPage = serverResponse.CurrentPage
		response.PerPage = serverResponse.PerPage
		response.TotalItem = serverResponse.TotalItem

	}
	return &response, nil
}

func (repo *PlatformOrderHTTP) GetByCustomerID(customerID *string, status *string) (*domain.Result, error) {
	response := domain.Result{}
	isExist := false
	url := fmt.Sprintf("%v/v1/api/order", repo.serviceURL)
	switch {
	case status != nil:
		url += fmt.Sprintf("?status=%v", *status)
		isExist = true
	case customerID != nil:
		if isExist {
			url += fmt.Sprintf("&member_id=%v", *customerID)
		} else {
			url += fmt.Sprintf("?member_id=%v", *customerID)
			isExist = true
		}

	}
	req, resp := repo.apprequest.NewRequest(nil, apprequest.GET, url)

	// making HTTP request ...
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			logrus.Error(err.Error())
			return &response, err
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	// data serializing ...
	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return &response, errors.New(string(body))
		}
		serverResponse, err := repo.serverResponseSuccess(body)
		if err != nil {
			return &response, err
		}
		var data v1.GetOrderResponse
		err = serverResponse.BodyParser(&data)
		if err != nil {
			return &response, err
		}

		for i := range data {
			response.OrderDisplay = append(response.OrderDisplay, repo.mapToOrder(data[i]))
		}

		response.CurrentPage = serverResponse.CurrentPage
		response.PerPage = serverResponse.PerPage
		response.TotalItem = serverResponse.TotalItem

	}
	return &response, nil
}

func (repo *PlatformOrderHTTP) GetByVendor(status *string) (*domain.Result, error) {
	response := domain.Result{}
	url := fmt.Sprintf("%v/v1/api/order/item", repo.serviceURL)
	switch {
	case status != nil:
		url += fmt.Sprintf("?status=%v", *status)
	}
	req, resp := repo.apprequest.NewRequest(nil, apprequest.GET, url)

	// making HTTP request ...
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			logrus.Error(err.Error())
			return &response, err
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	// data serializing ...
	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return &response, errors.New(string(body))
		}
		serverResponse, err := repo.serverResponseSuccess(body)
		if err != nil {
			return &response, err
		}
		var data v1.GetOrderItemResponse
		err = serverResponse.BodyParser(&data)
		if err != nil {
			return &response, err
		}

		for i := range data {
			response.OrderItemDisplay = append(response.OrderItemDisplay, repo.mapToOrderItem(data[i]))
		}

		response.CurrentPage = serverResponse.CurrentPage
		response.PerPage = serverResponse.PerPage
		response.TotalItem = serverResponse.TotalItem

	}
	return &response, nil
}

func (repo *PlatformOrderHTTP) GetByVendorID(VendorID *string, status *string) (*domain.Result, error) {
	response := domain.Result{}
	isExist := false
	url := fmt.Sprintf("%v/v1/api/order/item", repo.serviceURL)
	switch {
	case status != nil:
		url += fmt.Sprintf("?status=%v", *status)
		isExist = true
	case VendorID != nil:
		if isExist {
			url += fmt.Sprintf("&vendor_id=%v", *VendorID)
		} else {
			url += fmt.Sprintf("?vendor_id=%v", *VendorID)
			isExist = true
		}
	}
	req, resp := repo.apprequest.NewRequest(nil, apprequest.GET, url)

	// making HTTP request ...
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			logrus.Error(err.Error())
			return &response, err
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	// data serializing ...
	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return &response, errors.New(string(body))
		}
		serverResponse, err := repo.serverResponseSuccess(body)
		if err != nil {
			return &response, err
		}
		var data v1.GetOrderItemResponse
		err = serverResponse.BodyParser(&data)
		if err != nil {
			return &response, err
		}

		for i := range data {
			response.OrderItemDisplay = append(response.OrderItemDisplay, repo.mapToOrderItem(data[i]))
		}

		response.CurrentPage = serverResponse.CurrentPage
		response.PerPage = serverResponse.PerPage
		response.TotalItem = serverResponse.TotalItem

	}
	return &response, nil
}

func (repo *PlatformOrderHTTP) UpdateStatus(orderID *int64, status *string) (*domain.Result, error) {
	response := domain.Result{}
	payload := v1.Order{
		ID:     orderID,
		Status: status,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%v/v1/api/order", repo.serviceURL)
	req, resp := repo.apprequest.NewRequest(b, apprequest.PUT, url)
	req.Header.SetContentTypeBytes(apprequest.ApplicationJSON)

	// making HTTP request ...
	{
		err = fasthttp.Do(req, resp)
		if err != nil {
			return nil, err
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	// data serializing ...
	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			return nil, errors.New(string(body))
		}
		serverResponse, err := repo.serverResponseSuccess(body)
		if err != nil {
			return nil, err
		}
		var data v1.GetOrderResponse
		err = serverResponse.BodyParser(&data)
		if err != nil {
			return nil, err
		}

		for i := range data {
			response.OrderDisplay = append(response.OrderDisplay, repo.mapToOrder(data[i]))
		}

		return &response, nil
	}
}

func (r *PlatformOrderHTTP) serverResponseSuccess(b []byte) (v1.ResponseBody, error) {
	var serverResponse v1.ResponseBody
	reader := bytes.NewBuffer(b)
	err := json.NewDecoder(reader).Decode(&serverResponse)
	if err != nil {
		return v1.ResponseBody{}, err
	}
	return serverResponse, nil
}

func (repo *PlatformOrderHTTP) mapToOrder(data v1.Order) domain.Order {
	return domain.Order{
		OrderID:          *data.ID,
		OrderUUID:        *data.UUID,
		CustomerID:       *data.MemberID,
		OrderTotal:       *data.TotalAmount,
		OrderDate:        *data.CreatedAt,
		NeedTaxInvoice:   *data.RequestTaxInvoice,
		TaxInvoiceNumber: "",
		VendorID:         "",
		OrderNumber:      "",
		CustomerName:     "",
		ShippingPrice:    0,
		Status:           *data.Status,
	}
}

func (repo *PlatformOrderHTTP) mapToOrderItem(data v1.OrderItem) domain.OrderItem {
	return domain.OrderItem{
		OrderID: *data.OrderID,
		// OrderUUID:      *data.OrderUUID,
		VendorID:       *data.VendorID,
		ItemID:         *data.ID,
		ItemUUID:       *data.UUID,
		Name:           *data.Name,
		Sku:            *data.Sku,
		Quantity:       *data.Quantity,
		DiscountAmount: *data.DiscountAmount,
		OrderDate:      *data.CreatedAt,
		ShippingPrice:  0,
		TotalAmount:    *data.TotalAmount,
		Status:         *data.Status,
	}
}
