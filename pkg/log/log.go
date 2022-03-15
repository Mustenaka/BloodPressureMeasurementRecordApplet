package log

import (
	"sync"
)

var (
	_logger *logger
	once    sync.Once
)
