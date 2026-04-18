
package main

import "testing"


func TestGetMessage(t *testing.T) {
    want := "Hello"
    got := GetMessage()
    if got != want {
        t.Errorf("want %q, got %q", want, got)
    }
}