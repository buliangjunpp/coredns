package whitelist

import (
	"context"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

type Whitelist struct {
	Next        plugin.Handler
	Whitelisted map[string]struct{}
}

func (wl Whitelist) ServeDNS(ctx context.Context, rw dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: rw, Req: r}

	qname := state.QName()

	// Check if the domain is in the whitelist
	if _, ok := wl.Whitelisted[qname]; ok {
		// If in whitelist, proceed to the next plugin
		return plugin.NextOrFailure(wl.Name(), wl.Next, ctx, rw, r)
	}

	// If not in whitelist, return REFUSED response
	return dns.RcodeRefused, nil
}

func (wl Whitelist) Name() string { return pluginName }
