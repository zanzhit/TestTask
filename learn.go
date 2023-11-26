package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func readUrls(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	read := bufio.NewScanner(file)

	for read.Scan() {
		urls = append(urls, read.Text())
	}

	if err := read.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func recordStream(urls []string, done chan struct{}) {
	var commands []*exec.Cmd

	for i, url := range urls {
		timestamp := time.Now().Format("20060102_150405")
		cmdPath := "C:\\Users\\Zanzhit\\Desktop\\ffmpeg\\ffmpeg-6.0-essentials_build\\bin\\ffmpeg.exe"
		outputFileName := fmt.Sprintf("stream_%d_%s.mp4", i+1, timestamp)
		cmdArgs := []string{"-i", url, "-c", "copy", outputFileName}

		cmd := exec.Command(cmdPath, cmdArgs...)

		stdinPipe, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		defer stdinPipe.Close()

		commands = append(commands, cmd)

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		go func() {
			select {
			case <-done:
				_, err := stdinPipe.Write([]byte("q"))
				if err != nil {
					fmt.Printf("Ошибка записи в стандартный ввод команды %d: %v\n", i+1, err)
				}
				err = stdinPipe.Close()
				if err != nil {
					fmt.Printf("Ошибка закрытия stdinPipe для команды %d: %v\n", i+1, err)
				}
			}
		}()
	}

	for i, cmd := range commands {
		err := cmd.Wait()
		if err != nil {
			fmt.Printf("Ошибка ожидания завершения команды %d: %v\n", i+1, err)
		}
	}
}

func main() {
	filePath := "urls.txt"
	urls, err := readUrls(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла: ", err)
		return
	}

	done := make(chan struct{})

	go func() {
		var input string
		fmt.Scanln(&input)
		close(done)
	}()

	recordStream(urls, done)

	fmt.Println("Программа завершена")
}
