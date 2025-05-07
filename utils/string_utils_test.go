package utils

import "testing"

func TestReverseString(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "Basic string",
            input:    "привет",
            expected: "тевирп",
        },
        {
            name:     "Empty string",
            input:    "",
            expected: "",
        },
        {
            name:     "Single character",
            input:    "а",
            expected: "а",
        },
        {
            name:     "Palindrome",
            input:    "шалаш",
            expected: "шалаш",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := ReverseString(tt.input)
            if result != tt.expected {
                t.Errorf("ReverseString(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
} 