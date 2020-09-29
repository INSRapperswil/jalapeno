package arangodb

import (
	"strconv"

	"github.com/sbezverk/gobmp/pkg/message"
)

type lsNodeArangoMessage struct {
	*message.LSNode
}

func (n *lsNodeArangoMessage) MakeKey() string {

	// The LSNode Key uses ProtocolID, DomainID, and IGP Area ID
	// to create unique Keys for DB entries in multi-area / multi-topology scenarios
	return strconv.Itoa(int(n.ProtocolID)) + "_" + strconv.Itoa(int(n.DomainID)) + "_" + n.ISISAreaID + "_" + n.OSPFAreaID + "_" + n.IGPRouterID
}
