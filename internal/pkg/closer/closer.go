package closer

import (
	"errors"
	"sync"
)

var globalCloser = New()

func Add(handler handler) {
	globalCloser.Add(handler)
}

func Close() error {
	return globalCloser.Close()
}

type handler func() error
type Closer struct {
	handlers []handler
	mux      sync.Mutex
}

func New() *Closer {
	return &Closer{}
}

func (c *Closer) Add(handler handler) {
	c.mux.Lock()
	c.handlers = append(c.handlers, handler)
	c.mux.Unlock()
}

func (c *Closer) Close() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	var (
		wg        sync.WaitGroup
		allErrors []error
		mux       sync.Mutex
	)

	wg.Add(len(c.handlers))

	for _, h := range c.handlers {
		go func() {
			defer wg.Done()
			if err := h(); err != nil {
				mux.Lock()
				allErrors = append(allErrors, err)
				mux.Unlock()
			}
		}()
	}
	wg.Wait()

	return errors.Join(allErrors...)
}
