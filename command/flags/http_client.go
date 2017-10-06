package flags

import (
	"flag"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/configutil"
)

type HTTPClientFlags struct {
	Flags *flag.FlagSet

	// flags
	httpAddr      configutil.StringValue
	token         configutil.StringValue
	caFile        configutil.StringValue
	caPath        configutil.StringValue
	certFile      configutil.StringValue
	keyFile       configutil.StringValue
	tlsServerName configutil.StringValue
}

func NewHTTPClientFlags() *HTTPClientFlags {
	f := &HTTPClientFlags{}
	f.Flags = flag.NewFlagSet("", flag.ContinueOnError)
	f.Flags.Var(&f.caFile, "ca-file",
		"Path to a CA file to use for TLS when communicating with Consul. This "+
			"can also be specified via the CONSUL_CACERT environment variable.")
	f.Flags.Var(&f.caPath, "ca-path",
		"Path to a directory of CA certificates to use for TLS when communicating "+
			"with Consul. This can also be specified via the CONSUL_CAPATH environment variable.")
	f.Flags.Var(&f.certFile, "client-cert",
		"Path to a client cert file to use for TLS when 'verify_incoming' is enabled. This "+
			"can also be specified via the CONSUL_CLIENT_CERT environment variable.")
	f.Flags.Var(&f.keyFile, "client-key",
		"Path to a client key file to use for TLS when 'verify_incoming' is enabled. This "+
			"can also be specified via the CONSUL_CLIENT_KEY environment variable.")
	f.Flags.Var(&f.httpAddr, "http-addr",
		"The `address` and port of the Consul HTTP agent. The value can be an IP "+
			"address or DNS address, but it must also include the port. This can "+
			"also be specified via the CONSUL_HTTP_ADDR environment variable. The "+
			"default value is http://127.0.0.1:8500. The scheme can also be set to "+
			"HTTPS by setting the environment variable CONSUL_HTTP_SSL=true.")
	f.Flags.Var(&f.token, "token",
		"ACL token to use in the request. This can also be specified via the "+
			"CONSUL_HTTP_TOKEN environment variable. If unspecified, the query will "+
			"default to the token of the Consul agent at the HTTP address.")
	f.Flags.Var(&f.tlsServerName, "tls-server-name",
		"The server name to use as the SNI host when connecting via TLS. This "+
			"can also be specified via the CONSUL_TLS_SERVER_NAME environment variable.")
	return f
}

func (f *HTTPClientFlags) AddTo(fs *flag.FlagSet) {
	merge(fs, f.Flags)
}

func (f *HTTPClientFlags) APIClient() (*api.Client, error) {
	config := api.DefaultConfig()

	f.httpAddr.Merge(&config.Address)
	f.token.Merge(&config.Token)
	f.caFile.Merge(&config.TLSConfig.CAFile)
	f.caPath.Merge(&config.TLSConfig.CAPath)
	f.certFile.Merge(&config.TLSConfig.CertFile)
	f.keyFile.Merge(&config.TLSConfig.KeyFile)
	f.tlsServerName.Merge(&config.TLSConfig.Address)

	return api.NewClient(config)
}
