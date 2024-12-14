/*
 * Copyright (c) 2024, licensed under the EUPL-1.2-or-later
 */

package server

import "math/rand"

// randomString generates a random string of length n
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
