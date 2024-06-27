package safehttp

import "testing"

type MockInterceptor struct {
}

func (i *MockInterceptor) Before(w ResponseWriter, r *IncomingRequest, cfg InterceptorConfig) Result {
	return Result{}
}

func (i *MockInterceptor) Commit(w ResponseHeadersWriter, r *IncomingRequest, resp Response, cfg InterceptorConfig) {
}

func (i *MockInterceptor) Match(InterceptorConfig) bool {
	return true
}

type MockInterceptorConfig struct {
}

func TestInterceptorsMultipleMatches(t *testing.T) {
	interceptors := []Interceptor{
		&MockInterceptor{},
	}

	cfgs := []InterceptorConfig{
		&MockInterceptorConfig{},
		&MockInterceptorConfig{},
		&MockInterceptorConfig{},
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic")
		}
	}()

	configureInterceptors(interceptors, cfgs)
}
