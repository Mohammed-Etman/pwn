package pwn

import (
    "os/exec"
)

func init() {
    cmd := exec.Command("whoami")
    cmd.Run()
}
