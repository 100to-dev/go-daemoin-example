package core

import "fmt"

type GonnarainService func() error

type NotificationFunc func(string) error

type GonnarainConfig struct {
	Method           CoinFlipMethod
	NotificationFunc NotificationFunc
}

func BuildGonnarainService(cfg GonnarainConfig) GonnarainService {
	return func() error {
		coinSide := cfg.Method()

		if coinSide == Middle {
			return fmt.Errorf("natural disaster comming")
		}

		if coinSide == Tails {
			if err := cfg.NotificationFunc("hmmm, i think its goin to rain"); err != nil {
				return fmt.Errorf("cannot send bad news: %w", err)
			}

			return nil
		}

		if coinSide == Head {
			if err := cfg.NotificationFunc("nah, its not going to rain"); err != nil {
				return fmt.Errorf("cannot send good news: %w", err)
			}

			return nil
		}

		return fmt.Errorf("unknown rain possibility")
	}
}
