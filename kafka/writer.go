// Package kafka provides a wrapper around the kafka-go library for simplified Kafka operations.
package kafka

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Writer wraps the kafka-go Writer to provide a simplified interface for writing messages to Kafka topics.
type Writer struct {
	*kafka.Writer
}

// NewWriter creates a new Kafka writer instance with the specified configuration.
// It sets up a writer with optimized settings for message delivery:
// - Uses LeastBytes balancer for efficient message distribution
// - Configures a 100ms batch timeout for message batching
// - Connects to the specified Kafka bootstrap servers
//
// Parameters:
//   - cfg: Kafka configuration containing broker addresses and other settings
//
// Returns:
//   - *Writer: A new Kafka writer instance
//   - error: Any error that occurred during writer creation
func NewWriter(
	cfg *Config,
) (*Writer, error) {
	logger := zerolog.DefaultContextLogger
	if logger == nil {
		l := log.With().Caller().Logger()
		logger = &l
	}

	logger.Info().Msgf("setting up kafka writer")
	w := &kafka.Writer{
		Addr:         kafka.TCP(cfg.BootstrapServers...),
		Balancer:     &kafka.LeastBytes{},
		Logger:       logger,
		BatchTimeout: 100 * time.Millisecond,
	}

	return &Writer{w}, nil
}

// Close closes the Kafka writer and releases any associated resources.
// It should be called when the writer is no longer needed.
func (w *Writer) Close() error {
	return w.Writer.Close()
}
