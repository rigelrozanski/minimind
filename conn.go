package main

const (
	connSize = 5
)

type ElementConnection struct {
	ConnID         uint64
	El1ID          uint64
	El2ID          uint64
	ConnectionKind string
}

// NewElementConnection creates a new ElementConnection object
func NewElementConnection(connID, el1ID, el2ID uint64, connectionKind string) ElementConnection {
	return ElementConnection{
		ConnID:         connID,
		El1ID:          el1ID,
		El2ID:          el2ID,
		ConnectionKind: connectionKind,
	}
}
