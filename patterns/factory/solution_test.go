package factory

import (
	"testing"
)

func TestNewRuleConfigParserFactory(t *testing.T) {
	tests := []struct {
		kind string
		nil  bool
	}{
		{kind: "json"},
		{kind: "yaml"},
		{kind: "toml", nil: true},
	}

	for _, tt := range tests {
		got := NewRuleConfigParserFactory(tt.kind)
		if tt.nil {
			if got != nil {
				t.Fatalf("NewRuleConfigParserFactory(%q) should return nil", tt.kind)
			}
			continue
		}
		if got == nil || got.CreateParser() == nil {
			t.Fatalf("NewRuleConfigParserFactory(%q) returned invalid factory", tt.kind)
		}
	}
}
