package monitor

import (
	"log"
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	log.SetPrefix("monitor: ")
	log.SetFlags(0)
	SysInfo()


}

var invoke Invoker

func SysInfo(){
	//cmd := exec.Command("/bin/sh", "shellscript/OSinfo.sh")
	bytes, err := invoke.Command("/bin/sh", "shellscript/OSinfo.sh")

	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	v, _ := mem.VirtualMemory()
	fmt.Println(v)


	fmt.Println(string(bytes))
}
