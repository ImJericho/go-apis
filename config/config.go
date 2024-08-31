package config

const (
	MONGO_API_KEY  = "e7863c2d-b9df-4d50-94bb-9fb5a09c700a"
	MONGO_URL      = "mongodb+srv://vivekpatidar00w:zRdspnY6JXoINhsV@cluster0.9fjpr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	MONGO_URI      = "mongodb+srv://vivekpatidar00w:zRdspnY6JXoINhsV@cluster0.9fjpr.mongodb.net/"
	MONGO_PASSWORD = "zRdspnY6JXoINhsV"
)

type Config struct {
	MongoApiKey string
	MongoURI    string
	MongoPasswd string
}

func NewConfig() *Config {
	return &Config{
		MongoApiKey: MONGO_API_KEY,
		MongoURI:    MONGO_URI,
		MongoPasswd: MONGO_PASSWORD,
	}
}
