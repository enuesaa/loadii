package usecase

import "github.com/enuesaa/loadii/internal/services/serve"

func (u *useCaseImpl) Serve(dir string, port int) error {
	servectl := serve.New(u.deps)
	servectl.Port = port
	servectl.Basepath = dir
	u.deps.Log.Info("Listening on %s", servectl.Addr())

	if err := servectl.Listen(); err != nil {
		return err
	}
	return nil
}
