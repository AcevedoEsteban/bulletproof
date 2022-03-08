package routes

import(
	"net/http"
"time"
"github.com/go-chi/cors"
)
type Config struct {
	timeout: time.Duration
}
func NewConfig() *Config {
	return &Config{}
}

//gives no ever if someone wants to acces it from another ip address
func (c *config)Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"}
		AllowCredentials: true,
		MaxAge: 5,

	}).Handler(next)

}

func (c *Config) SetTimeout(timeInSecond int) *Config{
	c.timeout = time.Duration(timeInSecond) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}