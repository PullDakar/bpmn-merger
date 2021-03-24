package main

import "encoding/xml"

type Definitions struct {
	XMLName xml.Name `xml:"definitions"`
	Process Process  `xml:"process"`
	Diagram Diagram  `xml:"BPMNDiagram"`
}

// Process block
type Process struct {
	StartEvent    StartEvent     `xml:"startEvent"`
	ServiceTasks  []ServiceTask  `xml:"serviceTask"`
	ScriptTasks   []ScriptTask   `xml:"scriptTask"`
	SequenceFlows []SequenceFlow `xml:"sequenceFlow"`
	EndEvent      EndEvent       `xml:"endEvent"`
}

type StartEvent struct {
	Id         string `xml:"id,attr"`
	OutgoingId string `xml:"outgoing"`
}

type ServiceTask struct {
	Id         string `xml:"id,attr"`
	Class      string `xml:"class,attr"`
	IncomingId string `xml:"incoming"`
	OutgoingId string `xml:"outgoing"`
}

type ScriptTask struct {
}

type SequenceFlow struct {
	Id        string `xml:"id,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}

type EndEvent struct {
	Id         string `xml:"id,attr"`
	IncomingId string `xml:"incoming"`
}

// Process block

// Diagram block
type Diagram struct {
	Id    string `xml:"id,attr"`
	Plane Plane  `xml:"BPMNPlane"`
}

type Plane struct {
	Id          string  `xml:"id,attr"`
	BpmnElement string  `xml:"bpmnElement,attr"`
	Edges       []Edge  `xml:"BPMNEdge"`
	Shapes      []Shape `xml:"BPMNShape"`
}

type Edge struct {
	Id          string     `xml:"id,attr"`
	BpmnElement string     `xml:"bpmnElement,attr"`
	Waypoints   []Waypoint `xml:"waypoint"`
}

type Waypoint struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

type Shape struct {
	Id          string `xml:"id,attr"`
	BpmnElement string `xml:"bpmnElement,attr"`
	Bound       Bound  `xml:"Bounds"`
}

type Bound struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

// Diagram block
