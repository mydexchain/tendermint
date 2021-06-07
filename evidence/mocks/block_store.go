// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/mydexchain/tendermint/types"
)

// BlockStore is an autogenerated mock type for the BlockStore type
type BlockStore struct {
	mock.Mock
}

// LoadBlockMeta provides a mock function with given fields: height
func (_m *BlockStore) LoadBlockMeta(height int64) *types.BlockMeta {
	ret := _m.Called(height)

	var r0 *types.BlockMeta
	if rf, ok := ret.Get(0).(func(int64) *types.BlockMeta); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockMeta)
		}
	}

	return r0
}
