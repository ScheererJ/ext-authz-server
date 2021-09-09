/*
 * SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package auth

import (
	"regexp"
)

// Services holds a regex which all allowed services should match.
type Services string

// Check checks if a svc is allowed.
func (s Services) Check(svc string) (bool, error) {
	match, err := regexp.MatchString(string(s), svc)
	if err != nil {
		return match, err
	}
	return match, nil
}
