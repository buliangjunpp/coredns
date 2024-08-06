
## Examples

~~~ corefile

.:53 {
    errors
    health {
      lameduck 5s
    }
    ready
    kubernetes cluster.local in-addr.arpa ip6.arpa {
      pods insecure
      fallthrough in-addr.arpa ip6.arpa
      ttl 30
    }
    prometheus :9153
    whitelist {
      example.com 
      allowed.com
    }
    forward . /etc/resolv.conf {
      prefer_udp
      max_concurrent 1000
    }
    cache 30
    loop
    reload
    loadbalance
}

~~~
