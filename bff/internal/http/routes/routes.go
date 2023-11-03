package routes

import (
	"fmt"
	"github.com/yosida95/uritemplate"
	"net/http"
)

var (
	GETApp = Route{
		Template: uritemplate.MustNew("/app"),
		Raw:      "/app/*",
	}
	GETAppNoSlash = Route{
		Template: uritemplate.MustNew("/app"),
		Raw:      "/app",
		Redirect: &Redirect{
			Target:     "/app/",
			StatusCode: http.StatusMovedPermanently,
		},
	}
	GETApi = Route{
		Template: uritemplate.MustNew("/api"),
		Raw:      "/api",
	}
	POSTLogin = Route{
		Template: uritemplate.MustNew("/login"),
		Raw:      "/login",
	}
	GETGqlPlayground = Route{
		Template: uritemplate.MustNew("/gql"),
		Raw:      "/gql",
	}
	POSTGql = Route{
		Template: uritemplate.MustNew("/gql"),
		Raw:      "/gql",
	}
	AnySubscriptions = Route{
		Template: uritemplate.MustNew("/sub"),
		Raw:      "/sub",
	}
)

type Route struct {
	Raw      string
	Template *uritemplate.Template
	Redirect *Redirect
}
type Redirect struct {
	Target     string
	StatusCode int
}

func (r Redirect) Handler() http.Handler {
	return http.RedirectHandler(r.Target, r.StatusCode)
}

func (r Route) MustExpand(vals uritemplate.Values) string {
	s, err := r.Template.Expand(vals)
	if err != nil {
		panic(fmt.Errorf("failed to expand %s: %w", r.Template.Raw(), err))
	}
	return s
}
