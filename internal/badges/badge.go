package badges

import "github.com/vitorfox/gobadge/package/svg"

type Badger interface {
	Build(map[string]string) *svg.Node
}
