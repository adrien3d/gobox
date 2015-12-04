// +build !windows
package util

import (
	s "syscall"
)

// Fonction spécifique pour windows
func initConn() error {
	return nil
}

func Write(sd s.Handle, _, b []byte) error {
	_, err := s.Write(sd, b)
	return err
}
