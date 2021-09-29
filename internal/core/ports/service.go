package ports

import "sara-platform-order-gateway-service/internal/core/domain"

type Service interface {
	GetByCustomer(domain.OrderRequest) (domain.Result, error)

	GetByVendor(domain.OrderRequest) (domain.Result, error)

	UpdateStatus(domain.OrderRequest) (domain.Result, error)
}
