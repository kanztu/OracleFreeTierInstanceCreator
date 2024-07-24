package core

import (
	"github.com/kanztu/OracleFreeTierInstanceCreator/oci"
	"go.uber.org/zap"
)

type Core struct {
	oci    *oci.OCI
	logger *zap.Logger
}

func New(c *oci.OCI, logger *zap.Logger) *Core {
	return &Core{c, logger}
}

func (c *Core) CreateInstance() error {
	c.logger.Info("Creating instance")
	id, err := c.oci.CreateInstance()
	if err != nil {
		return err
	}
	c.logger.Info("Created instance", zap.String("id", id))
	return nil
}
