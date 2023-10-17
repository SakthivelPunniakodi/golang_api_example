// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package env

import (
	"fmt"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/env"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/utils"
)

type Env struct {
	env.Env
}

// ParseEnv.
func ParseEnv() (*Env, error) {
	o := Env{}

	parsers := []env.ParseEnvFn{}

	if err := o.Parse(parsers...); err != nil {
		return nil, fmt.Errorf("Env.Parse: %w", err)
	}

	return &o, nil
}

func (o Env) String() string {
	s := [][]string{}

	return fmt.Sprintf("%s %s", &o.Env, utils.ToEnvString(s))
}
