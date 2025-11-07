package helpers

import "testing"

func TestNormalizeString_RemovesNonAlnum(t *testing.T) {
    got, err := NormalizeString("12-34.56 AB!@#")
    if err != nil {
        t.Fatalf("erro inesperado: %v", err)
    }
    want := "123456 AB"
    if got != want {
        t.Fatalf("NormalizeString: got=%q want=%q", got, want)
    }
}

