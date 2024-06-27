package safehttp_test

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-safeweb/safehttp"
	"github.com/google/safehtml"
)

type panickingInterceptor struct {
	before, commit, onError bool
}

func (p panickingInterceptor) Before(w safehttp.ResponseWriter, _ *safehttp.IncomingRequest, cfg safehttp.InterceptorConfig) safehttp.Result {
	if p.before {
		panic("before")
	}
	return safehttp.NotWritten()
}

func (p panickingInterceptor) Commit(w safehttp.ResponseHeadersWriter, r *safehttp.IncomingRequest, resp safehttp.Response, cfg safehttp.InterceptorConfig) {
	if p.commit {
		panic("commit")
	}
}

func (panickingInterceptor) Match(safehttp.InterceptorConfig) bool {
	return false
}

func TestFlightInterceptorPanic(t *testing.T) {
	tests := []struct {
		desc        string
		interceptor panickingInterceptor
		wantPanic   bool
	}{
		{
			desc:        "panic in Before",
			interceptor: panickingInterceptor{before: true},
			wantPanic:   true,
		},
		{
			desc:        "panic in Commit",
			interceptor: panickingInterceptor{commit: true},
			wantPanic:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			mb := safehttp.NewServeMuxConfig(nil)
			mb.Intercept(tc.interceptor)
			mux := mb.Mux()

			mux.Handle("/search", safehttp.MethodGet, safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {

				w.Header().Set("foo", "bar")
				return w.Write(safehtml.HTMLEscaped("<h1>Hello World!</h1>"))
			}))

			req := httptest.NewRequest(safehttp.MethodGet, "http://foo.com/search", nil)
			rw := httptest.NewRecorder()

			defer func() {
				r := recover()
				if !tc.wantPanic {
					if r != nil {
						t.Fatalf("unexpected panic %v", r)
					}
					return
				}
				if r == nil {
					t.Fatal("expected panic")
				}
				// Good, the panic got propagated.
				if len(rw.Header()) > 0 {
					t.Errorf("ResponseWriter.Header() got %v, want empty", rw.Header())
				}
			}()
			mux.ServeHTTP(rw, req)
		})
	}
}

func TestFlightHandlerPanic(t *testing.T) {
	mb := safehttp.NewServeMuxConfig(nil)
	mux := mb.Mux()

	mux.Handle("/search", safehttp.MethodGet, safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {

		w.Header().Set("foo", "bar")
		panic("handler")
	}))

	req := httptest.NewRequest(safehttp.MethodGet, "http://foo.com/search", nil)
	rw := httptest.NewRecorder()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic")
		}
		// Good, the panic got propagated.
		if len(rw.Header()) > 0 {
			t.Errorf("ResponseWriter.Header() got %v, want empty", rw.Header())
		}
	}()
	mux.ServeHTTP(rw, req)
}

func TestFlightDoubleWritePanics(t *testing.T) {
	writeFuncs := map[string]func(safehttp.ResponseWriter, *safehttp.IncomingRequest) safehttp.Result{
		"Write": func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
			return w.Write(safehtml.HTMLEscaped("Hello"))
		},
		"WriteError": func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
			return w.WriteError(safehttp.StatusPreconditionFailed)
		},
	}

	for firstWriteName, firstWrite := range writeFuncs {
		for secondWriteName, secondWrite := range writeFuncs {
			t.Run(fmt.Sprintf("%s->%s", firstWriteName, secondWriteName), func(t *testing.T) {
				mb := safehttp.NewServeMuxConfig(nil)
				mux := mb.Mux()
				mux.Handle("/search", safehttp.MethodGet, safehttp.HandlerFunc(func(w safehttp.ResponseWriter, r *safehttp.IncomingRequest) safehttp.Result {
					firstWrite(w, r)
					secondWrite(w, r) // this should panic
					t.Fatal("should never reach this point")
					return safehttp.Result{}
				}))

				req := httptest.NewRequest(safehttp.MethodGet, "http://foo.com/search", nil)
				rw := httptest.NewRecorder()
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("expected panic")
					}
					// Good, the panic got propagated.
					// Note: we are not testing the response headers here, as the first write might have already succeeded.
				}()
				mux.ServeHTTP(rw, req)
			})

		}
	}

}

type MockFlightContext struct {
}

func (m MockFlightContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (m MockFlightContext) Done() <-chan struct{} {
	return nil
}

func (m MockFlightContext) Err() error {
	return nil
}

func (m MockFlightContext) Value(key interface{}) interface{} {
	return nil
}

func TestFlightValueNil(t *testing.T) {
	safehttp.FlightValues(MockFlightContext{})
}

func TestCoverageFlightWrite(t *testing.T) {
	safehttp.InitializeCoverageMap()
	TestFlightInterceptorPanic(t)
	TestFlightHandlerPanic(t)
	TestFlightDoubleWritePanics(t)
	safehttp.PrintCoverage()
}
