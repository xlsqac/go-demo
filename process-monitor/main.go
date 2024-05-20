/*
程序未运行时自动运行
可执行文件为 ~/script/go/xxx-monitor
*/
package main

import (
    "github.com/shirou/gopsutil/process"
    "os/exec"
)

func main() {
    processes, _ := process.Processes()
    var hasPro bool = false

    for _, p := range processes {
        name, _ := p.Name()
        if name == "Xnip" {
            hasPro = true
            break
        }
    }
    if !hasPro {
        cmd := exec.Command("open", "/Applications/Xnip.app")
        err := cmd.Run()
        if err != nil {
            return
        }
    }
}
