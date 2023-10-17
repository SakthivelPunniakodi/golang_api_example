// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package utils

import "strings"

func ToEnvString(strArray [][]string) string {
	s := []string{}
	for _, str := range strArray {
		s = append(s, strings.Join(str, "="))
	}

	return strings.Join(s, "|")
}
