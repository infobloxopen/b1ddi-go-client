package ip_space

import (
	"context"
	"crypto/tls"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"github.com/infobloxopen/b1ddi-go-client/runtimetest"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		expectedRequest  http.Request
		testMethodName   string
		testMethodParams interface{}
	}{
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space/bulk_copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"copy_objects\":null}\n")),
			},
			"IPSpaceBulkCopy",
			&IPSpaceBulkCopyParams{Body: &models.IpamsvcBulkCopyIPSpace{}, Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space/ip-space-copy-id/copy"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"IPSpaceCopy",
			&IPSpaceCopyParams{ID: "ip-space-copy-id", Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space"},
				Method: http.MethodPost,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Test Comment\",\"name\":\"ip_space_go_client_test\"}\n")),
			},
			"IPSpaceCreate",
			&IPSpaceCreateParams{Body: &models.IpamsvcIPSpace{
				Name:    swag.String("ip_space_go_client_test"),
				Comment: "Test Comment",
			}, Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space/delete-id"},
				Method: http.MethodDelete,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"IPSpaceDelete",
			&IPSpaceDeleteParams{ID: "delete-id", Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"IPSpaceList",

			&IPSpaceListParams{Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space/ip-space-read-id"},
				Method: http.MethodGet,
				Body:   io.NopCloser(strings.NewReader("")),
			},
			"IPSpaceRead",
			&IPSpaceReadParams{ID: "ip-space-read-id", Context: context.TODO()},
		},
		{
			http.Request{
				URL:    &url.URL{Path: "/api/ddi/v1/ipam/ip_space/ip-space-update-id"},
				Method: http.MethodPatch,
				Body:   io.NopCloser(strings.NewReader("{\"comment\":\"Updated comment\"}\n")),
			},
			"IPSpaceUpdate",
			&IPSpaceUpdateParams{
				ID:      "ip-space-update-id",
				Body:    &models.IpamsvcIPSpace{Comment: "Updated comment"},
				Context: context.TODO(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testMethodName, func(t *testing.T) {
			// Get the mock server response body
			resp := reflect.ValueOf(&Client{}).MethodByName(tc.testMethodName).Type().Out(0)
			// Initialize mock server
			s := runtimetest.InitTestServer(t, tc.expectedRequest, reflect.New(resp))

			// Initialize the client
			c := initIPSpaceTestClient(s.URL)

			methodParams := []reflect.Value{
				reflect.ValueOf(tc.testMethodParams),
				reflect.New(reflect.TypeOf((*runtime.ClientAuthInfoWriter)(nil)).Elem()).Elem(),
			}

			// Call the method by the name specified in the test case
			results := reflect.ValueOf(c).MethodByName(tc.testMethodName).Call(methodParams)
			if err := results[1].Interface(); err != nil {
				t.Fatal(err)
			}

			// Close the mock server
			s.Close()
		})
	}
}

func initIPSpaceTestClient(url string) ClientService {
	transport := httptransport.New(
		strings.TrimPrefix(url, "https://"), "api/ddi/v1", nil,
	)
	// Disable TLS verify for the mock server
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	transport.Transport = tr
	// create the IP Space API client
	return New(transport, strfmt.Default)
}