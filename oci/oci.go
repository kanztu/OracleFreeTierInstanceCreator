package oci

import (
	"context"

	"github.com/kanztu/OracleFreeTierInstanceCreator/config"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

type OCI struct {
	client             core.ComputeClient
	compartment        string
	availabilityDomain string
	displayName        string
	imageID            string
	subnetID           string
	sshKeys            string
}

func New(c *config.Config) (*OCI, error) {
	client, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		return nil, err
	}

	return &OCI{client, c.CompartmentOCID, c.AvailabilityDomain, c.DisplayName, c.ImageOCID, c.SubnetOCID, c.SshKey}, nil
}

func (o *OCI) CreateInstance() (string, error) {
	request := core.LaunchInstanceRequest{
		LaunchInstanceDetails: core.LaunchInstanceDetails{
			CompartmentId:      common.String(o.compartment),
			AvailabilityDomain: common.String(o.availabilityDomain),
			Shape:              common.String("VM.Standard.A1.Flex"),
			ShapeConfig:        &core.LaunchInstanceShapeConfigDetails{Ocpus: common.Float32(4), MemoryInGBs: common.Float32(24)},
			DisplayName:        common.String(o.displayName),
			ImageId:            common.String(o.imageID),
			SubnetId:           common.String(o.subnetID),
			Metadata:           map[string]string{"ssh_authorized_keys": o.sshKeys},
			CreateVnicDetails:  &core.CreateVnicDetails{SubnetId: &o.subnetID, AssignPublicIp: common.Bool(true)},
		},
	}

	response, err := o.client.LaunchInstance(context.Background(), request)
	if err != nil {
		return "", err
	}
	return *response.Id, nil
}
