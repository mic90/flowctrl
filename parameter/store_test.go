package parameter_test

import (
	"github.com/mic90/flowctrl/buffer"
	"github.com/mic90/flowctrl/parameter"
	"github.com/mic90/flowctrl/parameter/adapter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseStore_StoreRestore(t *testing.T) {
	//GIVEN
	expectedStore := int8(100)
	expectedRestore := int8(0)
	paramA := adapter.NewInt8(expectedRestore, parameter.New(buffer.Int8, "paramA", "description", parameter.UnitNone, false, true))
	paramB := adapter.NewInt8(expectedRestore, parameter.New(buffer.Int8, "paramB", "description", parameter.UnitPercent, false, true))
	container := parameter.NewContainer(paramA, paramB)
	store := parameter.NewBaseStore()
	//WHEN
	storeErr := store.Store(container)
	paramA.Set(expectedStore)
	storeRead := paramA.Get()
	restoreErr := store.Restore(container)
	restoreRead := paramA.Get()
	//THEN
	assert.Nil(t, storeErr)
	assert.Nil(t, restoreErr)
	assert.Equal(t, expectedStore, storeRead)
	assert.Equal(t, expectedRestore, restoreRead)
}
