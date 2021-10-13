package kill

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func YetAnotherKill(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

func YetAnotherPwd() error {
	output, err := filepath.Abs(".")
	if err == nil {
		fmt.Print(output)
	}

	return err
}

func YetAnotherCd(path string) error {
	return os.Chdir(path)
}

func YetAnotherEcho(text string) {
	if strings.HasPrefix(text, "$") {
		fmt.Print(os.Getenv(strings.TrimPrefix(text, "$")))
	} else {
		fmt.Print(text)
	}
}

func ParseProc() ([]*os.Process, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer func(d *os.File) {
		err := d.Close()
		if err != nil {

		}
	}(d)

	results := make([]*os.Process, 0, 50)
	for {
		names, err := d.Readdirnames(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, name := range names {
			// We only care if the name starts with a numeric
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			// From this point forward, any errors we just ignore, because
			// it might simply be that the process doesn't exist anymore.
			pid, err := strconv.Atoi(name)
			if err != nil {
				continue
			}

			p, err := os.FindProcess(pid)
			if err != nil {
				continue
			}

			results = append(results, p)
		}
	}

	return results, nil
}

func YetAnotherPs() error {
	procs, err := ParseProc()
	if err != nil {
		return err
	}
	for _, proc := range procs {
		_, err = fmt.Fprintf(os.Stdout, "%d", proc.Pid)
	}

	return nil
}
