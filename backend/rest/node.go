package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/prototype/service"
)

type nodeRestAPI struct {
	node_service service.NodeService
}

func NewNodeRestAPI(node_service service.NodeService) nodeRestAPI {
	return nodeRestAPI{node_service: node_service}
}

func (h nodeRestAPI) GetNodeFromNodeID(c *fiber.Ctx) error {
	node_id, err := c.ParamsInt("node_id")
	if err != nil {
		return err
	}
	node, err := h.node_service.GetNode(node_id)
	if err != nil {
		return err
	}
	return c.JSON(node)
}

func (h nodeRestAPI) GetAllNodes(c *fiber.Ctx) error {
	nodes, err := h.node_service.GetAllNodes()
	if err != nil {
		panic(err)
	}
	return c.JSON(nodes)
}

func (h nodeRestAPI) CreateNewNode(c *fiber.Ctx) error {
	var new_nodes service.NodeRequest
	c.BodyParser(&new_nodes)
	node, err := h.node_service.CreateNewNode(new_nodes)
	if err != nil {
		return err
	}
	return c.JSON(node)
}
