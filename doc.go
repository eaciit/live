// Copyright 2015 EACIIT

/*
Package live implements a monitoring server and keep that service alive.

Feature :
- 	Ping Service
	This feature will be used for check the service is run or not, and if the service is not running the script will trigger to execute turn on service.
	This feature can run over net, http, REST and in the future will support other method.
- 	Turn on and turn off service
	This feature used for turn off and turn off the service with any method such us local command run, ssh and REST
- 	Mail if any warning or error found
	Send email if any error or warning found
*/
package live // import "github.com/eaciit/live"
