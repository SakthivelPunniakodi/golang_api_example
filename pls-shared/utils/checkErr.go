// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package utils

import (
	"sampleapi/pls-shared/logger"
)

func CheckErr(logger logger.Logger, fn func() error) {
	if err := fn(); err != nil {
		logger.Errorf("CheckErr: %v", err)
	}
}
