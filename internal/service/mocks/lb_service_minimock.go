// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i testappwink/internal/service.LBService -o lb_service_minimock.go -n LBServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// LBServiceMock implements service.LBService
type LBServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetURL          func(ctx context.Context, inputURL string) (s1 string, err error)
	inspectFuncGetURL   func(ctx context.Context, inputURL string)
	afterGetURLCounter  uint64
	beforeGetURLCounter uint64
	GetURLMock          mLBServiceMockGetURL
}

// NewLBServiceMock returns a mock for service.LBService
func NewLBServiceMock(t minimock.Tester) *LBServiceMock {
	m := &LBServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetURLMock = mLBServiceMockGetURL{mock: m}
	m.GetURLMock.callArgs = []*LBServiceMockGetURLParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLBServiceMockGetURL struct {
	mock               *LBServiceMock
	defaultExpectation *LBServiceMockGetURLExpectation
	expectations       []*LBServiceMockGetURLExpectation

	callArgs []*LBServiceMockGetURLParams
	mutex    sync.RWMutex
}

// LBServiceMockGetURLExpectation specifies expectation struct of the LBService.GetURL
type LBServiceMockGetURLExpectation struct {
	mock    *LBServiceMock
	params  *LBServiceMockGetURLParams
	results *LBServiceMockGetURLResults
	Counter uint64
}

// LBServiceMockGetURLParams contains parameters of the LBService.GetURL
type LBServiceMockGetURLParams struct {
	ctx      context.Context
	inputURL string
}

// LBServiceMockGetURLResults contains results of the LBService.GetURL
type LBServiceMockGetURLResults struct {
	s1  string
	err error
}

// Expect sets up expected params for LBService.GetURL
func (mmGetURL *mLBServiceMockGetURL) Expect(ctx context.Context, inputURL string) *mLBServiceMockGetURL {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("LBServiceMock.GetURL mock is already set by Set")
	}

	if mmGetURL.defaultExpectation == nil {
		mmGetURL.defaultExpectation = &LBServiceMockGetURLExpectation{}
	}

	mmGetURL.defaultExpectation.params = &LBServiceMockGetURLParams{ctx, inputURL}
	for _, e := range mmGetURL.expectations {
		if minimock.Equal(e.params, mmGetURL.defaultExpectation.params) {
			mmGetURL.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetURL.defaultExpectation.params)
		}
	}

	return mmGetURL
}

// Inspect accepts an inspector function that has same arguments as the LBService.GetURL
func (mmGetURL *mLBServiceMockGetURL) Inspect(f func(ctx context.Context, inputURL string)) *mLBServiceMockGetURL {
	if mmGetURL.mock.inspectFuncGetURL != nil {
		mmGetURL.mock.t.Fatalf("Inspect function is already set for LBServiceMock.GetURL")
	}

	mmGetURL.mock.inspectFuncGetURL = f

	return mmGetURL
}

// Return sets up results that will be returned by LBService.GetURL
func (mmGetURL *mLBServiceMockGetURL) Return(s1 string, err error) *LBServiceMock {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("LBServiceMock.GetURL mock is already set by Set")
	}

	if mmGetURL.defaultExpectation == nil {
		mmGetURL.defaultExpectation = &LBServiceMockGetURLExpectation{mock: mmGetURL.mock}
	}
	mmGetURL.defaultExpectation.results = &LBServiceMockGetURLResults{s1, err}
	return mmGetURL.mock
}

// Set uses given function f to mock the LBService.GetURL method
func (mmGetURL *mLBServiceMockGetURL) Set(f func(ctx context.Context, inputURL string) (s1 string, err error)) *LBServiceMock {
	if mmGetURL.defaultExpectation != nil {
		mmGetURL.mock.t.Fatalf("Default expectation is already set for the LBService.GetURL method")
	}

	if len(mmGetURL.expectations) > 0 {
		mmGetURL.mock.t.Fatalf("Some expectations are already set for the LBService.GetURL method")
	}

	mmGetURL.mock.funcGetURL = f
	return mmGetURL.mock
}

// When sets expectation for the LBService.GetURL which will trigger the result defined by the following
// Then helper
func (mmGetURL *mLBServiceMockGetURL) When(ctx context.Context, inputURL string) *LBServiceMockGetURLExpectation {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("LBServiceMock.GetURL mock is already set by Set")
	}

	expectation := &LBServiceMockGetURLExpectation{
		mock:   mmGetURL.mock,
		params: &LBServiceMockGetURLParams{ctx, inputURL},
	}
	mmGetURL.expectations = append(mmGetURL.expectations, expectation)
	return expectation
}

// Then sets up LBService.GetURL return parameters for the expectation previously defined by the When method
func (e *LBServiceMockGetURLExpectation) Then(s1 string, err error) *LBServiceMock {
	e.results = &LBServiceMockGetURLResults{s1, err}
	return e.mock
}

// GetURL implements service.LBService
func (mmGetURL *LBServiceMock) GetURL(ctx context.Context, inputURL string) (s1 string, err error) {
	mm_atomic.AddUint64(&mmGetURL.beforeGetURLCounter, 1)
	defer mm_atomic.AddUint64(&mmGetURL.afterGetURLCounter, 1)

	if mmGetURL.inspectFuncGetURL != nil {
		mmGetURL.inspectFuncGetURL(ctx, inputURL)
	}

	mm_params := LBServiceMockGetURLParams{ctx, inputURL}

	// Record call args
	mmGetURL.GetURLMock.mutex.Lock()
	mmGetURL.GetURLMock.callArgs = append(mmGetURL.GetURLMock.callArgs, &mm_params)
	mmGetURL.GetURLMock.mutex.Unlock()

	for _, e := range mmGetURL.GetURLMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmGetURL.GetURLMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetURL.GetURLMock.defaultExpectation.Counter, 1)
		mm_want := mmGetURL.GetURLMock.defaultExpectation.params
		mm_got := LBServiceMockGetURLParams{ctx, inputURL}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetURL.t.Errorf("LBServiceMock.GetURL got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetURL.GetURLMock.defaultExpectation.results
		if mm_results == nil {
			mmGetURL.t.Fatal("No results are set for the LBServiceMock.GetURL")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmGetURL.funcGetURL != nil {
		return mmGetURL.funcGetURL(ctx, inputURL)
	}
	mmGetURL.t.Fatalf("Unexpected call to LBServiceMock.GetURL. %v %v", ctx, inputURL)
	return
}

// GetURLAfterCounter returns a count of finished LBServiceMock.GetURL invocations
func (mmGetURL *LBServiceMock) GetURLAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetURL.afterGetURLCounter)
}

// GetURLBeforeCounter returns a count of LBServiceMock.GetURL invocations
func (mmGetURL *LBServiceMock) GetURLBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetURL.beforeGetURLCounter)
}

// Calls returns a list of arguments used in each call to LBServiceMock.GetURL.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetURL *mLBServiceMockGetURL) Calls() []*LBServiceMockGetURLParams {
	mmGetURL.mutex.RLock()

	argCopy := make([]*LBServiceMockGetURLParams, len(mmGetURL.callArgs))
	copy(argCopy, mmGetURL.callArgs)

	mmGetURL.mutex.RUnlock()

	return argCopy
}

// MinimockGetURLDone returns true if the count of the GetURL invocations corresponds
// the number of defined expectations
func (m *LBServiceMock) MinimockGetURLDone() bool {
	for _, e := range m.GetURLMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetURLMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetURL != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetURLInspect logs each unmet expectation
func (m *LBServiceMock) MinimockGetURLInspect() {
	for _, e := range m.GetURLMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LBServiceMock.GetURL with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetURLMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		if m.GetURLMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to LBServiceMock.GetURL")
		} else {
			m.t.Errorf("Expected call to LBServiceMock.GetURL with params: %#v", *m.GetURLMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetURL != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		m.t.Error("Expected call to LBServiceMock.GetURL")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LBServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetURLInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LBServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *LBServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetURLDone()
}
