package server

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/acme/autocert"

	"github.com/photoprism/photoprism/internal/config"
)

// AutoTLS enables automatic HTTPS via Let's Encrypt.
func AutoTLS(conf *config.Config) (*autocert.Manager, error) {
	var siteDomain, tlsEmail, certDir string

	// Enable automatic HTTPS via Let's Encrypt?
	if !conf.SiteHttps() {
		return nil, fmt.Errorf("tls disabled")
	} else if siteDomain = conf.SiteDomain(); !strings.Contains(siteDomain, ".") {
		return nil, fmt.Errorf("no fully qualified site domain")
	} else if tlsEmail = conf.TLSEmail(); tlsEmail == "" {
		return nil, fmt.Errorf("automatic tls disabled")
	} else if certDir = conf.CertificatesPath(); certDir == "" {
		return nil, fmt.Errorf("certificates path not found")
	}

	// Create Let's Encrypt cert manager.
	m := &autocert.Manager{
		Email:      tlsEmail,
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(siteDomain),
		Cache:      autocert.DirCache(certDir),
	}

	return m, nil
}
