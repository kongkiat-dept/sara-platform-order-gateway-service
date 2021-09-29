package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type GetOrderResponse []Order
type GetOrderItemResponse []OrderItem

type Order struct {
	ID                      *int64       `json:"id,omitempty"`
	UUID                    *string      `json:"uuid,omitempty"`
	MemberID                *string      `json:"member_id,omitempty"`
	Name                    *string      `json:"name,omitempty"`
	BillTo                  *string      `json:"bill_to,omitempty"`
	ShipTo                  *string      `json:"ship_to,omitempty"`
	ContactNumber           *string      `json:"contact_number,omitempty"`
	ContactEmail            *string      `json:"contact_email,omitempty"`
	VatRate                 *float64     `json:"vat_rate,omitempty"`
	VatAmount               *float64     `json:"vat_amount,omitempty"`
	PaymentMethod           *string      `json:"payment_method,omitempty"`
	Amount                  *float64     `json:"amount,omitempty"`
	PlatformDiscountAmount  *float64     `json:"platform_discount_amount,omitempty"`
	VendorDiscountAmount    *float64     `json:"vendor_discount_amount,omitempty"`
	TotalItemDiscountAmount *float64     `json:"total_item_discount_amount,omitempty"`
	TotalAmount             *float64     `json:"total_amount,omitempty"`
	RequestTaxInvoice       *bool        `json:"request_tax_invoice,omitempty"`
	Campaigns               *[]Campaigns `json:"campaigns,omitempty"`
	Coupons                 *[]Coupons   `json:"coupons,omitempty"`
	Status                  *string      `json:"status,omitempty"`
	IsActive                *bool        `json:"is_active,omitempty"`
	IsDeleted               *bool        `json:"is_deleted,omitempty"`
	OrderItems              *[]OrderItem `json:"order_items,omitempty"`
	CreatedBy               *string      `json:"created_by,omitempty"`
	CreatedAt               *time.Time   `json:"created_at,omitempty"`
	UpdatedBy               *string      `json:"updated_by,omitempty"`
	UpdatedAt               *time.Time   `json:"updated_at,omitempty"`
}

type OrderItem struct {
	ID             *int64       `json:"id,omitempty"`
	UUID           *string      `json:"uuid,omitempty"`
	OrderID        *int64       `json:"order_id,omitempty"`
	OrderUUID      *string      `json:"order_uuid,omitempty"`
	VendorID       *int64       `json:"vendor_id,omitempty"`
	ShopID         *int64       `json:"shop_id,omitempty"`
	ProductID      *int64       `json:"product_id,omitempty"`
	Sku            *string      `json:"sku,omitempty"`
	Name           *string      `json:"name,omitempty"`
	Quantity       *float64     `json:"quantity,omitempty"`
	PricePerUnit   *float64     `json:"price_per_unit,omitempty"`
	Campaigns      *[]Campaigns `json:"campaigns,omitempty"`
	Coupons        *[]Coupons   `json:"coupons,omitempty"`
	Amount         *float64     `json:"amount,omitempty"`
	DiscountAmount *float64     `json:"discount_amount,omitempty"`
	TotalAmount    *float64     `json:"total_amount,omitempty"`
	Status         *string      `json:"status,omitempty"`
	IsActive       *bool        `json:"is_active,omitempty"`
	IsDeleted      *bool        `json:"is_deleted,omitempty"`
	CreatedBy      *string      `json:"created_by,omitempty"`
	CreatedAt      *time.Time   `json:"created_at,omitempty"`
	UpdatedBy      *string      `json:"updated_by,omitempty"`
	UpdatedAt      *time.Time   `json:"updated_at,omitempty"`
}

type MultiLanguage struct {
	TH string `json:"th,omitempty" form:"th"`
	EN string `json:"en,omitempty" form:"en"`
}

type Coupons struct {
	ID             string  `json:"id,omitempty" form:"id"`
	UUID           string  `json:"uuid,omitempty" form:"uuid"`
	Name           string  `json:"name,omitempty" form:"name"`
	CouponCode     string  `json:"coupon_code,omitempty" form:"coupon_code"`
	DiscountAmount float64 `json:"discount_amount,omitempty" form:"discount_amount"`
}

type Campaigns struct {
	ID             string  `json:"id,omitempty" form:"id"`
	UUID           string  `json:"uuid,omitempty" form:"uuid"`
	Name           string  `json:"name,omitempty" form:"name"`
	DiscountAmount float64 `json:"discount_amount,omitempty" form:"discount_amount"`
}

type SortMethod struct {
	Asc     bool   `json:"asc" query:"asc"`
	OrderBy string `json:"order_by" query:"order_by"`
}

type Pagination struct {
	// use 10 limit by default
	Limit  int `json:"limit" query:"limit" validate:"gte=-1,lte=100"`
	Offset int `json:"offset" query:"offset"`
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

func (r *ResponseBody) BodyParser(inf interface{}) error {
	if r.Data == nil {
		return errors.New("cannot parse nil interface data")
	}
	b, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, inf)
	if err != nil {
		return err
	}

	return nil
}
