package uaa

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/httpclient"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Factory struct {
	logTag string
	logger boshlog.Logger
}

func NewFactory(logger boshlog.Logger) Factory {
	return Factory{
		logTag: "uaa.Factory",
		logger: logger,
	}
}

func (f Factory) New(config Config) (UAA, error) {
	err := config.Validate()
	if err != nil {
		return UAAImpl{}, bosherr.WrapErrorf(
			err, "Validating UAA connection config")
	}

	client, err := f.httpClient(config)
	if err != nil {
		return UAAImpl{}, err
	}

	return UAAImpl{client: client}, nil
}

func (f Factory) httpClient(config Config) (Client, error) {
	certPool, err := config.CACertPool()
	if err != nil {
		return Client{}, err
	}

	if certPool == nil {
		f.logger.Debug(f.logTag, "Using default root CAs")
	} else {
		f.logger.Debug(f.logTag, "Using custom root CAs")
	}

	rawClient := httpclient.CreateDefaultClientInsecureSkipVerify()

	SetTransportMapping(rawClient.Transport.(*http.Transport), config.FloatingIP, config.Host)

	retryClient := httpclient.NewNetworkSafeRetryClient(rawClient, 5, 500*time.Millisecond, f.logger)

	httpClient := httpclient.NewHTTPClient(retryClient, f.logger)

	endpoint := url.URL{
		Scheme: "https",
		Host:   net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port)),
		Path:   config.Path,
	}

	return NewClient(endpoint.String(), config.Client, config.ClientSecret, httpClient, f.logger), nil
}

func SetTransportMapping(transport *http.Transport, fip, hostname string) {
	transport.TLSClientConfig.ServerName = hostname
	transport.DialTLS = func(network, addr string) (conn net.Conn, e error) {
		parts := strings.Split(addr, ":")
		host := fip
		port := "443"
		if len(parts) > 1 {
			port = parts[1]
		}
		return tls.Dial(network, strings.Join([]string{host, port}, ":"), transport.TLSClientConfig)
	}
}
