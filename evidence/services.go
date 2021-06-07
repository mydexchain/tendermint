package evidence

import (
	"github.com/mydexchain/tendermint/types"
)

//go:generate mockery --case underscore --name BlockStore

type BlockStore interface {
	LoadBlockMeta(height int64) *types.BlockMeta
}
