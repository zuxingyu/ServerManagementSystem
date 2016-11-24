package monitor

import (
	"log"
	//"ServerManagementSystem/ServerMonitor/common"
	"os/exec"
	"fmt"
	"strings"
)

//var invoke monitor.Invoker

func init() {
	log.SetPrefix("monitor: ")
	log.SetFlags(0)
	//invoke = monitor.Invoke{}
	SysInfo()
}



func SysInfo(){

	//cmd := exec.Command("/bin/sh", "shellscript/OSinfo.sh")
	cmd := exec.Command("/bin/sh", "top")

	out, err := cmd.Output()
	if err != nil {

	}

	fmt.Println(string(out))

	for _ , line := range strings.Split(string(out), "\n") {
		fmt.Println(line)
		//values := strings.Fields(line)
		//if len(values) < 1 {
		//	continue
		//}
		//if strings.HasPrefix(line, "Processes") {
		//	fmt.Println(line)
		//}

	}

	//out, err := invoke.Command("/bin/sh", "shellscript/OSinfo.sh")

	//sysctl, err := exec.LookPath("/usr/sbin/sysctl")
	//if err != nil {
	//
	//}
	//
	//out, err := invoker.Command(sysctl, "machdep.cpu")
	//
	//if err != nil {
	//	fmt.Println("cmd.Output: ", err)
	//	return
	//}
	//fmt.Println(string(out))

}
