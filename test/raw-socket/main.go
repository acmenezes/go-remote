package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/containernetworking/plugins/pkg/ns"
)

func main() {

	pid := "3258834"
	// Get the pods namespace object
	targetNS, err := ns.GetNS("/tmp/proc/" + pid + "/ns/net")

	if err != nil {
		fmt.Printf("Error getting Pod network namespace: %v", err)
	}
	err = targetNS.Do(func(hostNs ns.NetNS) error {
		fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
		f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
		fileInfo, err := f.Stat()
		if err != nil {
			fmt.Printf("%v", err)
			return err
		}
		fmt.Printf("%v", fileInfo)
		return nil
	})
	if err != nil {
		fmt.Printf("%v", err)
	}

}
