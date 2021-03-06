package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/whitekid/go-utils/log"
)

func TestSignal(t *testing.T) {
	signals := []os.Signal{syscall.SIGUSR1}
	defer signal.Reset(signals...)

	go func() {
		<-time.After(time.Second)
		log.Debugf("timeout 1 second")

		process, err := os.FindProcess(os.Getpid())
		require.NoError(t, err)
		log.Debugf("send signal to me, pid=%d", os.Getpid())
		process.Signal(syscall.SIGUSR1)
	}()

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)

	<-SetupSignal(ctx, signals...).Done()
	cancel()

	require.Fail(t, "@@@@")
}
