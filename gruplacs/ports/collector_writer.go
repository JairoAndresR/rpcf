package ports

import "rpcf/gruplacs"

type CollectorWriter interface {

	// Write the content of the grupLAC scraped in the Queue to be processed.
	Write(payload gruplacs.CollectorPayload) error
}
