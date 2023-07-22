package color

import "fmt"

var (
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

func ColorStr(str interface{}, color string) string {
	return fmt.Sprintf("%s%v%s", color, str, Reset)
}
