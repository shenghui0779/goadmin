package yiigo

import (
	"github.com/pelletier/go-toml"
	"github.com/philchia/agollo/v3"
	"go.uber.org/zap"
)

const defaultNamespace = "application"

type apolloConfig struct {
	AppID              string   `toml:"appid"`
	Cluster            string   `toml:"cluster"`
	Address            string   `toml:"address"`
	Namespace          []string `toml:"namespace"`
	CacheDir           string   `toml:"cache_dir"`
	AccesskeySecret    string   `toml:"accesskey_secret"`
	InsecureSkipVerify bool     `toml:"insecure_skip_verify"`
}

func initApollo() {
	node, ok := env.get("apollo").(*toml.Tree)

	if !ok {
		return
	}

	cfg := new(apolloConfig)

	if err := node.Unmarshal(cfg); err != nil {
		logger.Error("yiigo: apollo init error", zap.Error(err))

		return
	}

	if err := agollo.StartWithConf(&agollo.Conf{
		AppID:              cfg.AppID,
		Cluster:            cfg.Cluster,
		NameSpaceNames:     cfg.Namespace,
		CacheDir:           cfg.CacheDir,
		MetaAddr:           cfg.Address,
		AccesskeySecret:    cfg.AccesskeySecret,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
	}); err != nil {
		logger.Error("yiigo: apollo init error", zap.Error(err))

		return
	}

	if !InStrings(defaultNamespace, cfg.Namespace...) {
		cfg.Namespace = append(cfg.Namespace, defaultNamespace)
	}

	env.withApollo(cfg.Namespace)

	logger.Info("yiigo: apollo is OK.")
}
