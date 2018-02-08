package store

import (
	"github.com/owainlewis/frequency/pkg/types"
)

type Store interface {
	GetBuild(id int) (*types.Build, error)
	CreateBuild(build types.Build)
}
