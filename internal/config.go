package internal

import "github.com/lifesoftserv/telnyx/config"

/*
*	holds the telnyx configuration
 */
type Config struct {
	Api       config.ApiKeys
	Messaging config.Messaging
}
