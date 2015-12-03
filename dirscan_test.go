package main_test

import "testing"

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}