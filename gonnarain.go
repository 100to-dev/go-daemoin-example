package gonnarain

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/100to-dev/go-daemon-example/internal/core"
	"github.com/godbus/dbus/v5"
)

func Run(ctx context.Context, cfg *Config, logger *log.Logger) error {
	logger.Printf("starting gonnarain service")

	dbusConn, err := dbus.ConnectSessionBus()

	if err != nil {
		return err
	}

	defer dbusConn.Close()

	notificationFunc := func(msg string) error {
		return sendNotification(dbusConn, msg)
	}

	gonnarainService := core.BuildGonnarainService(core.GonnarainConfig{
		Method:           core.DefaultCoinFlip,
		NotificationFunc: notificationFunc,
	})

	idbus := dbusInterface{notificationFunc: notificationFunc}

	if err := exportDbusInterface(dbusConn, idbus); err != nil {
		return fmt.Errorf("dbus export error: %w", err)
	}

	ticker := time.NewTicker(cfg.Interval)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := gonnarainService(); err != nil {
				logger.Printf("gonnarain service error: %s", err)
			}
		}
	}
}

type dbusInterface struct {
	notificationFunc core.NotificationFunc
}

func (di dbusInterface) DisasterAlert(msg string) *dbus.Error {
	if err := di.notificationFunc(msg); err != nil {
		return &dbus.Error{
			Name: "NotificationError",
			Body: []interface{}{"Cannot alert people about the disaster"},
		}
	}

	return nil
}
