package domain

import (
	"net/http"
	"time"
)

var (
	Success             = Status{Code: http.StatusOK, Message: []string{"Success"}}
	BadRequest          = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Not responding because of incorrect syntax"}}
	Unauthorized        = Status{Code: http.StatusUnauthorized, Message: []string{"Sorry, We are not able to process your request. Please try again"}}
	Forbidden           = Status{Code: http.StatusForbidden, Message: []string{"Sorry, Permission denied"}}
	InternalServerError = Status{Code: http.StatusInternalServerError, Message: []string{"Internal Server Error"}}
	ConFlict            = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Data is conflict"}}
	FieldsPermission    = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Fields are not able to update"}}
)

type Result struct {
	OrderDisplay     []Order     `json:"order,omitempty"`
	OrderItemDisplay []OrderItem `json:"order_item,omitempty"`

	CurrentPage *int `json:"current_page,omitempty"`
	PerPage     *int `json:"per_page,omitempty"`
	TotalItem   *int `json:"total_item,omitempty"`
}

type Order struct {
	OrderID          int64       `json:"order_id,omitempty"`
	OrderUUID        string      `json:"order_uuid,omitempty"`
	CustomerID       string      `json:"customer_id,omitempty"`
	VendorID         string      `json:"vendor_id,omitempty"`
	OrderNumber      string      `json:"order_number"`
	CustomerName     string      `json:"customer_name"`
	OrderTotal       float64     `json:"order_total"`
	ShippingPrice    float64     `json:"shipping_price"`
	OrderDate        time.Time   `json:"order_date" `
	NeedTaxInvoice   bool        `json:"need_tax_invoice"`
	TaxInvoiceNumber string      `json:"tax_invoice_number"`
	Status           string      `json:"status"`
	OrderItems       []OrderItem `json:"item,omitempty"`
}

type OrderItem struct {
	OrderID        int64     `json:"order_id,omitempty"`
	OrderUUID      string    `json:"order_uuid,omitempty"`
	ItemID         int64     `json:"item_id,omitempty"`
	ItemUUID       string    `json:"item_uuid,omitempty"`
	VendorID       int64     `json:"vendor_id,omitempty"`
	VendorName     string    `json:"vendor_name"`
	Name           string    `json:"name"`
	Sku            string    `json:"sku"`
	Quantity       float64   `json:"quantity"`
	DiscountAmount float64   `json:"discount_amount"`
	TotalAmount    float64   `json:"total_amount"`
	ShippingPrice  float64   `json:"shipping_price"`
	ShippingName   string    `json:"shipping_name"`
	Status         string    `json:"status"`
	OrderDate      time.Time `json:"order_date" `
}

type MultiLanguage struct {
	TH string `json:"th,omitempty" form:"th"`
	EN string `json:"en,omitempty" form:"en"`
}

type Pagination struct {
	// use 10 limit by default
	Limit  int `json:"limit" query:"limit" validate:"gte=-1,lte=100"`
	Offset int `json:"offset" query:"offset"`
}

type SortMethod struct {
	Asc     bool   `json:"asc" query:"asc"`
	OrderBy string `json:"order_by" query:"order_by"`
}

// ResponseBody struct
type ResponseBody struct {
	Status Status      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`

	CurrentPage *int `json:"current_page,omitempty"`
	PerPage     *int `json:"per_page,omitempty"`
	TotalItem   *int `json:"total_item,omitempty"`
}

// Status struct
type Status struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message,omitempty"`
}
