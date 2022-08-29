package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("New returns nil")
	} else {
		tracer.Trace("Hello Tracer")
		if buf.String() != "Hello Tracer\n" {
			t.Errorf("Error output: '%s'", buf.String())
		}
	}

}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("data")
}
