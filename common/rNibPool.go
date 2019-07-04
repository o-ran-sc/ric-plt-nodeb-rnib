package common

import (
	"sync/atomic"
)

type Pool struct {
	New     func() interface{}
	Destroy func(interface{})
	pool    chan interface{}
	created int32 //Number of objects created
}

/*
NewPool creates thread safe Pool object and returns a pointer to it.
poolSize int - sets the capacity of the pool
newObj func - specifies a function to generate a value (pool element)
destroyObj func - specifies a function to destroy a value (pool element)
*/
func NewPool(poolSize int, newObj func() interface{}, destroyObj func(interface{})) *Pool{
	return &Pool{
		New:     newObj,
		Destroy: destroyObj,
		pool:    make(chan interface{}, poolSize),
	}
}

/*
Retrieve an object from the pool.
If the pool is empty and the number of used object is less than capacity, a new object is created by calling New.
Otherwise, the method blocks until an object is returned to the pool.
*/
func (p *Pool) Get() interface{} {
	select {
	case obj := <-p.pool:
		return obj
	default:
		if atomic.AddInt32(&p.created, 1) <= int32(cap(p.pool)) && p.New != nil {
			p.pool <- p.New()
		}
	}
	return <-p.pool //block waiting
}

/*
Return an object to the pool.
If capacity is exceeded the object is discarded after calling Destroy on it if Destroy is not nil.
*/
func (p *Pool) Put(obj interface{}) {
	if obj != nil {
		select {
		case p.pool <- obj:
		default:
			if p.Destroy != nil {
				p.Destroy(obj)
			}
		}
	}
}

/*
 Closes the pool and if Destroy is not nil, call Destroy on each object in the pool
 The pool must not be used once this method is called.
*/
func (p *Pool) Close() {
	close(p.pool)
	available := len(p.pool)
	if p.Destroy != nil {
		for obj := range p.pool {
			p.Destroy(obj)

		}
	}
	atomic.AddInt32(&p.created, -int32(available))
}
/*
Return statistics.
available - the number of available instances
created - the number of created instances
*/
func (p *Pool) Stats() (available int, created int) {
	return len(p.pool), int(p.created)
}
