package controller

import (
	"golang.org/x/exp/maps"
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}

type ListProviderResponseItem struct {
	Name string `json:"name"`
	Command string `json:"command"`
}
type ListProviderResponse struct {
	Items []ListProviderResponseItem `json:"items"`
}

func ListProviders(c *gin.Context) {
	res := ListProviderResponse{
		Items: make([]ListProviderResponseItem, 0),
	}
	for _, name := range maps.Keys(service.ProviderCommandMap) {
		res.Items = append(res.Items, ListProviderResponseItem{
			Name: name,
		})
	}
	c.JSON(200, res)
}

func DescribeProvider(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)
	info, err := providerSrv.GetInfo()
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Info] {
		Data: info,
	}
	c.JSON(200, res)
}

type AddProviderRequest struct {
	Name string `json:"name" validate:"required"`
	Command string `json:"command" validate:"required"`
}
func AddProvider(c *gin.Context) {
	var reqbody AddProviderRequest
	if err := c.ShouldBindJSON(&reqbody); err != nil {
		AbortOnError(c, err)
		return
	}
	v := validator.New()
	if err := v.Struct(reqbody); err != nil {
		AbortOnError(c, err)
		return
	}
	res := ApiResponse[struct {}] {}
	c.JSON(200, res)
}

func DescribeCard(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	cardName := c.Param("cardName")
	card, err := providerSrv.DescribeCard(cardName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Card] {
		Data: card,
	}
	c.JSON(200, res)
}

func DescribePanel(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	panelName := c.Param("panelName")
	panel, err := providerSrv.DescribePanel(panelName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Panel] {
		Data: panel,
	}
	c.JSON(200, res)
}
