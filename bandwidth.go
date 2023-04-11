package bandwidth

import (
	"net/http"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
	caddyhttp.RegisterHandlerDirective("bandwidth", parseCaddyfile)
}

type Middleware struct {
	Limit int `json:"limit,omitempty"`
}

func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.bandwidth",
		New: func() caddy.Module { return new(Middleware) },
	}
}

func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		args := d.RemainingArgs()
		if len(args) > 0 {
			limit, err := strconv.Atoi(args[0])
			if err != nil {
				return d.Errf("parsing limit value: %v", err)
			}
			m.Limit = limit
		}
	}
	return nil
}

func (m *Middleware) Provision(ctx caddy.Context) error {
	return nil
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if m.Limit > 0 {
		w = &limitedResponseWriter{
			ResponseWriter: w,
			remaining:      int64(m.Limit),
		}
	}
	return next.ServeHTTP(w, r)
}

type limitedResponseWriter struct {
	http.ResponseWriter
	remaining int64
}

func (l *limitedResponseWriter) Write(p []byte) (int, error) {
	if l.remaining <= 0 {
		return len(p), nil
	}

	if int64(len(p)) > l.remaining {
		p = p[:l.remaining]
	}

	n, err := l.ResponseWriter.Write(p)
	l.remaining -= int64(n)

	return n, err
}

func parseCaddyfile(h caddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}
