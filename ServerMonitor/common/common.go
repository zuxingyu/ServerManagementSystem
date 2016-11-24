package monitor

import (
	"os/exec"
	"time"
	"bytes"
	"log"
	"errors"
	"runtime"
	"strings"
	"net/url"
	"path"
	"io/ioutil"
	"os"
)

var (
	Timeout = 3 * time.Second        // 设置超时时间
	ErrTimeout = errors.New("Command执行超时")
)

type Invoker interface {
	Command(string, ...string) ([]byte, error)
}

type Invoke struct{}

func (i Invoker) Command(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return CombinedOutputTimeout(cmd, Timeout)
}

type FakeInvoke struct {
	CommandExpectedDir string // CommandExpectedDir specifies dir which includes expected outputs.
	Suffix             string // Suffix species expected file name suffix such as "fail"
	Error              error  // If Error specfied, return the error.
}

// Command in FakeInvoke returns from expected file if exists.
func (i FakeInvoke) Command(name string, arg ...string) ([]byte, error) {
	if i.Error != nil {
		return []byte{}, i.Error
	}

	arch := runtime.GOOS

	fname := strings.Join(append([]string{name}, arg...), "")
	fname = url.QueryEscape(fname)
	var dir string
	if i.CommandExpectedDir == "" {
		dir = "expected"
	} else {
		dir = i.CommandExpectedDir
	}
	fpath := path.Join(dir, arch, fname)
	if i.Suffix != "" {
		fpath += "_" + i.Suffix
	}
	if PathExists(fpath) {
		return ioutil.ReadFile(fpath)
	}
	return exec.Command(name, arg...).Output()
}

func PathExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}


// 这个方法用来运行执行所给的command及timeout
// copied from https://github.com/shirou/gopsutil
func CombinedOutputTimeout(c *exec.Cmd, timeout time.Duration) ([]byte, error) {
	var b bytes.Buffer
	c.Stdout = &b
	c.Stderr = &b
	if err := c.Start(); err != nil {
		return nil, err
	}
	err := WaitTimeout(c, timeout)
	return b.Bytes(), err
}

// 等待命令执行
func WaitTimeout(c *exec.Cmd, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	done := make(chan error)
	go func() {
		done <- c.Wait()
	}()
	select {
	case err := <-done:
		timer.Stop()
		return err
	case <-timer.C:
		if err := c.Process.Kill(); err != nil {
			log.Printf("FATAL error killing process: %s", err)
			return err
		}
		// wait for the command to return after killing it
		<-done
		return ErrTimeout
	}
}