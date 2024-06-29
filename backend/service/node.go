package service

import "time"

type NodeRequest struct {
	OrderID   int    `json:"order_id"`
	Reference string `json:"reference"`
	Duration  int    `json:"duration"`
}

type NodeResponse struct {
	NodeID         int       `json:"node_id"`
	Reference      string    `json:"reference"`
	OrderID        int       `json:"order_id"`
	IsActive       bool      `json:"is_active"`
	DeploymentDate time.Time `json:"deployement_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type NodeService interface {
	CreateNewNode(NodeRequest) (*NodeResponse, error)
	GetNode(int) (*NodeResponse, error)
	GetAllNodes() ([]NodeResponse, error)
}
