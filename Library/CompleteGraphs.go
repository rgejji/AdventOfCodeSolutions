package aoc2017

import (
	"fmt"
	"strings"
)

var (
	//assignedCluster is tracks the cluster the given node belongs to
	assignedCluster map[string]string
	//clusterList tracks the nodess that the given cluster id contains.
	clusterList map[string][]string
)

//CountNumberOfClusters counts the number of clusters that belong to given cluster
//connection map lists all the connections from node to list of nodes
func CountNumberOfClusters(connectionMap map[string][]string) int {
	assignedCluster = make(map[string]string)
	clusterList = make(map[string][]string)
	//Trace clusters
	clusterCnt := 0
	for strVal := range connectionMap {
		//If we already assigned the node to a cluster, we have already explored it.
		//So skip
		if _, ok := assignedCluster[strVal]; ok {
			continue
		}
		assignedCluster[strVal] = strVal
		clusterList[strVal] = []string{strVal}
		traceCluster(strVal, strVal, connectionMap)
		clusterCnt++
		//fmt.Printf("Cluster %v has values %v\n", strVal, clusterList[strVal])
	}
	fmt.Printf("Have %d unique clusters\n", clusterCnt)

	return clusterCnt
}

//MakeGraphFromInputStr creates a graph connection map
func MakeGraphFromInputStr(inputStr string) map[string][]string {
	connectionMap := make(map[string][]string)
	lines := strings.Split(inputStr, "\n")
	for _, line := range lines {
		firstSplit := strings.Split(line, "<->")
		value := strings.TrimSpace(firstSplit[0])
		connections := strings.Split(firstSplit[1], ",")

		connectionMap[value] = []string{}
		for _, connection := range connections {
			connectionMap[value] = append(connectionMap[value], strings.TrimSpace(connection))
		}
	}
	return connectionMap
}

func traceCluster(currVal string, clusterVal string, connectionMap map[string][]string) {
	neighbors := connectionMap[currVal]
	for _, neighbor := range neighbors {
		//if already found, continue
		if _, ok := assignedCluster[neighbor]; ok {
			continue
		}
		//otherwise add to list and expand that neighbor
		assignedCluster[neighbor] = clusterVal
		clusterList[clusterVal] = append(clusterList[clusterVal], neighbor)

		traceCluster(neighbor, clusterVal, connectionMap)
	}

}
