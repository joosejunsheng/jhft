package core

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/joosejunsheng/jhft/config"
	"github.com/joosejunsheng/jhft/log"
	"github.com/joosejunsheng/jhft/proxy"
)

var CoreCmd = &cobra.Command{
	Use:   "hft",
	Short: "run hft",
	Long:  "run hft",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	var (
		err    error
		logger *zap.Logger
		p      proxy.Func
	)

	logText := os.Getenv("LOG_LEVEL")
	if logText == "" {
		logText = "INFO" // Default to INFO if not set
	}

	logLevel, err := zapcore.ParseLevel(logText)
	if err != nil {
		panic(err)
	}
	plugin := log.NewStdoutPlugin(logLevel)
	logger = log.NewLogger(plugin)
	logger.Info("log init end")

	// set zap global logger
	zap.ReplaceGlobals(logger)

	client := &http.Client{}

	// url := "/api/v1/market/orderbook/level1?symbol=BTC-USDT"
	// endpoint := "/api/v2/user-info"
	endpoint := "/api/v1/market/orderbook/level2_20?symbol=BTC-USDT"
	domain := "https://api.kucoin.com"
	req, err := http.NewRequest("GET", domain+endpoint, nil)
	if err != nil {
		fmt.Printf("Error Request : %s - %s \n", endpoint, err)
	}

	for {
		tsNow := strconv.Itoa(int(time.Now().UTC().Unix() * 1000))

		apiKey := config.KucoinConf.APIKey
		apiSecret := config.KucoinConf.APISecret
		passPhrase := config.KucoinConf.PassPhrase

		req.Header.Set("KC-API-KEY", apiKey)
		req.Header.Set("KC-API-SIGN", generateSignature(tsNow, "GET", endpoint, "", apiSecret))
		req.Header.Set("KC-API-TIMESTAMP", tsNow)
		req.Header.Set("KC-API-PASSPHRASE", generatePassphrase(apiSecret, passPhrase))
		req.Header.Set("KC-API-KEY-VERSION", "3")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error Request : %s - %s \n", domain+endpoint, err)
		}

		bodyReader := bufio.NewReader(resp.Body)
		e := DetermineEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

		a, _ := io.ReadAll(utf8Reader)
		bla := string(a)

		fmt.Println(bla)

		time.Sleep(500 * time.Millisecond)
	}
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)

	if err != nil {
		zap.L().Error("fetch failed", zap.Error(err))

		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

func generatePassphrase(apiSecret, passPhrase string) string {

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(passPhrase))
	encryptedPassphrase := h.Sum(nil)

	passphraseBase64 := base64.StdEncoding.EncodeToString(encryptedPassphrase)

	return passphraseBase64
}

func generateSignature(tsNow, method, endpoint, body, apiSecret string) string {
	secret := []byte(apiSecret)
	signature := tsNow + method + endpoint + body

	h := hmac.New(sha256.New, secret)

	h.Write([]byte(signature))

	hash := h.Sum(nil)

	signatureBase64 := base64.StdEncoding.EncodeToString(hash)

	return signatureBase64
}
