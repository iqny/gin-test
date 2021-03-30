package conf

import (
	"flag"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	App *AppConfig
}

type AppConfig struct {
	Timezone string
	Host     string
	Gzip     bool
	Debug    bool
}

var (
	confPath string
	once     sync.Once
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default")
}
func Init() *Config {
	return load()
}

func load() (cfg *Config) {
	once.Do(func() {
		path, err := filepath.Abs(confPath)
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(path, &cfg); err != nil {
			panic(err)
		}
	})
	return
}
