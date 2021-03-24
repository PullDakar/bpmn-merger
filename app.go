package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("diagram_1.bpmn")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened diagram_1.bpmn")

	defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)
	var defs Definitions

	err = xml.Unmarshal(bytes, &defs)
	if err != nil {
		fmt.Println(err)
	}

	firstGraph := NewDirectedGraph()
	start := defs.Process.StartEvent
	serviceTask := defs.Process.ServiceTasks[0]
	end := defs.Process.EndEvent
	startNode := GraphNode{Id: start.Id, Value: start}
	serviceTaskNode := GraphNode{Id: serviceTask.Id, Value: serviceTask}
	endNode := GraphNode{Id: end.Id, Value: end}
	firstGraph.AddEdge(startNode, append([]GraphNode{}, serviceTaskNode))
	firstGraph.AddEdge(serviceTaskNode, append([]GraphNode{}, endNode))

}
