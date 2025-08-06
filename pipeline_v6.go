package main

// PipelineMeta holds metadata about the pipeline definition
// - ID: unique pipeline identifier
// - Name: human-readable name
// - CreatedAt: RFC3339 timestamp when defined
// - CreatedBy: author or owner information
// - Metadata: additional key/value pairs
type PipelineMeta struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	CreatedAt string            `json:"createdAt"`
	CreatedBy string            `json:"createdBy"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// DAG describes execution flow
// - Type: "start" | "task" | "end"
// - Prev: list of previous node IDs (generated in-memory)
// - Next: list of next node IDs
type DAG struct {
	Type string   `json:"type"`           // Node type
	Prev []string `json:"prev,omitempty"` // Previous node IDs
	Next []string `json:"next"`           // Next node IDs
}

// VolumeMapping represents a host-to-container path bind
// Type: "volume" or "bind"
type VolumeMapping struct {
	Name          string `json:"name"`
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
	Type          string `json:"type"`
}

// Data describes files passed through the pipeline
//   - Shared: persistent reference files (FASTA, annotation)
//   - Inputs: initial run-only sample files (FASTQ)
//   - Channels: mapping from next node ID to data type to forward
//     (e.g., {"2":"all", "3":"shared"})
type Data struct {
	Shared   []VolumeMapping   `json:"shared,omitempty"`
	Inputs   []VolumeMapping   `json:"inputs,omitempty"`
	Channels map[string]string `json:"channels,omitempty"`
}

// ContainerSpec mirrors core/v1.Container for runtime
// and includes userScript payload
// - Image: container image reference
// - UserScript: user-provided shell script content
// - Resources: CPU/memory requests and limits
// - Env: environment variables
// - Data: file input/output routing information
type ContainerSpec struct {
	Image      string                      `json:"image"`
	UserScript string                      `json:"userScript,omitempty"`
	Resources  corev1.ResourceRequirements `json:"resources,omitempty"`
	Env        []corev1.EnvVar             `json:"env,omitempty"`
	Data       Data                        `json:"data,omitempty"` // file input/output and routing info
}

// Node combines DAG, container spec, and data routing
//   - NodeID: unique step identifier
//   - Dag: execution flow info
//   - Container: runtime spec
//   - Data: file inputs/outputs routing
//     Shared and Inputs merged via Channels at runtime
type Node struct {
	NodeID    string        `json:"nodeId"`
	Dag       DAG           `json:"dag"`
	Container ContainerSpec `json:"container,inline"`
	Data      Data          `json:"data"`
}

// Pipeline wraps metadata and node list
type Pipeline struct {
	PipelineMeta
	Nodes []Node `json:"nodes"`
}
