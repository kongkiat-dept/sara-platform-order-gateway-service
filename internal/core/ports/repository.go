package ports

import "sara-platform-order-gateway-service/internal/core/domain"

type PlatformOrderHTTPRepository interface {
	GetByCustomer(status *string) (*domain.Result, error)
	GetByCustomerID(customerID *string, status *string) (*domain.Result, error)

	GetByVendor(status *string) (*domain.Result, error)
	GetByVendorID(customerID *string, status *string) (*domain.Result, error)

	UpdateStatus(orderID *int64, status *string) (*domain.Result, error)
}
