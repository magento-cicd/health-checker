package options

import "github.com/sirupsen/logrus"

// The options accepted by this CLI tool
type Options struct {
	Ports    []int
	Listener string
	Logger   *logrus.Logger
}
