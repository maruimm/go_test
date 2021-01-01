package tcpServer

import "fmt"

type LogLevel int

const (
	_ LogLevel = iota
	LOGFANAL
	LOGERROR
	LOGWARING
	LOGINFO
	LOGDEBUG

)

func PrintLog(_ LogLevel, formatter string, args ...interface{}) {
	fmt.Printf(formatter, args...)
}
