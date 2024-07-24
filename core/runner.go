package core

import (
	"strings"
	"time"

	"github.com/kanztu/OracleFreeTierInstanceCreator/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func CreateInstance(shutdowner fx.Shutdowner, m *Core, c *config.Config) error {
	defer shutdowner.Shutdown()
	for {
		err := m.CreateInstance()
		if err == nil {
			break
		}
		timeSleep := time.Second * time.Duration(c.IntervalSec)
		if strings.Contains(err.Error(), "Out of host capacity") {
			m.logger.Warn("Out of host capacity")
		} else if strings.Contains(err.Error(), "TooManyRequests") {
			timeSleep *= 2
			m.logger.Warn("TooManyRequests")
		} else {
			m.logger.Error("Failed to create instance", zap.Error(err))
		}
		time.Sleep(timeSleep)
		continue
	}
	return nil
}
