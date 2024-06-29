package service

import (
	"fmt"
	"time"

	"github.com/sebsvt/prototype/repository"
)

type nodeService struct {
	node_repo repository.NodeRepository
}

func NewNodeService(node_repo repository.NodeRepository) NodeService {
	return nodeService{node_repo: node_repo}
}

// CreateNewNode implements ApplicationService.
func (srv nodeService) CreateNewNode(new_node NodeRequest) (*NodeResponse, error) {
	fmt.Println(new_node)
	entity, err := srv.node_repo.Create(repository.Node{
		OrderID:        new_node.OrderID,
		Reference:      new_node.Reference,
		Duration:       new_node.Duration,
		IsActive:       true,
		DeploymentDate: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	expiration_date := entity.DeploymentDate.AddDate(0, entity.Duration, 0)
	return &NodeResponse{
		NodeID:         entity.NodeID,
		Reference:      entity.Reference,
		OrderID:        entity.OrderID,
		IsActive:       entity.IsActive,
		DeploymentDate: entity.DeploymentDate,
		ExpirationDate: expiration_date,
	}, nil
}

// GetAllNodes implements ApplicationService.
func (srv nodeService) GetAllNodes() ([]NodeResponse, error) {
	var nodes_res []NodeResponse
	all_nodes, err := srv.node_repo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, node := range all_nodes {
		nodes_res = append(nodes_res, NodeResponse{
			NodeID:         node.NodeID,
			Reference:      node.Reference,
			OrderID:        node.OrderID,
			IsActive:       node.IsActive,
			DeploymentDate: node.DeploymentDate,
			ExpirationDate: node.DeploymentDate.AddDate(0, node.Duration, 0),
		})
	}
	return nodes_res, nil
}

// GetNode implements ApplicationService.
func (srv nodeService) GetNode(node_id int) (*NodeResponse, error) {
	node, err := srv.node_repo.GetByID(node_id)
	if err != nil {
		return nil, err
	}
	return &NodeResponse{
		NodeID:         node.NodeID,
		Reference:      node.Reference,
		OrderID:        node.OrderID,
		IsActive:       node.IsActive,
		DeploymentDate: node.DeploymentDate,
		ExpirationDate: node.DeploymentDate.AddDate(0, node.Duration, 0),
	}, nil
}
