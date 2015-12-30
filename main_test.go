package main

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	cmd("rm -rf tmp")
	cmd("mkdir -p tmp/foo.d")
	cmd("echo bar > tmp/foo.d/bar")
	cmd("echo baz > tmp/foo.d/baz")

	cmd("go run main.go -once -dir tmp/foo.d -file tmp/foo")

	data, err := ioutil.ReadFile("tmp/foo")
  check(err)

	if string(data) != "bar\nbaz\n" {
		t.Errorf("Got: %s", string(data))
	}
}

func cmd(s string) {
  err := exec.Command("sh", "-c", s).Run()
  check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
