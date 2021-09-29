package domain

type OrderRequest struct {
	OrderID    *int64  `json:"order_id" form:"order_id"`
	OrderUUID  *string `json:"order_uuid" form:"order_uuid"`
	CustomerID *string `json:"customer_id" form:"customer_id"`
	VendorID   *string `json:"vendor_id" form:"vendor_id"`
	Status     *string `json:"status" form:"status"`
}
