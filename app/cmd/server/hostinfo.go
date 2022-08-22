package main

import "os"

type HostInfo struct {
	Hostname string `json:"hostname"`
}

func NewHostInfo() HostInfo {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "UNKNOWN"
	}
	return HostInfo{
		Hostname: hostName,
	}
}
