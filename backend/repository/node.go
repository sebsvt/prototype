package repository

import "time"

type Node struct {
	NodeID         int       `db:"node_id"`
	Reference      string    `db:"reference"`
	OrderID        int       `db:"order_id"`
	Duration       int       `db:"duration"`
	IsActive       bool      `db:"is_active"`
	DeploymentDate time.Time `db:"deployment_date"`
}

type NodeRepository interface {
	GetByID(node_id int) (*Node, error)
	Create(Node) (*Node, error)
	GetAll() ([]Node, error)
}
