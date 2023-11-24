package wt

import (
	"log"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/webtransport", new(RootModule))
}

type RootModule struct{}

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &Connection{}
)

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {

	metrics, err := registerMetrics(vu)
	if err != nil {
		log.Fatalln(vu.Runtime(), err)
	}

	return &Connection{vu: vu, metrics: metrics}
}

func (conn *Connection) Exports() modules.Exports {
	return modules.Exports{Default: conn}
}
