//go:build with_utls

package v2rayxhttp

import "github.com/sagernet/sing-box/common/tls"

func isRealityClient(config tls.Config) bool {
	switch config := config.(type) {
	case *tls.RealityClientConfig:
		return true
	default:
		return false
	}
}
