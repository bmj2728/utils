package strutil

import "testing"

func TestNewComparisonManager(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"CM1"},
		{"CM2"},
		{"CM3"},
		{"CM4"},
		{"CM5"},
		{"CM6"},
		{"CM7"},
		{"CM8"},
		{"CM9"},
		{"CM10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := NewComparisonManager()
			if cm.ComparisonResults == nil ||
				cm.SimilarityResults == nil ||
				cm.ShingleResults == nil ||
				cm.LCSResults == nil {
				t.Errorf("NewComparisonManager() = %v, want %v", cm, nil)
			}
		})
	}
}
