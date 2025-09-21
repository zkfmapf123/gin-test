package concurrency

import "go.uber.org/zap"

func Dispatcher(logger zap.Logger, w <-chan func()) {
	logger.Info("Background job dispatcher start...")

	for {
		select {
		case job, ok := <-w:

			if !ok {
				logger.Info("Job channel closed, stopping dispatcher...")
				return
			}

			logger.Info("Processing background job")

			go func() {

				defer func() {
					if r := recover(); r != nil {
						logger.Error("Job panic recovered,", zap.Any("error", r))
					}
				}()
				job()
			}()
		}
	}
}
