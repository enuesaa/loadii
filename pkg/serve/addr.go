package serve

import (
	"fmt"
)

func (ctl *Servectl) Addr() string {
	return fmt.Sprintf(":%d", ctl.Port)
}
