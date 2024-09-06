package config

type GRPCConfig interface {
	Address() string
}

type LoadBalancerConfig interface {
	OriginalHost() string
	CDNHost() string
}
