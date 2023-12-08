package app

import (
	"github.com/go-chi/chi/v5"
	temporalClient "github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"github.com/temporalio/temporal-demo-infra/bff/internal/http/routes"
	"go.temporal.io/sdk/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type pageCase struct {
	desc             string
	path             string
	expectBody       string
	expectStatusCode int
}

func TestGET(t *testing.T) {

	cases := []pageCase{
		{
			desc:             "not found",
			path:             "/doesntexist",
			expectStatusCode: http.StatusNotFound,
		},
		{
			desc:             "root index served without trailing",
			path:             "/app",
			expectStatusCode: http.StatusMovedPermanently,
		},
		{
			desc:             "root index served by dir",
			path:             "/app/",
			expectBody:       "mynameisindex.html",
			expectStatusCode: http.StatusOK,
		},
		{
			desc:             "/app/login",
			path:             "/app/login",
			expectBody:       "mynameislogin.html",
			expectStatusCode: http.StatusOK,
		},
		{
			desc:             "/app/deep/route (subroutes)",
			path:             "/app/deep/route",
			expectBody:       "mynameisdeeproute.html",
			expectStatusCode: http.StatusOK,
		},
		{
			desc:             "/app/deep (route with subroutes)",
			path:             "/app/deep",
			expectBody:       "mynameisdeep.html",
			expectStatusCode: http.StatusOK,
		},
	}
	tc := new(mocks.Client)

	router := chi.NewRouter()
	handlers, err := NewHandlers(
		WithMountPath(routes.GETApp.Raw),
		WithGeneratedAppDirectory("/testfiles"),
		WithTemporalClients(&temporalClient.Clients{Client: tc}),
	)
	if err != nil {
		t.Fatal(err)
	}
	handlers.Register(router)
	testserver := httptest.NewServer(router)
	defer testserver.Close()
	u, err := url.Parse(testserver.URL)
	if err != nil {
		t.Fatal("unable to server url")
	}
	//vals := uritemplate.Values{}
	//// add values if needed
	//path, err := routes.GETApp.Template.Expand(vals)
	//if err != nil {
	//	t.Fatal("unable to expand uri template")
	//}

	for _, testCase := range cases {
		t.Run(testCase.desc, func(t *testing.T) {
			httpClient := &http.Client{
				Transport: nil,
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
				Jar:     nil,
				Timeout: 0,
			}
			p, err := url.Parse(testCase.path)
			if err != nil {
				t.Fatal("unable to parse path")
			}
			u = u.ResolveReference(p)
			resp, err := httpClient.Get(u.String())
			if err != nil {
				t.Fatalf("failed to GET: %v", err)
			}
			if resp.StatusCode != testCase.expectStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					resp.StatusCode, testCase.expectStatusCode)
			}
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read body %v", err)
			}
			if testCase.expectBody != "" {
				if string(bytes) != testCase.expectBody {
					t.Errorf("handler returned wrong body: got %v want %v", string(bytes), testCase.expectBody)
				}
			}
		})

	}

}
