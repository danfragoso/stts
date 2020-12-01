package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

func main() {
	sendInit()

	for {
		update()
		wait()
	}
}

func sendInit() {
	fmt.Print("{ \"version\": 1 } [")
}

func update() {
	fmt.Print("[")
	sendWindow()
	next()
	sendBattery()
	next()
	sendTime()
	fmt.Print("], ")
}

func next() {
	fmt.Print(", ")
}

func sendWindow() {
	res, _ := exec.Command("xdotool", "getactivewindow", "getwindowname").Output()
	windowTitle := string(res)

	windowTitle = strings.ReplaceAll(windowTitle, "{", "\\{")
	windowTitle = strings.ReplaceAll(windowTitle, "}", "\\}")

	windowTitle = strings.ReplaceAll(windowTitle, "[", "\\[")
	windowTitle = strings.ReplaceAll(windowTitle, "]", "\\]")

	windowTitle = strings.ReplaceAll(windowTitle, "\"", "\\\"")

	fmt.Print(`{"full_text": "` + strings.TrimSpace(windowTitle) + ` "}`)
}

func sendBattery() {
	res, _ := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")
	status := strings.ToUpper(string(res)[:3])

	if status != "FUL" {
		res, _ = ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
		percentage := " " + strings.TrimSpace(string(res)) + "%"

		fmt.Print(`{"full_text": " ` + status + percentage + ` "}`)
	} else {
		fmt.Print(`{"full_text": ""}`)
	}
}

func sendTime() {
	now := time.Now()
	fmt.Print(`{"full_text": " ` + now.Format("Jan, 02 15:04") + `"}`)
}

func wait() {
	time.Sleep(250 * time.Millisecond)
}
