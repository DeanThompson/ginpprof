package ginpprof

import (
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func newServer() *gin.Engine {
	r := gin.Default()
	return r
}

func TestWrap(t *testing.T) {
	r := newServer()

	Wrap(r)

	expectedRouters := map[string]string{
		"/debug/pprof/":             "IndexHandler",
		"/debug/pprof/heap":         "HeapHandler",
		"/debug/pprof/goroutine":    "GoroutineHandler",
		"/debug/pprof/block":        "BlockHandler",
		"/debug/pprof/threadcreate": "ThreadCreateHandler",
		"/debug/pprof/cmdline":      "CmdlineHandler",
		"/debug/pprof/profile":      "ProfileHandler",
		"/debug/pprof/symbol":       "SymbolHandler",
		"/debug/pprof/trace":        "TraceHandler",
	}

	routers := r.Routes()
	for _, router := range routers {
		//fmt.Println(router.Path, router.Method, router.Handler)
		name, ok := expectedRouters[router.Path]
		if !ok {
			t.Errorf("missing router %s", router.Path)
		}
		if !strings.Contains(router.Handler, name) {
			t.Errorf("handler for %s should contain %s, got %s", router.Path, name, router.Handler)
		}
	}
}
