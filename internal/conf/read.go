package conf

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func read(path string) (config models.Conf) {

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8840")
	viper.SetDefault("THEME", "sand")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("NODEPATH", "")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("ARP_ARGS", "")
	viper.SetDefault("ARP_STRS_JOINED", "")
	viper.SetDefault("IFACES", "")
	viper.SetDefault("TIMEOUT", 120)
	viper.SetDefault("TRIM_HIST", 48)
	viper.SetDefault("SHOUTRRR_URL", "")

	viper.SetDefault("USE_DB", "sqlite")
	viper.SetDefault("PG_CONNECT", "")

	viper.SetDefault("INFLUX_ENABLE", false)

	viper.SetDefault("PROMETHEUS_ENABLE", false)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host = viper.Get("HOST").(string)
	config.Port = viper.Get("PORT").(string)
	config.Theme = viper.Get("THEME").(string)
	config.Color = viper.Get("COLOR").(string)
	config.NodePath = viper.Get("NODEPATH").(string)
	config.LogLevel = viper.Get("LOG_LEVEL").(string)
	config.ArpArgs = viper.Get("ARP_ARGS").(string)
	config.ArpStrs = viper.GetStringSlice("ARP_STRS")
	config.Ifaces = viper.Get("IFACES").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.TrimHist = viper.GetInt("TRIM_HIST")
	config.ShoutURL = viper.Get("SHOUTRRR_URL").(string)

	config.UseDB = viper.Get("USE_DB").(string)
	config.PGConnect = viper.Get("PG_CONNECT").(string)

	config.InfluxEnable = viper.GetBool("INFLUX_ENABLE")
	config.InfluxSkipTLS = viper.GetBool("INFLUX_SKIP_TLS")
	config.InfluxAddr, _ = viper.Get("INFLUX_ADDR").(string)
	config.InfluxToken, _ = viper.Get("INFLUX_TOKEN").(string)
	config.InfluxOrg, _ = viper.Get("INFLUX_ORG").(string)
	config.InfluxBucket, _ = viper.Get("INFLUX_BUCKET").(string)

	config.PrometheusEnable = viper.GetBool("PROMETHEUS_ENABLE")

	joined := viper.Get("ARP_STRS_JOINED").(string)
	// slog.Info("ARP_STRS_JOINED: " + joined)

	if joined != "" {
		config.ArpStrs = strings.Split(joined, ",")
	}

	return config
}
