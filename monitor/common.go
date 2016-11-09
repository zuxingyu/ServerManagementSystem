package monitor

import (
	"os/exec"
	"time"
	"bytes"
	"log"
	"errors"
	//"fmt"
)

var (
	Timeout = 3 * time.Second        // 设置超时时间
	ErrTimeout = errors.New("Command执行超时")
)


type Invoker interface {
	Command(string, ...string) ([]byte, error)
}

type Invoke struct{}

func (i Invoke)Command(name string, arg ...string) ([]byte, error) {
	//fmt.Printf("%s-----%s", name, arg)

	cmd := exec.Command(name, arg...)
	return CombinedOutputTimeout(cmd, Timeout)
}
// 这个方法用来运行执行所给的command及timeout
// copied from https://github.com/shirou/gopsutil
func CombinedOutputTimeout(c *exec.Cmd, timeout time.Duration) ([]byte, error) {
	var b bytes.Buffer
	// Stdout and Stderr specify the process's standard output and error.
	//
	// If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	//
	// If Stdout and Stderr are the same writer, at most one
	// goroutine at a time will call Write.
	c.Stdout = &b
	c.Stderr = &b
	if err := c.Start(); err != nil {
		return nil, err
	}
	err := WaitTimeout(c, timeout)
	return b.Bytes(), err
}

// 等待命令执行
func WaitTimeout(c *exec.Cmd, timeOut time.Duration) error{
	timer := time.NewTimer(timeOut)
	done := make(chan error)
	go func() {
		done <- c.Wait()
	}()
	select {
	case err := <- done:
		timer.Stop()
		return err
	case <- timer.C:
		if err := c.Process.Kill(); err != nil {
			log.Printf("执行错误，关闭进程：%s", err)
		}
		// 待杀掉进程后等待命令执行
		<-done
		return ErrTimeout
	}

}