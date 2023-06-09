// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gridironOne/gridiron/eth/core"
	"sync"
)

// Ensure, that ConfigurationPluginMock does implement core.ConfigurationPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.ConfigurationPlugin = &ConfigurationPluginMock{}

// ConfigurationPluginMock is a mock implementation of core.ConfigurationPlugin.
//
//	func TestSomethingThatUsesConfigurationPlugin(t *testing.T) {
//
//		// make and configure a mocked core.ConfigurationPlugin
//		mockedConfigurationPlugin := &ConfigurationPluginMock{
//			ChainConfigFunc: func() *params.ChainConfig {
//				panic("mock out the ChainConfig method")
//			},
//			ExtraEipsFunc: func() []int {
//				panic("mock out the ExtraEips method")
//			},
//			FeeCollectorFunc: func() *common.Address {
//				panic("mock out the FeeCollector method")
//			},
//			PrepareFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Prepare method")
//			},
//		}
//
//		// use mockedConfigurationPlugin in code that requires core.ConfigurationPlugin
//		// and then make assertions.
//
//	}
type ConfigurationPluginMock struct {
	// ChainConfigFunc mocks the ChainConfig method.
	ChainConfigFunc func() *params.ChainConfig

	// ExtraEipsFunc mocks the ExtraEips method.
	ExtraEipsFunc func() []int

	// FeeCollectorFunc mocks the FeeCollector method.
	FeeCollectorFunc func() *common.Address

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(contextMoqParam context.Context)

	// calls tracks calls to the methods.
	calls struct {
		// ChainConfig holds details about calls to the ChainConfig method.
		ChainConfig []struct {
		}
		// ExtraEips holds details about calls to the ExtraEips method.
		ExtraEips []struct {
		}
		// FeeCollector holds details about calls to the FeeCollector method.
		FeeCollector []struct {
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockChainConfig  sync.RWMutex
	lockExtraEips    sync.RWMutex
	lockFeeCollector sync.RWMutex
	lockPrepare      sync.RWMutex
}

// ChainConfig calls ChainConfigFunc.
func (mock *ConfigurationPluginMock) ChainConfig() *params.ChainConfig {
	if mock.ChainConfigFunc == nil {
		panic("ConfigurationPluginMock.ChainConfigFunc: method is nil but ConfigurationPlugin.ChainConfig was just called")
	}
	callInfo := struct {
	}{}
	mock.lockChainConfig.Lock()
	mock.calls.ChainConfig = append(mock.calls.ChainConfig, callInfo)
	mock.lockChainConfig.Unlock()
	return mock.ChainConfigFunc()
}

// ChainConfigCalls gets all the calls that were made to ChainConfig.
// Check the length with:
//
//	len(mockedConfigurationPlugin.ChainConfigCalls())
func (mock *ConfigurationPluginMock) ChainConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockChainConfig.RLock()
	calls = mock.calls.ChainConfig
	mock.lockChainConfig.RUnlock()
	return calls
}

// ExtraEips calls ExtraEipsFunc.
func (mock *ConfigurationPluginMock) ExtraEips() []int {
	if mock.ExtraEipsFunc == nil {
		panic("ConfigurationPluginMock.ExtraEipsFunc: method is nil but ConfigurationPlugin.ExtraEips was just called")
	}
	callInfo := struct {
	}{}
	mock.lockExtraEips.Lock()
	mock.calls.ExtraEips = append(mock.calls.ExtraEips, callInfo)
	mock.lockExtraEips.Unlock()
	return mock.ExtraEipsFunc()
}

// ExtraEipsCalls gets all the calls that were made to ExtraEips.
// Check the length with:
//
//	len(mockedConfigurationPlugin.ExtraEipsCalls())
func (mock *ConfigurationPluginMock) ExtraEipsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockExtraEips.RLock()
	calls = mock.calls.ExtraEips
	mock.lockExtraEips.RUnlock()
	return calls
}

// FeeCollector calls FeeCollectorFunc.
func (mock *ConfigurationPluginMock) FeeCollector() *common.Address {
	if mock.FeeCollectorFunc == nil {
		panic("ConfigurationPluginMock.FeeCollectorFunc: method is nil but ConfigurationPlugin.FeeCollector was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFeeCollector.Lock()
	mock.calls.FeeCollector = append(mock.calls.FeeCollector, callInfo)
	mock.lockFeeCollector.Unlock()
	return mock.FeeCollectorFunc()
}

// FeeCollectorCalls gets all the calls that were made to FeeCollector.
// Check the length with:
//
//	len(mockedConfigurationPlugin.FeeCollectorCalls())
func (mock *ConfigurationPluginMock) FeeCollectorCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFeeCollector.RLock()
	calls = mock.calls.FeeCollector
	mock.lockFeeCollector.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *ConfigurationPluginMock) Prepare(contextMoqParam context.Context) {
	if mock.PrepareFunc == nil {
		panic("ConfigurationPluginMock.PrepareFunc: method is nil but ConfigurationPlugin.Prepare was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(contextMoqParam)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedConfigurationPlugin.PrepareCalls())
func (mock *ConfigurationPluginMock) PrepareCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}
