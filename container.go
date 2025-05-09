package main

import (
	"encoding/json"
	"fmt"
	"github.com/seoyhaein/dag-go"
)

// PipelineConfig: 최상위 JSON 객체 ("pipeline" 키를 포함)
/*type PipelineConfig struct {
	Pipeline Pipeline `json:"pipeline"`
}*/

// Pipeline: 파이프라인 전반적인 정보를 담음
type Pipeline struct {
	Name        string      `json:"name"`
	GlobalFiles GlobalFiles `json:"globalFiles"`
	Nodes       []Node      `json:"nodes"`
}

// GlobalFiles: 파이프라인 전역에서 사용하는 파일 경로들
type GlobalFiles struct {
	ExecutorShell    string `json:"executorShell"`
	DockerfilePath   string `json:"dockerfilePath,omitempty"` // dockerfilePath는 옵션
	HealthcheckShell string `json:"healthcheckShell"`
	InstallShell     string `json:"installShell"`
	UserScriptShell  string `json:"userScriptShell"`
}

// Node: 각 노드(컨테이너)의 기본 정보를 담음
type Node struct {
	NodeID  string       `json:"nodeId,omitempty"`
	Type    string       `json:"type"`
	Next    []string     `json:"next"`
	Details *NodeDetails `json:"details,omitempty"` // start/end 노드는 details가 없을 수 있으므로 pointer 처리
}

// NodeDetails: 노드의 세부 설정 (이미지 빌드와 컨테이너 실행 설정)
type NodeDetails struct {
	Image     ImageConfig     `json:"image"`
	Container ContainerConfig `json:"container"`
}

// ImageConfig: 이미지 빌드에 필요한 설정
type ImageConfig struct {
	ID              string `json:"id"`
	DockerfilePath  string `json:"dockerfilePath,omitempty"` // Dockerfile 경로 (필요한 경우)
	SourceImageName string `json:"sourceImageName"`
	TargetImageName string `json:"targetImageName"`
	ImageSavePath   string `json:"imageSavePath"`
}

// ContainerConfig: 컨테이너 실행에 필요한 모든 설정
type ContainerConfig struct {
	Directories     []string            `json:"directories"`     // 컨테이너 내부에서 생성할 디렉토리 목록
	WorkDir         string              `json:"workDir"`         // 컨테이너의 작업 디렉토리
	ScriptMap       map[string][]string `json:"scriptMap"`       // 각 디렉토리에 복사할 스크립트 파일 목록
	PermissionFiles []string            `json:"permissionFiles"` // 권한 설정이 필요한 파일 목록 (최종 경로 기준)
	Cmd             []string            `json:"cmd"`             // 컨테이너 시작 시 실행할 명령어
	Resources       ResourcesConfig     `json:"resources"`       // 리소스 제한 설정
	Volumes         []VolumeConfig      `json:"volumes"`         // 볼륨 마운트 설정
}

// ResourcesConfig: 컨테이너 리소스 제한 설정
type ResourcesConfig struct {
	CPU      CPUConfig    `json:"cpu"`
	Memory   MemoryConfig `json:"memory"`
	OomScore int          `json:"oomScore"`
}

// CPUConfig: CPU 관련 설정
type CPUConfig struct {
	CPUQuota  int `json:"cpuQuota"`  // 한 주기 동안 사용할 수 있는 최대 CPU 시간
	CPUPeriod int `json:"cpuPeriod"` // CPU 제한 주기
	CPUShares int `json:"cpuShares"` // 상대적인 CPU 가중치
}

// MemoryConfig: 메모리 제한 설정
type MemoryConfig struct {
	MemLimit int `json:"memLimit"` // 메모리 제한 (바이트 단위)
}

// VolumeConfig: 컨테이너와 호스트 간의 볼륨 마운트 설정
type VolumeConfig struct {
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
}

// NewPipelineFromJSON json 을 통해 Pipeline 생성
func NewPipelineFromJSON(data []byte) (*Pipeline, error) {
	var pipeline Pipeline
	if err := json.Unmarshal(data, &pipeline); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &pipeline, nil
}

// CreateDagFromPipeline 파싱한 Pipeline 을 바탕으로 Dag 를 생성
// TODO json node id 와 dag 의 node id 가 같은지 확인해야 함.
// 이것도 dag 인터페이스에 넣어야 함.
func (p *Pipeline) CreateDagFromPipeline() (*dag_go.Dag, error) {
	// DAG 초기화: InitDag 내부에서 start 노드를 생성함.
	dag, err := dag_go.InitDag()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dag: %w", err)
	}

	// JSON 에 정의된 노드들 중, "start"와 "end"는 내부에서 생성되므로 제외하고 나머지 노드를 생성하여 nodeMap 에 저장함
	nodeMap := make(map[string]*dag_go.Node)

	for _, n := range p.Nodes {
		if n.NodeID == "start" || n.NodeID == "end" {
			continue
		}
		node := dag.CreateNode(n.NodeID)
		if node == nil {
			// 이미 생성된 경우 nodeMap 에서 가져옴.
			if existing, ok := nodeMap[n.NodeID]; ok {
				node = existing
			} else {
				return nil, fmt.Errorf("failed to create node with id: %s", n.NodeID)
			}
		}
		nodeMap[node.ID] = node
		// TODO 필요하다면 여기서 n.Details 등을 사용해 추가 정보를 설정할 수 있음. 그런데 이렇게 하면 node 구조체를 수정해줘야 하는데.
	}

	// 각 노드의 "next" 배열을 순회하며 edge 를 생성
	for _, n1 := range p.Nodes {
		// reserved 노드 처리: "start"는 내부 start 노드의 ID로, "end"는 건너뜀.
		// start 노드는 edge 연결을 해줘야 하지만, end 노드의 경우는 FinishDag() 에서 edge 연결을 해줌. 따라서 건너뜀.
		// FinishDag() 에서 end node 도 생성해줌. 그전까지는 일단 dag 에서 end node 는 없음.
		var currentNodeID string
		switch n1.NodeID {
		case "start":
			currentNodeID = dag.StartNode.ID
		case "end":
			continue
		default:
			currentNodeID = n1.NodeID
		}

		for _, nextID := range n1.Next {
			// "start"가 next 에 포함되면 오류를 반환
			if nextID == "start" {
				return nil, fmt.Errorf("node %s cannot reference 'start' in its next array", currentNodeID)
			}
			// "end"는 FinishDag()에서 처리되므로 건너뜀
			if nextID == "end" {
				continue
			}
			// nextID가 JSON 에 정의된 노드인지 확인합
			if _, ok := nodeMap[nextID]; !ok {
				return nil, fmt.Errorf("node with id %s referenced in next is not defined in JSON", nextID)
			}
			// 존재하는 노드들 간에 간선을 생성
			if err := dag.AddEdgeIfNodesExist(currentNodeID, nextID); err != nil {
				return nil, fmt.Errorf("failed to add edge from %s to %s: %w", currentNodeID, nextID, err)
			}
		}
	}

	// FinishDag() 호출 시 내부적으로 end 노드를 생성하고 연결 작업을 수행
	if err := dag.FinishDag(); err != nil {
		return nil, fmt.Errorf("failed to finish dag: %w", err)
	}
	return dag, nil
}

// RunE 여기서 컨테이너 생성하고, 실행해주는 역활을 해줌.
func (p *Pipeline) RunE(a interface{}) error {
	// type assertion: 입력이 *dag.Node 타입인지 검증
	node, ok := a.(*dag_go.Node)
	if !ok {
		return fmt.Errorf("invalid input: expected *dag.Node")
	}

	// Pipeline 내의 노드 목록에서 node.Id와 일치하는 노드를 찾습니다.
	var pipelineNode *Node
	for i, n := range p.Nodes {
		if n.NodeID == node.ID {

			pipelineNode = &p.Nodes[i]
			break
		}
	}
	if pipelineNode == nil {
		return fmt.Errorf("no container configuration found for node id: %s", node.ID)
	}

	// pipelineNode.Details에는 이미지 빌드 및 컨테이너 실행에 필요한 설정이 포함되어 있습니다.
	// 이곳에서 BuildImage(pipelineNode.Details.Image) 또는 RunContainer(pipelineNode.Details.Container) 같은
	// 실제 로직을 호출하여 컨테이너 작업을 수행할 수 있습니다.

	return nil
}
