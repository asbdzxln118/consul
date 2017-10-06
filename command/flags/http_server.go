package flags

import (
	"flag"

	"github.com/hashicorp/consul/configutil"
)

type HTTPServerFlags struct {
	Flags *flag.FlagSet

	// flags
	datacenter configutil.StringValue
	stale      configutil.BoolValue
}

// httpFlagsServer is the list of.Flags that apply to HTTP connections.
func NewServerFlags() *HTTPServerFlags {
	f := &HTTPServerFlags{}
	f.Flags = flag.NewFlagSet("", flag.ContinueOnError)
	f.Flags.Var(&f.datacenter, "datacenter",
		"Name of the datacenter to query. If unspecified, this will default to "+
			"the datacenter of the queried agent.")
	f.Flags.Var(&f.stale, "stale",
		"Permit any Consul server (non-leader) to respond to this request. This "+
			"allows for lower latency and higher throughput, but can result in "+
			"stale data. This option has no effect on non-read operations. The "+
			"default value is false.")
	return f
}

func (f *HTTPServerFlags) AddTo(fs *flag.FlagSet) {
	merge(fs, f.Flags)
}

func merge(dst, src *flag.FlagSet) {
	src.VisitAll(func(f *flag.Flag) {
		dst.Var(f.Value, f.Name, f.DefValue)
	})
}
