package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func NowPlaying() {
	if isSpotifyIsPlay() {
		printCurrentPlaying()
	}
}

func isSpotifyIsPlay() bool {
	cmd := exec.Command(
		"dbus-send",
		"--print-reply",
		"--dest=org.mpris.MediaPlayer2.spotify",
		"/org/mpris/MediaPlayer2",
		"org.freedesktop.DBus.Properties.Get",
		"string:org.mpris.MediaPlayer2.Player",
		"string:PlaybackStatus",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return false
	}

	output := out.String()
	statusRegex := regexp.MustCompile(`variant.*?"(.*?)"`)
	status := statusRegex.FindStringSubmatch(output)[1]

	return status == "Playing"
}

func printCurrentPlaying() {
	cmd := exec.Command(
		"dbus-send",
		"--print-reply",
		"--dest=org.mpris.MediaPlayer2.spotify",
		"/org/mpris/MediaPlayer2",
		"org.freedesktop.DBus.Properties.Get",
		"string:org.mpris.MediaPlayer2.Player",
		"string:Metadata",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	output := out.String()

	titleRegex := regexp.MustCompile(`(?s)title.*?variant.*?"(.*?)"`)
	artistRegex := regexp.MustCompile(`(?s)artist.*?variant.*?"(.*?)"`)

	titleMatch := titleRegex.FindStringSubmatch(output)
	artistMatch := artistRegex.FindStringSubmatch(output)

	if len(titleMatch) > 1 && len(artistMatch) > 1 {
		title := titleMatch[1]
		artist := artistMatch[1]
		fmt.Printf("%s - %s\n", artist, title)
	}
}
