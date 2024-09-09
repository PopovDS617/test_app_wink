package lb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCdnHostURL = "cdn.com"
var testURLNotCorrect = "http://test/video/hello"
var testURLCorrect = "http://s1.test.com/video/hello"

func TestLBService_Fail(t *testing.T) {

	ctx := context.Background()
	srvc := NewService(testCdnHostURL)

	_, err := srvc.GetURL(ctx, testURLNotCorrect)

	if err == nil {
		t.Errorf("expected err not to be nil, but received: %v", err)
	}
}

func TestLBService_Success(t *testing.T) {

	ctx := context.Background()
	srvc := NewService(testCdnHostURL)

	for i := 1; i < 55; i++ {
		receivedURL, err := srvc.GetURL(ctx, testURLCorrect)

		if err != nil {
			t.Errorf("expected err to be nil, but received: %v", err)
		}

		if i%10 == 0 {
			if !assert.Equal(t, testURLCorrect, receivedURL) {
				t.Errorf("expected to receive %s, but received %s", testURLCorrect, receivedURL)
			}
			continue
		}

		expectedURL := "http://cdn.com/s1/video/hello"

		if !assert.Equal(t, expectedURL, receivedURL) {
			t.Errorf("expected to receive %s, but received %s", expectedURL, receivedURL)
		}
	}
}
