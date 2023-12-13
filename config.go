package main

type Handler func(any) any

type KClientConfig struct {
	Kafka       *KafkaConfig
	Handler     *Handler
	ConsumerNum int
	CacheSize   int
}

type KafkaConfig struct {
	Brokers  []string
	Group    string
	Version  string
	Topics   []string
	Assignor string
	Oldest   bool
	Verbose  bool
}
