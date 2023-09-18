package realip

import (
	"fmt"
	"net/http"
	"testing"
)

type testIP struct {
	name     string
	request  *http.Request
	expected string
}

func TestFromRequest(t *testing.T) {
	newRequest := func(remoteAddr string, headers map[string]string) *http.Request {
		r := &http.Request{}
		r.RemoteAddr = remoteAddr
		r.Header = http.Header{}

		for header, value := range headers {
			r.Header.Add(header, value)
		}

		return r
	}

	testData := []testIP{
		{
			name:     "No header",
			request:  newRequest("144.12.54.87", map[string]string{}),
			expected: "144.12.54.87",
		},
		{
			name:     "Has X-Forwarded-For",
			request:  newRequest("", map[string]string{"X-Forwarded-For": "144.12.54.87"}),
			expected: "144.12.54.87",
		},
		{
			name: "Has multiple X-Forwarded-For",
			request: newRequest("", map[string]string{
				"X-Forwarded-For": fmt.Sprintf("%s,%s,%s", "119.14.55.11", "144.12.54.87", "127.0.0.0"),
			}),
			expected: "119.14.55.11",
		},
		{
			name:     "Has X-Real-IP",
			request:  newRequest("", map[string]string{"X-Real-IP": "144.12.54.87"}),
			expected: "144.12.54.87",
		},
		{
			name: "Has multiple X-Forwarded-For and X-Real-IP",
			request: newRequest("", map[string]string{
				"X-Real-IP":       "119.14.55.11",
				"X-Forwarded-For": fmt.Sprintf("%s,%s", "144.12.54.87", "127.0.0.0"),
			}),
			expected: "144.12.54.87",
		},
	}

	// Run test
	for _, v := range testData {
		if actual := FromRequest(v.request); v.expected != actual {
			t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
		}
	}

}
