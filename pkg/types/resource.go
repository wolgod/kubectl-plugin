package types

import "sort"

// NodeResourceList xxx
type NodeResourceList struct {
	Name     string
	Status   string
	PodCount int

	CPU    float64
	Memory float64

	// request resource
	CPURequests    float64
	MemoryRequests float64

	// request resource usage
	CPURequestsUsage    string
	MemoryRequestsUsage string

	// limit resource
	CPULimits    float64
	MemoryLimits float64

	// limit resource usage
	CPULimitsUsage    string
	MemoryLimitsUsage string
}

// NodeResourceColumnDefinitions xxx
type NodeResourceColumnDefinitions struct {
	Name     string
	Status   string
	PodCount int

	// requests resource
	CPU         string
	CPURequests string
	CPULimits   string

	// limits resource
	Memory         string
	MemoryRequests string
	MemoryLimits   string
}
type ByName []NodeResourceColumnDefinitions

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortNodes sorts nodes by name.
func SortNodes(nodes []NodeResourceColumnDefinitions) {
	sort.Sort(ByName(nodes))
}
