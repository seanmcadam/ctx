package ctx

import "testing"

func TestCtx_std_create(t *testing.T) {

	var c, c2 *Ctx

	if c = New(); c == nil {
		t.Fatalf("New() returned nil\n")
	}

	if c2 = c.WithCancel(); c2 == nil {
		t.Fatalf("WithCancel() returned nil\n")
	}

	if c.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	if c2.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

}

func TestCtx_create_cancel_child(t *testing.T) {

	var c, c2 *Ctx

	if c = New(); c == nil {
		t.Fatalf("New() returned nil\n")
	}

	if c2 = c.WithCancel(); c2 == nil {
		t.Fatalf("WithCancel() returned nil\n")
	}

	if c.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	if c2.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	c2.Cancel()

	if c.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	if !c2.Done() {
		t.Fatalf("Ctx Done returned false\n")
	}

}

func TestCtx_create_cancel_parent(t *testing.T) {

	var c, c2 *Ctx

	if c = New(); c == nil {
		t.Fatalf("New() returned nil\n")
	}

	if c2 = c.WithCancel(); c2 == nil {
		t.Fatalf("WithCancel() returned nil\n")
	}

	if c.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	if c2.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	c.Cancel()

	if !c.Done() {
		t.Fatalf("Ctx Done returned false\n")
	}

	if !c2.Done() {
		t.Fatalf("Ctx Done returned false\n")
	}

}


func TestCtx_create_get_context(t *testing.T) {

	var c *Ctx

	if c = New(); c == nil {
		t.Fatalf("New() returned nil\n")
	}

	if c.mycontext != c.Context(){
		t.Fatalf("Context() does not match\n")
	}


}

func TestCtx_create_wo_cancel(t *testing.T) {

	var c, c2 *Ctx

	if c = New(); c == nil {
		t.Fatalf("New() returned nil\n")
	}

	if c2 = c.WithoutCancel(); c2 == nil {
		t.Fatalf("WithCancel() returned nil\n")
	}

	if c.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	if c2.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

	c.Cancel()

	if !c.Done() {
		t.Fatalf("Ctx Done returned false\n")
	}

	if c2.Done() {
		t.Fatalf("Ctx Done returned true\n")
	}

}