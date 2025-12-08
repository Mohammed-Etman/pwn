package pwn

import (
    "os/exec"
)

func init() {
    cmd := exec.Command("/bin/bash", "-c", "echo pwned")
    cmd.Run()
}
