package v1alpha1

import "encoding/json"

type RuleJSON struct {
	ID       string `json:"id"`
	RuleSpec `json:",inline"`
}

type UpstreamJSON struct {
	URL          string  `json:"url"`
	StripPath    *string `json:"strip_path,omitempty"`
	PreserveHost bool    `json:"preserve_host"`
}

func (rj RuleJSON) MarshalJSON() ([]byte, error) {

	type Alias RuleJSON

	if rj.Upstream.PreserveHost == nil {

	}

	return json.Marshal(&struct {
		Upstream *UpstreamJSON `json:"upstream"`
		Alias
	}{
		Upstream: &UpstreamJSON{
			URL:          rj.Upstream.URL,
			PreserveHost: parsePreserveHost(rj.Upstream.PreserveHost),
			StripPath:    rj.Upstream.StripPath,
		},
		Alias: (Alias)(rj),
	})
}

func parsePreserveHost(b *bool) bool {

	if b == nil {
		return false
	}

	return *b
}
