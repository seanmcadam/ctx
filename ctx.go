// Package provide a simple contex management module
package ctx

import "context"

// Ctx is a struct hosding the
type Ctx struct {
	mycontext context.Context
	cancel    func()
}

// New() returns a fresh context that is used at the top of the
// context hierarchy
func New() (c *Ctx) {
	ctx, cancel := context.WithCancel(context.Background())
	c = &Ctx{
		mycontext: ctx,
		cancel:    cancel,
	}
	return c
}

// NewWithCancel() extends an existing Ctx object
// The new object will be canceled with the parent, and
// can be canceled independantly of the parent
func (c *Ctx) WithCancel() (d *Ctx) {
	ctx, cancel := context.WithCancel(c.mycontext)
	d = &Ctx{
		mycontext: ctx,
		cancel:    cancel,
	}
	return d
}

// NewWithCancel() extends an existing Ctx object
// The new object will not be canceled with the parent
func (c *Ctx) WithoutCancel() (d *Ctx) {
	ctx := context.WithoutCancel(c.mycontext)
	d = &Ctx{
		mycontext: ctx,
		cancel:    nil,
	}
	return d
}

// Context() returns the context of the Ctx struct
func (c *Ctx) Context() (ctx context.Context) {
	return c.mycontext
}

// Cancel() send the signal to stop the object that Ctx is embedded in
func (c *Ctx) Cancel() {
	if c.cancel != nil {
		c.cancel()
	}
}

// DoneChan() returns the Done channel
func (c *Ctx) DoneChan() <-chan struct{} {
	return c.mycontext.Done()
}

// Done() returns a boolean for the done status
func (c *Ctx) Done() bool {
	select {
	case <-c.DoneChan():
		return true
	default:
		return false
	}
}
