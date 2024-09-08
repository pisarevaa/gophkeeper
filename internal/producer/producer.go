package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Kafka struct {
	*kafka.Conn
}

func NewProducer(server string, logger *zap.SugaredLogger) *Kafka {
	conn, err := kafka.DialContext(context.Background(), "tcp", server)
	if err != nil {
		logger.Error("Failed to connect kafka: %v", err)
		return nil
	}
	logger.Info("Connected to kafka")
	kafka := &Kafka{conn}
	return kafka
}

func (k *Kafka) CloseConnection() {
	k.Close()
}
