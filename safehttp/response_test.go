package safehttp

import (
	"testing"
)

type RedirectResponseWriter struct {
	MockHeadersWriter
}

func (w RedirectResponseWriter) Write(resp Response) Result {
	_, ok := resp.(RedirectResponse)
	if !ok {
		panic("Expected RedirectResponse")
	}
	return Result{}
}

func (w RedirectResponseWriter) WriteError(resp ErrorResponse) Result {
	return Result{}
}

type MockHeadersWriter struct {
}

func (w MockHeadersWriter) Header() Header {
	return Header{}
}

func (w MockHeadersWriter) AddCookie(c *Cookie) error {
	return nil
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()

	f()
}

func TestResponseRedirect(t *testing.T) {
	mockResponseWriter := RedirectResponseWriter{}
	incomingReq := &IncomingRequest{}

	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("Expected no panic, but panic occurred: %v", r)
		}
	}()

	Redirect(mockResponseWriter, incomingReq, "http://example.com", StatusFound)

	assertPanic(t, func() {
		Redirect(mockResponseWriter, incomingReq, "http://example.com", StatusNotFound)
	})
}

type JSONResponseWriter struct {
	MockHeadersWriter
}

func (w JSONResponseWriter) Write(resp Response) Result {
	_, ok := resp.(JSONResponse)
	if !ok {
		panic("Expected JSONResponse")
	}
	return Result{}
}

func (w JSONResponseWriter) WriteError(resp ErrorResponse) Result {
	return Result{}
}

func TestWriteJSON(t *testing.T) {
	w := JSONResponseWriter{}

	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("Expected no panic, but panic occurred: %v", r)
		}
	}()

	jsondata := map[string]interface{}{
		"key": "value",
	}
	WriteJSON(w, jsondata)
}
