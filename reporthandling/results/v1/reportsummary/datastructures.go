package reportsummary

// SummaryDetails detailed summary of the scanning. will contain versions, counters, etc.
type SummaryDetails struct {
	Frameworks       []FrameworkSummary        `json:"frameworks"`         // list of framework summary
	Controls         map[string]ControlSummary `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	ResourceCounters `json:",inline"`
}

// FrameworkSummary summary of scanning from a single framework perspective
type FrameworkSummary struct {
	Name             string                    `json:"name"`               // framework name
	Score            float32                   `json:"score"`              // framework score
	Version          string                    `json:"version"`            // framework version
	Controls         map[string]ControlSummary `json:"controls,omitempty"` // mapping of control - map[<control ID>]<control summary>
	ResourceCounters `json:",inline"`
}

// ControlSummary summary of scanning from a single control perspective
type ControlSummary struct {
	Name             string  `json:"name"`
	Score            float32 `json:"score"`
	ResourceCounters `json:",inline"`
}

type ResourceCounters struct {
	PassedResources   int `json:"passedResources"`
	FailedResources   int `json:"failedResources"`
	ExcludedResources int `json:"excludedResources"`
	SkippedResources  int `json:"skippedResources"`
}
