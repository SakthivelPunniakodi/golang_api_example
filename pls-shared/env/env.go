// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package env

import (
	"fmt"
	"os"

	ierr "sampleapi/pls-shared/errors"
	"sampleapi/pls-shared/utils"
)

const (
	ServicePort     = "SERVICE_PORT"
	ServiceVersion  = "SERVICE_VERSION"
	ServiceLogLevel = "SERVICE_LOG_LEVEL"
)

// Env of application.
type Env struct {
	Service Service
}

type Service struct {
	Port     string
	Version  string
	LogLevel string
}

type ParseEnvFn func() error

func (o *Env) Parse(parsers ...ParseEnvFn) error {
	parsers = append(
		parsers,
		o.parseServiceEnv,
	)

	for _, parser := range parsers {
		if err := parser(); err != nil {
			return fmt.Errorf("parser: %w", err)
		}
	}

	return nil
}

func (o *Env) parseServiceEnv() error {
	servicePort, ok := os.LookupEnv(ServicePort)
	if !ok {
		return ierr.NewErrorf(ierr.NotFound, "environment variable not found: %s", ServicePort)
	}

	o.Service.Port = servicePort

	serviceVersion, ok := os.LookupEnv(ServiceVersion)
	if !ok {
		return ierr.NewErrorf(ierr.NotFound, "environment variable not found: %s", ServiceVersion)
	}

	o.Service.Version = serviceVersion

	serviceLogLevel, ok := os.LookupEnv(ServiceLogLevel)
	if !ok {
		return ierr.NewErrorf(ierr.NotFound, "environment variable not found: %s", ServiceLogLevel)
	}

	o.Service.LogLevel = serviceLogLevel

	return nil
}

func (e Env) String() string {
	s := [][]string{
		{ServicePort, e.Service.Port},
		{ServiceVersion, e.Service.Version},
		{ServiceLogLevel, e.Service.LogLevel},
	}

	return utils.ToEnvString(s)
}
