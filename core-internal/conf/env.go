package conf

type Env struct {
	Lala           string `mapstructure:"LALA"`
	Port           string `mapstructure:"PORT" json:"PORT"`
	Debugging      bool   `mapstructure:"DEBUGGING"`
	ContextTimeOut int    `mapstructure:"CONTEXT_TIMEOUT"`

	GRPCHost string `mapstructure:"GRPC_HOST"`
	GRPCPort string `mapstructure:"GRPC_PORT"`

	GRPCProductHost string `mapstructure:"GRPC_PRODUCT_HOST"`

	OtelHttpExporter string `mapstructure:"OTEL_HTTP_EXPORTER"`
	OtelGrpcExporter string `mapstructure:"OTEL_GRPC_EXPORTER"`

	DBReadHost   string `mapstructure:"DB_READ_HOST"`
	DBReadPort   string `mapstructure:"DB_READ_PORT"`
	DBReadName   string `mapstructure:"DB_READ_NAME"`
	DBReadUser   string `mapstructure:"DB_READ_USER"`
	DBReadPass   string `mapstructure:"DB_READ_PASS"`
	DBReadSchema string `mapstructure:"DB_READ_SCHEMA"`

	DBWriteHost   string `mapstructure:"DB_WRITE_HOST"`
	DBWritePort   string `mapstructure:"DB_WRITE_PORT"`
	DBWriteName   string `mapstructure:"DB_WRITE_NAME"`
	DBWriteUser   string `mapstructure:"DB_WRITE_USER"`
	DBWritePass   string `mapstructure:"DB_WRITE_PASS"`
	DBWriteSchema string `mapstructure:"DB_WRITE_SCHEMA"`

	KafkaHost          string `mapstructure:"KAFKA_HOST"`
	KafkaConsumerGroup string `mapstructure:"KAFKA_CONSUMER_GROUP"`

	RedisHost    string `mapstructure:"REDIS_HOST_PORT"`
	RedisPass    string `mapstructure:"REDIS_PASS"`
	RedisDB      string `mapstructure:"REDIS_DB"`
	RedisTimeOut int64  `mapstructure:"REDIS_TIMEOUT"`

	ElasticHost string `mapstructure:"ELASTIC_HOST"`
	ElasticUser string `mapstructure:"ELASTIC_USER"`
	ElasticPass string `mapstructure:"ELASTIC_PASS"`
}
