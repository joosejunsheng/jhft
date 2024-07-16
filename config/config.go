package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type KucoinConfig struct {
	APIKey     string `yaml:"api_key"`
	APISecret  string `yaml:"api_secret"`
	PassPhrase string `yaml:"pass_phrase"`
}

type BinanceConfig struct {
	APIKey     string `yaml:"api_key"`
	APISecret  string `yaml:"api_secret"`
	PassPhrase string `yaml:"pass_phrase"`
}

type MySQLConfig struct {
}

var KucoinConf *KucoinConfig
var BinanceConf *BinanceConfig
var MySQLConf MySQLConfig

func InitializeConfig() {
	kucoinConfig, kucoinConfigErr := loadKucoinConfig("config/kucoin.yaml")
	if kucoinConfigErr != nil {
		fmt.Println(kucoinConfigErr)
	}
	KucoinConf = kucoinConfig

	binanceConfig, binanceConfigErr := loadBinanceConf("config/binance.yaml")
	if binanceConfigErr != nil {
		fmt.Println(binanceConfigErr)
	}
	BinanceConf = binanceConfig
}

func loadKucoinConfig(path string) (*KucoinConfig, error) {
	var config *KucoinConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func loadBinanceConf(path string) (*BinanceConfig, error) {
	var config *BinanceConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func loadMySQLConfig() (MySQLConfig, error) {
	var config MySQLConfig

	return config, nil
}
