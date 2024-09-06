package config

type GRPCConfig interface {
	Address() string
}

type LoadBalancerConfig interface {
	CDNHost() string
}

type LoggerConfig interface {
	Stage() string
}
