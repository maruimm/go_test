package main

import (
	"context"
	"github/maruimm/myGoLearning/tcpServer"
	"os"
	"os/signal"
	"syscall"
)


func setupSig(cancelFunc context.CancelFunc) {

	sigRecv := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRecv, sigs...)
	for sig := range sigRecv {
		tcpServer.PrintLog(tcpServer.LOGDEBUG, "recv a sig:%+v\n", sig)
		cancelFunc()
		break
	}

}

func main() {
	ctx, stopServer := context.WithCancel(context.Background())
	s, err := tcpServer.NewTcpServer(ctx,  "",
	 8899)
	if err != nil {
		tcpServer.PrintLog(tcpServer.LOGDEBUG, "error:%+v\n",err)
		return
	}

	go s.Run()
	setupSig(stopServer)
	s.ShutDown()
	tcpServer.PrintLog(tcpServer.LOGDEBUG, "main groutine exit\n")
}
