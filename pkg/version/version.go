package version

import "fmt"

var (
	Version   = ""
	Revision  = ""
	BuildDate = ""
)

func String() string {
	return fmt.Sprintf(
		// e.g.) Version: v0.1.0, Revision: 49b9aa0, BuildDate: 2020-09-28T00:15:27Z
		"Version: %s, Revision: %s, BuildDate: %s",
		Version,
		Revision,
		BuildDate,
	)
}
