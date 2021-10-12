package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("YESH")
	for {
		var b bytes.Buffer
		fmt.Print(">")
		var text string
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			text = scanner.Text()
		}

		err := Exec(&b, text)
		if err != nil {
			log.Printf("exec error: %v\n", err)
		}
		_, err = io.Copy(os.Stdout, &b)
		if err != nil {
			log.Printf("exec error: %v\n", err)
		}
	}
}

func Execute(outputBuffer *bytes.Buffer, stack ...*exec.Cmd) (err error) {
	var errorBuffer bytes.Buffer
	pipeStack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		stdinPipe, stdoutPipe := io.Pipe()
		stack[i].Stdout = stdoutPipe
		stack[i].Stderr = &errorBuffer
		stack[i+1].Stdin = stdinPipe
		pipeStack[i] = stdoutPipe
	}
	stack[i].Stdout = outputBuffer
	stack[i].Stderr = &errorBuffer

	if err := call(stack, pipeStack); err != nil {
		log.Println(string(errorBuffer.Bytes()), err)
	}
	return err
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}

func Exec(b *bytes.Buffer, mArgv string) error {
	argvSlice := strings.Split(mArgv, "|")
	var (
		err error
	)

	var NewCmd = func(commandRaw string) (*exec.Cmd, error) {
		args := strings.Fields(commandRaw)
		command := args[0]
		if len(args) == 1 {
			cmd := exec.Command(command)
			return cmd, nil
		} else {
			args = args[1:]
			arg := strings.Join(args, " ")
			cmd := exec.Command(command, arg)
			return cmd, nil
		}
	}

	commandChain := make([]*exec.Cmd, 0, len(argvSlice))

	for _, rawCommand := range argvSlice {
		cmd, err := NewCmd(rawCommand)
		if err != nil {
			return err
		}
		commandChain = append(commandChain, cmd)
	}
	err = Execute(b, commandChain...)
	if err != nil {
		return err
	}
	return err
}

/*func ForkExec(mArgv string) error{
	argvSlice := strings.Split(mArgv, "|")
	var(
		w *os.File
		r = os.Stdin
		attr syscall.ProcAttr
		pid int
	)

	var NewCmd = func (id int, out *os.File) error{
		var err error
		commandRaw := argvSlice[id]



		attr  = syscall.ProcAttr{
			Env:   []string{},
			Files: []uintptr{r.Fd(), out.Fd(), os.Stderr.Fd()},
		}
		args := strings.Fields(commandRaw)
		command := args[0]
		if len(args) == 1 {
			args = []string{""}
		}else{
			args = args[1:]
		}
		log.Println(args)

		var path string

		path, err = exec.LookPath(command)
		if err != nil {
			return err
		}

		pid, err = syscall.ForkExec(path,args,&attr)
		if err != nil {
			return err
		}

		r = w
		return nil
	}

	w, err := os.Create(fmt.Sprintf("pipe%d", 0))
	if err != nil {
		return err
	}

	defer w.Close()
	if len(argvSlice)==1{
		w = os.Stdout
	}
	err = NewCmd(0, w)

	if err != nil{
		return err
	}

	for i := 1; i < len(argvSlice)-1; i++ {
		log.Println("soooo.....")
		w, err = os.Create(fmt.Sprintf("pipe%d", 0))
		if err != nil {
			return err
		}
		defer w.Close()

		err = NewCmd(i,w)
		if err != nil{
			return err
		}
	}
	if len(argvSlice) >= 2{
		err = NewCmd(len(argvSlice)-1, os.Stdout)
		if err != nil{
			return err
		}
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	_, err = proc.Wait()
	return err
} */
