package core

import "go.uber.org/fx"

func CreateInstance(shutdowner fx.Shutdowner, m *Core) error {
	defer shutdowner.Shutdown()
	return m.CreateInstance()
}
