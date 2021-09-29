package handlers

import (
	"sara-platform-order-gateway-service/internal/core/domain"
	"sara-platform-order-gateway-service/internal/core/ports"
	"sara-platform-order-gateway-service/pkg/apprequest"
	"sara-platform-order-gateway-service/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HTTPHandler struct {
	svc        ports.Service
	serviceURL string
	apprequest apprequest.HTTPRequest
	validator  validator.Validator
}

func NewHTTPHandler(svc ports.Service, serviceURL string) *HTTPHandler {
	return &HTTPHandler{
		svc:        svc,
		validator:  validator.New(),
		serviceURL: serviceURL,
		apprequest: apprequest.NewRequester(),
	}
}

func (hdl *HTTPHandler) SearchOrderByCustomer(c *fiber.Ctx) error {
	condition := domain.OrderRequest{}

	err := c.QueryParser(&condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	err = hdl.validator.ValidateStruct(condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	idStr := c.Params("id")
	if idStr != "" {
		condition.CustomerID = &idStr
	}

	result, err := hdl.svc.GetByCustomer(condition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{
		Status:      domain.Success,
		Data:        result.OrderDisplay,
		CurrentPage: result.CurrentPage,
		PerPage:     result.PerPage,
		TotalItem:   result.TotalItem,
	})
}

func (hdl *HTTPHandler) SearchOrderByVendor(c *fiber.Ctx) error {
	condition := domain.OrderRequest{}

	err := c.QueryParser(&condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	err = hdl.validator.ValidateStruct(condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	idStr := c.Params("id")
	if idStr != "" {
		condition.VendorID = &idStr
	}

	result, err := hdl.svc.GetByVendor(condition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{
		Status:      domain.Success,
		Data:        result.OrderItemDisplay,
		CurrentPage: result.CurrentPage,
		PerPage:     result.PerPage,
		TotalItem:   result.TotalItem,
	})
}

func (hdl *HTTPHandler) UpdateStatus(c *fiber.Ctx) error {
	condition := domain.OrderRequest{}
	err := c.BodyParser(&condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})

	}
	err = hdl.validator.ValidateStruct(condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	idStr := c.Params("id")
	if idStr != "" {
		condition.VendorID = &idStr
	}

	result, err := hdl.svc.UpdateStatus(condition)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{
		Status:      domain.Success,
		Data:        result.OrderDisplay,
		CurrentPage: result.CurrentPage,
		PerPage:     result.PerPage,
		TotalItem:   result.TotalItem,
	})
}

func (hdl *HTTPHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{Status: domain.Success, Data: ""})
}
