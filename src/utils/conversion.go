package utils

import (
	"github.com/spacerouter/marketplace/models"
	"strconv"
	"strings"
)

func ConvertToStack(compose models.Compose) (*models.Stack, error) {
	networks, err := convertToNetworkDeclarations(compose.Networks)
	if err != nil {
		return nil, err
	}

	volumes, err := convertToVolumeDeclarations(compose.Volumes)
	if err != nil {
		return nil, err
	}

	services, err := convertToServices(compose.Services)
	if err != nil {
		return nil, err
	}
	return &models.Stack{
		Networks: networks,
		Volumes:  volumes,
		Services: services,
	}, nil
}

func convertToNetworkDeclarations(networks map[string]models.ComposeNetworkDeclaration) ([]models.NetworkDeclaration, error) {
	var returnNetworks []models.NetworkDeclaration
	for name, _ := range networks {
		returnNetworks = append(returnNetworks, models.NetworkDeclaration{
			Name: name,
		})
	}

	return returnNetworks, nil
}

func convertToVolumeDeclarations(networks map[string]models.ComposeVolumeDeclaration) ([]models.VolumeDeclaration, error) {
	var returnVolumes []models.VolumeDeclaration
	for name, _ := range networks {
		returnVolumes = append(returnVolumes, models.VolumeDeclaration{
			Name: name,
		})
	}

	return returnVolumes, nil
}

func convertToServices(services map[string]models.ComposeService) ([]models.Service, error) {
	var returnService []models.Service
	for name, service := range services {
		ports, err := convertToPorts(service.Ports)

		if err != nil {
			return nil, err
		}

		envs, err := convertToEnvs(service.Environment)
		if err != nil {
			return nil, err
		}

		vol, err := convertToVolumes(service.Volumes)
		if err != nil {
			return nil, err
		}

		networks, err := convertToNetworks(service.Networks)
		if err != nil {
			return nil, err
		}

		imageInfo := strings.Split(service.Image, ":")
		imageVersion := "latest"
		if len(imageInfo) > 1 {
			imageVersion = imageInfo[1]
		}

		returnService = append(returnService, models.Service{
			Name:         name,
			Image:        imageInfo[0],
			ImageVersion: imageVersion,
			Ports:        ports,
			Envs:         envs,
			Volumes:      vol,
			Networks:     networks,
		})
	}

	return returnService, nil
}

func convertToPorts(ports []string) ([]models.Port, error) {
	var returnPorts []models.Port
	for _, value := range ports {
		parts := strings.Split(value, ":")
		port, err := strconv.ParseUint(parts[0], 10, 32)
		if err != nil {
			return nil, err
		}
		returnPorts = append(returnPorts, models.Port{
			Port: uint(port),
		})
	}

	return returnPorts, nil
}

func convertToEnvs(ports map[string]string) ([]models.EnvVar, error) {
	var returnEnvs []models.EnvVar
	for name, value := range ports {
		returnEnvs = append(returnEnvs, models.EnvVar{
			Name:         name,
			DefaultValue: value,
		})
	}

	return returnEnvs, nil
}

func convertToVolumes(ports []string) ([]models.Volume, error) {
	var returnVolumes []models.Volume
	for _, value := range ports {
		parts := strings.Split(value, ":")

		returnVolumes = append(returnVolumes, models.Volume{
			Name:       parts[0],
			MountPoint: parts[0],
		})
	}

	return returnVolumes, nil
}

func convertToNetworks(ports []string) ([]models.Network, error) {
	var returnNetworks []models.Network
	for _, value := range ports {
		returnNetworks = append(returnNetworks, models.Network{
			Name: value,
		})
	}

	return returnNetworks, nil
}
