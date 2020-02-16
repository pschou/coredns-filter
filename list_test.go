package filter

import (
	"testing"

	"github.com/caddyserver/caddy"
)

func TestFilter_Load(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{`filter {
			allow ./testdata/whitelist.txt
			block ./testdata/blacklist.txt
		}`, false},
	}
	for _, tt := range tests {
		c := caddy.NewTestController("dns", tt.input)
		f, err := parseConfig(c)
		if err != nil {
			t.Fatal(err)
		}

		if err := f.Load(); (err != nil) != tt.wantErr {
			t.Errorf("Filter.Load() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
