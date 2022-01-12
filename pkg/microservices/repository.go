package microservices

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type MSRepository interface {
}

type MSNeo4jRepository struct {
	Driver neo4j.Driver
}

func (m *MSNeo4jRepository) saveMicroService(tx neo4j.Transaction, ms *MicroService) (interface{}, error) {
	query := "CREATE (:MicroService {})"
	parameters := map[string]interface{}{
		"timestamp":         ms.Timestamp,
		"msname":            ms.MsName,
		"msinstanceid":      ms.MsInstanceId,
		"nodeid":            ms.NodeId,
		"cpuutilization":    ms.CPUUtilization,
		"memoryutilization": ms.MemoryUtilization,
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}

func (m *MSNeo4jRepository) listMicroService(tx neo4j.Transaction) (interface{}, error) {
	cypher := "MATCH (:MicroService {})"
	result, err := tx.Run(cypher, nil)
	if err != nil {
		return nil, err
	}

	var list []MicroService

	for result.Next() {
		record := result.Record()
		timestamp, _ := record.Get("timestamp")
		msname, _ := record.Get("msname")
		msinstanceid, _ := record.Get("msinstanceid")
		nodeid, _ := record.Get("nodeid")
		cpuutilization, _ := record.Get("cpuutilization")
		memoryutilization, _ := record.Get("memoryutilization")

		list = append(list, MicroService{
			Timestamp:         timestamp.(string),
			MsName:            msname.(string),
			MsInstanceId:      msinstanceid.(string),
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
