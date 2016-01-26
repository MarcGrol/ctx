// +build appengine

package ctx

import (
	"net/http"

	"google.golang.org/appengine"

	"golang.org/x/net/context"
)

type appEngineContextFactory struct {
}

func (f appEngineContextFactory) CreateContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

func init() {
	New = appEngineContextFactory{}
}
