package builder

import (
	"strings"
)

type GroupBy []string

func (this GroupBy) IsEmpty() bool {
	return len(this) == 0
}

func (this GroupBy) String() string {
	if this.IsEmpty() {
		return ""
	}

	return strings.Join(this, ",")
}
