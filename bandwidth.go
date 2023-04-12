package bandwidth

import (
	"net/http"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"golang.org/x/time/rate"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("bandwidth", parseCaddyfile)
}

type Middleware struct {
	Limit   int           `json:"limit,omitempty"`
	limiter *rate.Limiter
}

func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.bandwidth",
		New: func() caddy.Module { return new(Middleware) },
	}
}

func (m *Middleware) Provision(ctx caddy.Context) error {
	if m.Limit > 0 {
		m.limiter = rate.NewLimiter(rate.Limit(m.Limit), m.Limit)
	}
	return nil
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if m.limiter != nil {
		w = &limitedResponseWriter{
			ResponseWriter: w,
			limiter:        m.limiter,
			r:              r,
		}
	}
	return next.ServeHTTP(w, r)
}

type limitedResponseWriter struct {
	http.ResponseWriter
	limiter *rate.Limiter
	r       *http.Request
}

func (l *limitedResponseWriter) Write(p []byte) (int, error) {
	n := len(p)
	err := l.limiter.WaitN(l.r.Context(), n)
	if err != nil {
		return 0, err
	}

	return l.ResponseWriter.Write(p)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware

	for h.Next() {
		for h.NextBlock(0) {
			switch h.Val() {
			case "limit":
				limitStr := h.RemainingArgs()
				if len(limitStr) != 1 {
					return nil, h.ArgErr()
				}
				var err error
				m.Limit, err = strconv.Atoi(limitStr[0])
				if err != nil {
					return nil, h.Errf("parsing limit value: %v", err)
				}
			default:
				return nil, h.Errf("unrecognized parameter '%s'", h.Val())
			}
		}
	}

	return m, nil
}
