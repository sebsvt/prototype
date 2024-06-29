package repository

import "github.com/jmoiron/sqlx"

type nodeRepositoryDB struct {
	db *sqlx.DB
}

func NewNodeRepositoryDB(db *sqlx.DB) NodeRepository {
	return nodeRepositoryDB{db: db}
}

// Create implements NodeRepository.
func (repo nodeRepositoryDB) Create(new_node Node) (*Node, error) {
	query := "insert into nodes (order_id, reference, duration, is_active, deployment_date) values (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, new_node.OrderID, new_node.Reference, new_node.Duration, new_node.IsActive, new_node.DeploymentDate)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	new_node.NodeID = int(id)
	return &new_node, nil
}

// GetAll implements NodeRepository.
func (repo nodeRepositoryDB) GetAll() ([]Node, error) {
	var nodes []Node
	query := "select node_id, reference, order_id, duration, is_active, deployment_date from nodes"
	err := repo.db.Select(&nodes, query)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetByID implements NodeRepository.
func (repo nodeRepositoryDB) GetByID(node_id int) (*Node, error) {
	var node Node
	query := "select node_id, reference, order_id, duration, is_active, deployment_date from nodes where node_id=?"
	err := repo.db.Get(&node, query, node_id)
	if err != nil {
		return nil, err
	}
	return &node, nil
}
