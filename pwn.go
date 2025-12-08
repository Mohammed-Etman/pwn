package pwn

import (
	"fmt"
	"os"
)

func init() {
    os.Exec("whoami")
}
