package kafka

type ProducerConfig struct {
	// BootStrapServers should be a comma seperated string of kafka brokers
	BootstrapServers string `mapstructure:"BOOTSTRAP_SERVERS"`

	// Uniquely identify a kafka client
	ClientID string `mapstructure:"CLIENT_ID"`

	// no. of acknowledgements client will receive before considering successful message delivery
	// "0" => not wait for acks
	// "1" => get ack from leader
	// "2" => get ack from all replicas
	Ack string `mapstructure:"ACK"`

	// Topic to produce messages to
	Topic string `mapstructure:"TOPIC"`
}

type ConsumerConfig struct {
	// BootStrapServers should be a comma seperated string of kafka brokers
	BootstrapServers string `mapstructure:"BOOTSTRAP_SERVERS"`

	// Uniquely identify a kafka consumer
	// multiple consumers can have same ConsumerGroupID
	// offset is unique for Partition X ConsumerGroupID
	ConsumerGroupID string `mapstructure:"CONSUMER_GROUP_ID"`

	// Topic to consumer message from
	Topic string `mapstructure:"TOPIC"`

	// Starting point in each partition for data consumption to start from
	AutoOffsetReset string `mapstructure:"AUTO_OFFSET_RESET"`

	// Polling timeout
	PollTimeout int `mapstructure:"POLL_TIMEOUT"`
}
