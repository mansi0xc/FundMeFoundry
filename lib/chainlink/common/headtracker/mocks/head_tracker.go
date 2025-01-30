// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
)

// HeadTracker is an autogenerated mock type for the HeadTracker type
type HeadTracker[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	mock.Mock
}

type HeadTracker_Expecter[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	mock *mock.Mock
}

func (_m *HeadTracker[H, BLOCK_HASH]) EXPECT() *HeadTracker_Expecter[H, BLOCK_HASH] {
	return &HeadTracker_Expecter[H, BLOCK_HASH]{mock: &_m.Mock}
}

// Backfill provides a mock function with given fields: ctx, headWithChain
func (_m *HeadTracker[H, BLOCK_HASH]) Backfill(ctx context.Context, headWithChain H) error {
	ret := _m.Called(ctx, headWithChain)

	if len(ret) == 0 {
		panic("no return value specified for Backfill")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, H) error); ok {
		r0 = rf(ctx, headWithChain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HeadTracker_Backfill_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Backfill'
type HeadTracker_Backfill_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// Backfill is a helper method to define mock.On call
//   - ctx context.Context
//   - headWithChain H
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) Backfill(ctx interface{}, headWithChain interface{}) *HeadTracker_Backfill_Call[H, BLOCK_HASH] {
	return &HeadTracker_Backfill_Call[H, BLOCK_HASH]{Call: _e.mock.On("Backfill", ctx, headWithChain)}
}

func (_c *HeadTracker_Backfill_Call[H, BLOCK_HASH]) Run(run func(ctx context.Context, headWithChain H)) *HeadTracker_Backfill_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(H))
	})
	return _c
}

func (_c *HeadTracker_Backfill_Call[H, BLOCK_HASH]) Return(err error) *HeadTracker_Backfill_Call[H, BLOCK_HASH] {
	_c.Call.Return(err)
	return _c
}

func (_c *HeadTracker_Backfill_Call[H, BLOCK_HASH]) RunAndReturn(run func(context.Context, H) error) *HeadTracker_Backfill_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *HeadTracker[H, BLOCK_HASH]) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HeadTracker_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type HeadTracker_Close_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) Close() *HeadTracker_Close_Call[H, BLOCK_HASH] {
	return &HeadTracker_Close_Call[H, BLOCK_HASH]{Call: _e.mock.On("Close")}
}

func (_c *HeadTracker_Close_Call[H, BLOCK_HASH]) Run(run func()) *HeadTracker_Close_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeadTracker_Close_Call[H, BLOCK_HASH]) Return(_a0 error) *HeadTracker_Close_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_Close_Call[H, BLOCK_HASH]) RunAndReturn(run func() error) *HeadTracker_Close_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *HeadTracker[H, BLOCK_HASH]) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// HeadTracker_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type HeadTracker_HealthReport_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) HealthReport() *HeadTracker_HealthReport_Call[H, BLOCK_HASH] {
	return &HeadTracker_HealthReport_Call[H, BLOCK_HASH]{Call: _e.mock.On("HealthReport")}
}

func (_c *HeadTracker_HealthReport_Call[H, BLOCK_HASH]) Run(run func()) *HeadTracker_HealthReport_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeadTracker_HealthReport_Call[H, BLOCK_HASH]) Return(_a0 map[string]error) *HeadTracker_HealthReport_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_HealthReport_Call[H, BLOCK_HASH]) RunAndReturn(run func() map[string]error) *HeadTracker_HealthReport_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// LatestAndFinalizedBlock provides a mock function with given fields: ctx
func (_m *HeadTracker[H, BLOCK_HASH]) LatestAndFinalizedBlock(ctx context.Context) (H, H, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestAndFinalizedBlock")
	}

	var r0 H
	var r1 H
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (H, H, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) H); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(H)
	}

	if rf, ok := ret.Get(1).(func(context.Context) H); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(H)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// HeadTracker_LatestAndFinalizedBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestAndFinalizedBlock'
type HeadTracker_LatestAndFinalizedBlock_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// LatestAndFinalizedBlock is a helper method to define mock.On call
//   - ctx context.Context
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) LatestAndFinalizedBlock(ctx interface{}) *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH] {
	return &HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH]{Call: _e.mock.On("LatestAndFinalizedBlock", ctx)}
}

func (_c *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH]) Run(run func(ctx context.Context)) *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH]) Return(latest H, finalized H, err error) *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH] {
	_c.Call.Return(latest, finalized, err)
	return _c
}

func (_c *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH]) RunAndReturn(run func(context.Context) (H, H, error)) *HeadTracker_LatestAndFinalizedBlock_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// LatestChain provides a mock function with given fields:
func (_m *HeadTracker[H, BLOCK_HASH]) LatestChain() H {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LatestChain")
	}

	var r0 H
	if rf, ok := ret.Get(0).(func() H); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(H)
	}

	return r0
}

// HeadTracker_LatestChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestChain'
type HeadTracker_LatestChain_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// LatestChain is a helper method to define mock.On call
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) LatestChain() *HeadTracker_LatestChain_Call[H, BLOCK_HASH] {
	return &HeadTracker_LatestChain_Call[H, BLOCK_HASH]{Call: _e.mock.On("LatestChain")}
}

func (_c *HeadTracker_LatestChain_Call[H, BLOCK_HASH]) Run(run func()) *HeadTracker_LatestChain_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeadTracker_LatestChain_Call[H, BLOCK_HASH]) Return(_a0 H) *HeadTracker_LatestChain_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_LatestChain_Call[H, BLOCK_HASH]) RunAndReturn(run func() H) *HeadTracker_LatestChain_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *HeadTracker[H, BLOCK_HASH]) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// HeadTracker_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type HeadTracker_Name_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) Name() *HeadTracker_Name_Call[H, BLOCK_HASH] {
	return &HeadTracker_Name_Call[H, BLOCK_HASH]{Call: _e.mock.On("Name")}
}

func (_c *HeadTracker_Name_Call[H, BLOCK_HASH]) Run(run func()) *HeadTracker_Name_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeadTracker_Name_Call[H, BLOCK_HASH]) Return(_a0 string) *HeadTracker_Name_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_Name_Call[H, BLOCK_HASH]) RunAndReturn(run func() string) *HeadTracker_Name_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *HeadTracker[H, BLOCK_HASH]) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HeadTracker_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type HeadTracker_Ready_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) Ready() *HeadTracker_Ready_Call[H, BLOCK_HASH] {
	return &HeadTracker_Ready_Call[H, BLOCK_HASH]{Call: _e.mock.On("Ready")}
}

func (_c *HeadTracker_Ready_Call[H, BLOCK_HASH]) Run(run func()) *HeadTracker_Ready_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeadTracker_Ready_Call[H, BLOCK_HASH]) Return(_a0 error) *HeadTracker_Ready_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_Ready_Call[H, BLOCK_HASH]) RunAndReturn(run func() error) *HeadTracker_Ready_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *HeadTracker[H, BLOCK_HASH]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HeadTracker_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type HeadTracker_Start_Call[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable] struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *HeadTracker_Expecter[H, BLOCK_HASH]) Start(_a0 interface{}) *HeadTracker_Start_Call[H, BLOCK_HASH] {
	return &HeadTracker_Start_Call[H, BLOCK_HASH]{Call: _e.mock.On("Start", _a0)}
}

func (_c *HeadTracker_Start_Call[H, BLOCK_HASH]) Run(run func(_a0 context.Context)) *HeadTracker_Start_Call[H, BLOCK_HASH] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *HeadTracker_Start_Call[H, BLOCK_HASH]) Return(_a0 error) *HeadTracker_Start_Call[H, BLOCK_HASH] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeadTracker_Start_Call[H, BLOCK_HASH]) RunAndReturn(run func(context.Context) error) *HeadTracker_Start_Call[H, BLOCK_HASH] {
	_c.Call.Return(run)
	return _c
}

// NewHeadTracker creates a new instance of HeadTracker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHeadTracker[H types.Head[BLOCK_HASH], BLOCK_HASH types.Hashable](t interface {
	mock.TestingT
	Cleanup(func())
}) *HeadTracker[H, BLOCK_HASH] {
	mock := &HeadTracker[H, BLOCK_HASH]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
