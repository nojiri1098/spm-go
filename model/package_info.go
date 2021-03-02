package model

type PackageInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Files []string `json:"files"`
	FilesCount int `json:"files_count"`
	Dependencies *DependenciesInfo `json:"dependencies"`
	Dependants []string `json:"dependants"`
	AfferentCoupling int `json:"afferent_coupling"`
	EfferentCoupling int `json:"efferent_coupling"`
	Instability float64 `json:"instability"`
}

type PackagesSummary struct {
	Packages []*PackageInfo `json:"packages"`
}

type DependenciesInfo struct {
	Standard []string `json:"standard"`
	Internals []string `json:"internals"`
	Externals []string `json:"externals"`
	StandardCount int `json:"standard_count"`
	InternalsCount int `json:"internals_count"`
	ExternalsCount int `json:"externals_count"`
	TotalCount int `json:"count"`
}
