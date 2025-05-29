// Package kafka provides a wrapper around the kafka-go library for simplified Kafka operations.
package kafka

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// Reader wraps the kafka-go Reader to provide a simplified interface for reading messages from Kafka topics.
type Reader struct {
	*kafka.Reader
}

// NewReader creates a new Kafka reader instance with the specified configuration.
// It sets up a consumer group with the given groupID and subscribes to the provided topics.
//
// Parameters:
//   - cfg: Kafka configuration containing bootstrap servers and other settings
//   - topics: List of topics to subscribe to
//   - groupID: Consumer group ID for the reader
//
// Returns:
//   - *Reader: A new Kafka reader instance
//   - error: Any error that occurred during reader creation
func NewReader(
	cfg *Config,
	topics []Topic,
	groupID string,
) (*Reader, error) {
	logger := zerolog.DefaultContextLogger
	if logger == nil {
		l := log.With().Caller().Logger()
		logger = &l
	}

	topicsArr := make([]string, len(topics))
	for i, t := range topics {
		topicsArr[i] = t.String()
	}

	logger.Info().Msgf("setting up kafka reader in group %s", groupID)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.BootstrapServers,
		GroupTopics: topicsArr,
		GroupID:     groupID,
	})

	return &Reader{r}, nil
}

// Close closes the Kafka reader and releases any associated resources.
// It should be called when the reader is no longer needed.
func (r *Reader) Close() error {
	return r.Reader.Close()
}
