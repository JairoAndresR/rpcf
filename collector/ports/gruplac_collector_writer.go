package ports

import (
	"rpcf/collector"
)

type GrupLACCollectorWriter interface {

	// Write the content of the grupLAC scraped in the Queue to be processed.
	Write(payload collector.Payload, queueName string) error
}
