package usecase

import (
	"fmt"

	"github.com/enuesaa/loadii/internal/services/serve"
)

func (u *useCaseImpl) Serve(dir string, port int) {
	servectl := serve.New(u.deps)
	servectl.Port = port
	servectl.Basepath = dir

	fmt.Printf("Listening on %s\n", servectl.Addr())

	if err := servectl.Listen(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}
