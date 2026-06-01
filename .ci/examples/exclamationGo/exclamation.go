package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/streamnative/pulsar-function-go/pf"
)

func HandleExclamation(ctx context.Context, in []byte) ([]byte, error) {
	// 1. unmarshal []byte to your struct, use any schema you want
	payload := string(in)

	// 2. do your logic
	if fc, ok := pf.FromContext(ctx); ok {
		for _, word := range strings.Split(payload, " ") {
			_ = fc.IncrCounter(word, 1)
			count, _ := fc.GetCounter(word)
			logrus.Infof("got word: %s for %d times", word, count)
		}
		// test user config and secrets
		cfg := fc.GetUserConfValue("configKey")
		sec, err := fc.GetSecret("secretKey")
		if err == nil {
			msg := fmt.Sprintf("config: %v, secret: %s", cfg, *sec)
			_, _ = fc.Publish("persistent://public/default/test-exec-package-serde-extra", []byte(msg))
		}
	}
	data := payload + "!"

	// 3. marshal your struct to []byte
	return []byte(data), nil

}

func main() {
	pf.Start(HandleExclamation)
}
