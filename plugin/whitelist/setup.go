package whitelist

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

const pluginName = "whitelist"

var log = clog.NewWithPlugin(pluginName)

func init() { plugin.Register(pluginName, setup) }

func setup(c *caddy.Controller) error {
	whitelist := make(map[string]struct{})

	for c.Next() {
		for c.NextBlock() {
			domain := c.Val()
			whitelist[domain] = struct{}{}
		}
	}

	w := Whitelist{Whitelisted: whitelist}
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		w.Next = next
		return w
	})

	log.Infof("whitelist: configured %d whitelisted domains", len(whitelist))
	return nil
}
