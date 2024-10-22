package main

import (
	"bufio"
	"fmt"
	"io"
	"net/netip"
	"os"
	"strings"
)

type Interface struct {
	Name    string
	IP      netip.Addr
	Status  string
	Protocl string
}

const (
	ipHdr    = "IP-Address"
	statHdr  = "Status"
	protoHdr = "Protocol"
)

func parseAddr(addr string) (netip.Addr, error) {
	if addr == "unassigned" {
		return netip.IPv4Unspecified(), nil
	}

	return netip.ParseAddr(addr)
}

func ParseIfaces(r io.Reader) ([]Interface, error) {
	s := bufio.NewScanner(r)
	var ipIdx, statIdx, protoIdx, lnum int
	var ifaces []Interface

	for s.Scan() {
		lnum++
		line := s.Text()
		if lnum == 1 {
			ipIdx = strings.Index(line, ipHdr)
			statIdx = strings.Index(line, statHdr)
			protoIdx = strings.Index(line, protoHdr)

			if ipIdx == -1 || statIdx == -1 || protoIdx == -1 {
				return nil, fmt.Errorf("bad header: %q", s.Text())
			}
			continue
		}

		name := strings.TrimSpace(line[:ipIdx])
		ipStr := strings.TrimSpace(line[ipIdx:statIdx])
		ip, err := parseAddr(ipStr)
		if err != nil {
			return nil, fmt.Errorf("%d: bad IP address - %w", lnum, err)
		}
		status := strings.TrimSpace(line[statIdx:protoIdx])
		proto := strings.TrimSpace(line[protoIdx:])

		iface := Interface{
			Name:    name,
			IP:      ip,
			Status:  status,
			Protocl: proto,
		}
		ifaces = append(ifaces, iface)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return ifaces, nil
}

func main() {
	file, err := os.Open("ip_brief.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	ifaces, err := ParseIfaces(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for _, iface := range ifaces {
		fmt.Printf("%+v\n", iface)
	}
}
