package controller

import (
	"github.com/foxfurry/go_delivery/internal/domain/dto"
	"github.com/foxfurry/go_delivery/internal/service/supervisor"
	"github.com/gin-gonic/gin"
)

type IController interface {
	RegisterDeliveryRoutes(c *gin.Engine)
}

type deliveryController struct {
	supervisor supervisor.ISupervisor
}

func NewDeliveryController() IController {
	return &deliveryController{
		supervisor: supervisor.NewDeliverySupervisor(),
	}
}

func (ctrl *deliveryController) RegisterDeliveryRoutes(c *gin.Engine){
	c.POST("/order", ctrl.order)
	c.GET("/menu", ctrl.menu)
	c.POST("/distribution", ctrl.distribute)
}

func (ctrl *deliveryController) menu(c *gin.Context){
	menu, err := ctrl.supervisor.Menu()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, menu)
}

func (ctrl *deliveryController) order(c *gin.Context){
	order := new(dto.ClientOrder)

	if err := c.ShouldBindJSON(order); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}

	if err := ctrl.supervisor.Order(order); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}

	c.JSON(200, "Received")
}

func (ctrl *deliveryController) distribute(c *gin.Context){
	order := new(dto.Distribution)

	if err := c.ShouldBindJSON(order); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}

	if err := ctrl.supervisor.Distribution(*order); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	}

	c.JSON(200, "Received")
}