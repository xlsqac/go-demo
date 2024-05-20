package main

import (
    "fmt"
    "os/exec"
    "strings"
    "time"
)

const (
    AD    = "ad"
    TRACK = "track"
)

func checkError(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func getCurrentSongType() string {
    // 获取当前歌曲类型
    // 广告：ad，正常歌曲：track
    out, err := exec.Command("osascript", "-e", "tell application \"Spotify\" to get spotify url of current track").Output()
    if err != nil {
        panic(err)
    }
    songType := strings.Split(string(out), ":")[1]
    return songType
}

func modifyVolumeForSongType(songType string) {
    // 根据歌曲类型调整音量
    if songType == AD {
        err := exec.Command("osascript", "-e", "tell application \"Spotify\" to set sound volume to 1").Run()
        checkError(err)
        err = exec.Command("osascript", "-e", "tell application \"Spotify\" to set player state to play").Run()
        checkError(err)
    } else if songType == TRACK {
        err := exec.Command("osascript", "-e", "tell application \"Spotify\" to set sound volume to 100").Run()
        checkError(err)
    }
}

func main() {
    for {
        songType := getCurrentSongType()
        modifyVolumeForSongType(songType)
        time.Sleep(500 * time.Millisecond)
    }
}
