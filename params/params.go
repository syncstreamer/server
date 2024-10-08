package params

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var InAddr string
var OutAddr string
var ServeStatic bool
var TimeframeDuration int
var TimeframeHistoryItems int
var CertPath string
var CertPrivateKeyPath string
var UseTLS bool

func ReadParams() {
	const (
		defaultTimeframeDuration     = 10_000
		defaultTimeframeHistoryItems = 5
	)
	inAddrEnv, _ := os.LookupEnv("SYNCSTREAMER_IN_ADDRESS")
	flag.StringVar(&InAddr, "in_addr", inAddrEnv, "Inbound address \"[host]:[port]\"")

	outAddrEnv, _ := os.LookupEnv("SYNCSTREAMER_OUT_ADDRESS")
	flag.StringVar(&OutAddr, "out_addr", outAddrEnv, "Outbound address \"[host]:[port]\"")

	serveStaticEnv, _ := os.LookupEnv("SYNCSTREAMER_SERVE_STATIC")
	serveStaticEnvBool := serveStaticEnv == "true"
	flag.BoolVar(&ServeStatic, "serve_static", serveStaticEnvBool,
		"set to true if the server should serve client static too, default: false")

	timeframeDurationEnv, _ := os.LookupEnv("SYNCSTREAME_TIMEFRAME_DURATION")
	timeframeDurationEnvInt, _ := strconv.ParseInt(timeframeDurationEnv, 10, 64)
	flag.IntVar(&TimeframeDuration, "timeframe_duration", int(timeframeDurationEnvInt),
		fmt.Sprintf("timeframe duration in ms, default: %d", defaultTimeframeDuration))

	timeframeHistoryItemsEnv, _ := os.LookupEnv("SYNCSTREAM_TIMEFRAME_HISTORY_ITEMS")
	timeframeHistoryItemsEnvInt, _ := strconv.ParseInt(timeframeHistoryItemsEnv, 10, 64)
	flag.IntVar(&TimeframeHistoryItems, "timeframe_history_items", int(timeframeHistoryItemsEnvInt),
		fmt.Sprintf("timeframe history items number, default: %d", defaultTimeframeHistoryItems))

	certPathEnv, _ := os.LookupEnv("SYNCSTREAMER_CERT_PATH")
	flag.StringVar(&CertPath, "cert_path", certPathEnv, "TLS cert path")

	certPrivateKeyEnv, _ := os.LookupEnv("SYNCSTREAME_CERT_PRIVATE_KEY_PATH")
	flag.StringVar(&CertPrivateKeyPath, "cert_private_key_path", certPrivateKeyEnv, "TLS private key path")

	useTLSEnv, _ := os.LookupEnv("SYNCSTREAME_USE_TLS")
	useTLSEnvBool := useTLSEnv == "true"
	flag.BoolVar(&UseTLS, "use_tls", useTLSEnvBool,
		"set to true if the server should use TLS, default: false")

	flag.Parse()

	if InAddr == "" || OutAddr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if TimeframeDuration == 0 {
		TimeframeDuration = defaultTimeframeDuration
	}

	if TimeframeHistoryItems == 0 {
		TimeframeHistoryItems = defaultTimeframeHistoryItems
	}
}
