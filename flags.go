package main

import "flag"

var ip = flag.String("ip", "localhost", "IP address/hostname to bind this service to")
var port = flag.Int("port", 8080, "Port number to bind this service to")
var enableColor = flag.Bool("enablecolor", false, "true to enable console color")
var logLevel = flag.String("loglevel", "debug", "Set minimum log level")
