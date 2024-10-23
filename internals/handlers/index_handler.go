package handlers

import (
	"context"
	"htmx-go/internals/services"
	"htmx-go/views"
	"htmx-go/views/components"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const APP_TIMEOUT = time.Second * 10

type IndexHandler struct {
	OrderService *services.OrderService
}

func NewIndexHandler(orderService *services.OrderService) *IndexHandler {
	return &IndexHandler{OrderService: orderService}
}

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (h *IndexHandler) IndexPageHandler(ctx *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), APP_TIMEOUT)
	defer cancel()

	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var viewsOrders []*components.Order
	for _, order := range orders {
		viewsOrder := &components.Order{
			Id: strconv.FormatUint(uint64(order.ID), 10),
			PatientFirstName: order.PatientFirstName,
			PatientLastName: order.PatientFirstName,
			ProviderFirstName: order.ProviderFirstName,
			ProviderLastName: order.ProviderLastName,
		}
		viewsOrders = append(viewsOrders, viewsOrder)
	}
	render(ctx, http.StatusOK, views.Index(viewsOrders))
}
