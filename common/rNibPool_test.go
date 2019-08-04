//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var	max int32
var	counter int32
var poolGlob *Pool
type instance struct{}

func(instance) up(){
	tmpc := atomic.AddInt32(&counter, 1)
	swapped:= false
	for !swapped {
		tmpm := atomic.LoadInt32(&max)
		if tmpc >tmpm {
			swapped = atomic.CompareAndSwapInt32(&max, tmpm, tmpc)
		} else {
			break
		}
	}
}

func(instance) down(){
	atomic.AddInt32(&counter, - 1)
}

func TestPoolMax(t *testing.T){
	counter = 0
	max = 0
	validateMaxLimit(1, 1, t)
	counter = 0
	max = 0
	validateMaxLimit(1, 2, t)
	counter = 0
	max = 0
	validateMaxLimit(5, 10, t)
}

func validateMaxLimit(size int, iterations int, t *testing.T) {
	poolGlob = NewPool(size, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	group := sync.WaitGroup{}
	for i := 0; i < iterations; i++ {
		go func() {
			group.Add(1)
			getPutInstance()
			group.Done()
		}()
	}
	time.Sleep(time.Second)
	group.Wait()
	assert.Equal(t, int32(size), max)
}

func getPutInstance() {
	inst := poolGlob.Get().(instance)
	inst.up()
	time.Sleep(time.Millisecond*10)
	inst.down()
	poolGlob.Put(inst)
}

func TestNewPool(t *testing.T){
	size := 5
	pool := NewPool(size, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	assert.NotNil(t, pool)
	assert.NotNil(t, pool.New)
	assert.NotNil(t, pool.Destroy)
	assert.NotNil(t, pool.pool)
	assert.Equal(t, cap(pool.pool), size, "the capacity of the pool should be " + string(size))
}

func TestGetCreated(t *testing.T) {
	pool := NewPool(1, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	pool.Get()
	available, created := pool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 1, created, "number of created objects in the pool should be 1")
	pool.Close()
}

func TestGetAndPut(t *testing.T) {
	pool := NewPool(1, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	pool.Put(pool.Get())
	available, created := pool.Stats()
	assert.Equal(t, 1, available, "number of available objects in the pool should be 1")
	assert.Equal(t, 1, created, "number of created objects in the pool should be 1")
	pool.Close()
}

func TestPutOutOfCapacity(t *testing.T) {
	pool := NewPool(1, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	pool.Put(pool.Get())
	pool.Put(new(instance))
	available, created := pool.Stats()
	assert.Equal(t, 1, available, "number of available objects in the pool should be 1")
	assert.Equal(t, 1, created, "number of created objects in the pool should be 1")
	pool.Close()
}

func TestNotInitializedPut(t *testing.T) {
	var poolEmpty Pool
	poolEmpty.Put(new(instance))
	available, created := poolEmpty.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the pool should be 0")
}

func TestPutNilObject(t *testing.T) {
	var poolEmpty Pool
	poolEmpty.Put(nil)
	available, created := poolEmpty.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the pool should be 0")
}

func TestGet(t *testing.T) {
	pool := NewPool(2, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	i1 := pool.Get()
	i2 := pool.Get()
	available, created := pool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 2, created, "number of created objects in the pool should be 2")
	pool.Put(i1)
	pool.Put(i2)
	pool.Put(new(instance))
	available, created = pool.Stats()
	assert.Equal(t, 2, available, "number of available objects in the pool should be 2")
	assert.Equal(t, 2, created, "number of created objects in the pool should be 2")
}

func TestClose(t *testing.T) {
	pool := NewPool(3, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	i1 := pool.Get()
	i2 := pool.Get()
	i3 := pool.Get()
	available, created := pool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 3, created, "number of created objects in the pool should be 3")
	pool.Put(i1)
	pool.Put(i2)
	pool.Put(i3)
	available, created = pool.Stats()
	assert.Equal(t, 3, available, "number of available objects in the pool should be 3")
	assert.Equal(t, 3, created, "number of created objects in the pool should be 3")
	pool.Close()
	i := pool.Get()
	assert.Nil(t, i)
	available, created = pool.Stats()
	assert.Equal(t, 0, available, "number of available objects in the pool should be 0")
	assert.Equal(t, 0, created, "number of created objects in the pool should be 0")
}

func TestPoolPutPanicsOnClosedChannel(t *testing.T){
	pool := NewPool(1, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	close(pool.pool)
	assert.Panics(t, func(){pool.Put(instance{})})
}

func TestPoolClosePanicsOnClosedChannel(t *testing.T) {
	pool := NewPool(1, func() interface{} {
		inst := instance{}
		return inst
	},
		func(obj interface{}) {
		},
	)
	close(pool.pool)
	assert.Panics(t, func(){pool.Close()})
}
