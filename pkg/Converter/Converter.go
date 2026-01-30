package converter

import "fmt"

func PortConverter(port int) string {
	return fmt.Sprintf("%04X", port)
}
