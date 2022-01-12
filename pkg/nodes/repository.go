package nodes

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type NodeRepository interface {
}

type NodeNeo4jRepository struct {
	Driver neo4j.Driver
}

func (n *NodeNeo4jRepository) Save(node *Node, database string) (err error) {

	session := n.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: database,
	})
	defer func() {
		err = session.Close()
	}()

	if _, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return saveNode(tx, node)
	}); err != nil {
		return err
	}
	return nil
}

func (n *NodeNeo4jRepository) Load(node *Node, database string) (err error) {

	session := n.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: database,
	})
	defer func() {
		err = session.Close()
	}()

	if _, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return listNode(tx)
	}); err != nil {
		return err
	}

	return nil
}

func saveNode(tx neo4j.Transaction, node *Node) (interface{}, error) {
	query := "CREATE (:Node {})"
	parameters := map[string]interface{}{
		"timestamp":         node.Timestamp,
		"nodeId":            node.NodeId,
		"cpuUtilization":    node.CPUUtilization,
		"memoryUtilization": node.MemoryUtilization,
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}

func listNode(tx neo4j.Transaction) (interface{}, error) {
	cypher := "MATCH (:Node {})"
	result, err := tx.Run(cypher, nil)
	if err != nil {
		return nil, err
	}

	var list []Node

	for result.Next() {
		record := result.Record()
		timestamp, _ := record.Get("tiemstamp")
		nodeid, _ := record.Get("nodeid")
		cpuutilization, _ := record.Get("cpuutilization")
		memoryutilization, _ := record.Get("memoryutilization")

		list = append(list, Node{
			Timestamp:         timestamp.(string),
			NodeId:            nodeid.(string),
			CPUUtilization:    cpuutilization.(float64),
			MemoryUtilization: memoryutilization.(float64),
		})
	}

	if err = result.Err(); err != nil {
		return nil, err
	}

	return list, result.Err()
}

// createNode in Neo4j
// e.g. Cypher: "CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node' + id(a)"
func (n *NodeNeo4jRepository) createNode(driver neo4j.Driver, Cypher string, db string) error {

	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: db,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(Cypher, nil)
		if err != nil {
			log.Println("write to DB with error:", err)
			return nil, err
		}
		return result.Consume()
	})

	return err
}

func (n *NodeNeo4jRepository) queryNode(driver neo4j.Driver, Cypher string, db string) ([]neo4j.Node, error) {

	var list []neo4j.Node
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: db,
	})
	defer session.Close()

	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(Cypher, nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			record := result.Record()
			if value, ok := record.Get("n"); ok {
				node := value.(neo4j.Node)
				list = append(list, node)
			}
		}
		if err = result.Err(); err != nil {
			return nil, err
		}

		return list, result.Err()
	})

	if err != nil {
		log.Println("Read error:", err)
	}
	return list, err
}
