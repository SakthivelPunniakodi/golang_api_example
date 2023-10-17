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
	"strconv"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/env"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/utils"
)

const (
	IS_BEST_FIT = "IS_BEST_FIT"
)

type Env struct {
	env.Env
	IsBestFit bool
}

// ParseEnv.
func ParseEnv() (*Env, error) {
	o := Env{}

	parsers := []env.ParseEnvFn{
		o.parseIsBestFitEnv,
	}

	if err := o.Parse(parsers...); err != nil {
		return nil, fmt.Errorf("Env.Parse: %w", err)
	}

	return &o, nil
}

func (o *Env) parseIsBestFitEnv() error {
	isBestFit, ok := os.LookupEnv(IS_BEST_FIT)
	if !ok {
		return ierr.NewErrorf(ierr.NotFound, "environment variable not found: %s", IS_BEST_FIT)
	}

	bestFit, _ := strconv.ParseBool(isBestFit)

	o.IsBestFit = bestFit

	return nil
}

func (o Env) String() string {
	s := [][]string{
		{IS_BEST_FIT, strconv.FormatBool(o.IsBestFit)},
	}

	return fmt.Sprintf("%s %s", &o.Env, utils.ToEnvString(s))
}
