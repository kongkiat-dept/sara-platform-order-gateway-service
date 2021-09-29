package services

import (
	"sara-platform-order-gateway-service/internal/core/domain"
	"sara-platform-order-gateway-service/internal/core/ports"

	"github.com/sirupsen/logrus"
)

type Service struct {
	storeRepo ports.PlatformOrderHTTPRepository
}

func New(storeRepo ports.PlatformOrderHTTPRepository) *Service {
	return &Service{
		storeRepo: storeRepo,
	}
}

/**
1. get order and order_item from platform order by customer
2. get customer's detail name from customer
3. get shipping's price by order_item from platform shiping
**/
func (svc *Service) GetByCustomer(r domain.OrderRequest) (domain.Result, error) {
	var totalShippingPrice float64
	var response *domain.Result
	var err error
	// 1. get order and order_item from platform order by customer
	if r.CustomerID != nil {
		response, err = svc.storeRepo.GetByCustomerID(r.CustomerID, r.Status)
		if err != nil {

			logrus.Error(err.Error())
			return *response, err
		}
	} else {
		response, err = svc.storeRepo.GetByCustomer(r.Status)
		if err != nil {

			logrus.Error(err.Error())
			return *response, err
		}
	}

	// 2. get customer's Detail name from customer

	// 3. get shipping's price by order_item from platform shiping
	// no loop
	// micro and add
	// talk with kee & p'bow abount logic
	// rabbitMQ in process
	for i, order := range response.OrderDisplay {
		totalShippingPrice = 0
		for _, item := range order.OrderItems {
			//update shiping detail

			//shippingDetail,err := svc.storeRepo.GetShiping(item.ItemID)
			if err != nil {
				return *response, err
			}
			item.ShippingName = "kerry"
			item.ShippingPrice = 0
			totalShippingPrice += item.ShippingPrice
		}
		response.OrderDisplay[i].ShippingPrice = totalShippingPrice
	}

	if response.OrderDisplay == nil {
		orderDisplay := []domain.Order{}
		response.OrderDisplay = orderDisplay
	}

	return *response, nil
}

/**
1. get order and order_item from platform order by vendor
2. get vendor's detail name from custome
3. get shipping's price by order_item from platform shiping
**/
func (svc *Service) GetByVendor(r domain.OrderRequest) (domain.Result, error) {
	var totalShippingPrice float64
	var response *domain.Result
	var err error
	// 1. get order and order_item from platform order by vendor
	if r.VendorID != nil {
		response, err = svc.storeRepo.GetByVendorID(r.VendorID, r.Status)
		if err != nil {
			return *response, err
		}
	} else {
		response, err = svc.storeRepo.GetByVendor(r.Status)
		if err != nil {
			return *response, err
		}
	}

	// 2. get vendor's detail name from vendor

	// 3. get shipping's price by order_item from platform shiping
	// don't have service
	for i, order := range response.OrderDisplay {
		totalShippingPrice = 0
		for _, item := range order.OrderItems {
			//update shiping detail

			//shippingDetail,err := svc.storeRepo.GetShiping(item.ItemID)
			if err != nil {
				return *response, err
			}
			item.ShippingName = "kerry"
			item.ShippingPrice = 0
			totalShippingPrice += item.ShippingPrice
		}
		response.OrderDisplay[i].ShippingPrice = totalShippingPrice
	}

	return *response, nil

}

/**
1. send update status to vendor api
2. update stautus to platform order
**/
func (svc *Service) UpdateStatus(r domain.OrderRequest) (domain.Result, error) {
	// 1. send update status to vendor api

	// 2. update stautus to platform order
	response, err := svc.storeRepo.UpdateStatus(r.OrderID, r.Status)
	if err != nil {
		logrus.Error(err.Error())
		return *response, err
	}

	return *response, nil

}
