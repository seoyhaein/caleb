package main

/*// PipelineConfig는 최상위 JSON 구조를 나타냅니다.
type PipelineConfig struct {
	Pipeline Pipeline `json:"pipeline"`
}

// Pipeline는 파이프라인의 이름, 글로벌 파일, 노드 목록을 포함합니다.
type Pipeline struct {
	Name        string       `json:"name"`
	GlobalFiles GlobalFiles  `json:"globalFiles"`
	Nodes       []NodeConfig `json:"nodes"`
}

// GlobalFiles는 글로벌 스크립트/파일 경로들을 정의합니다.
type GlobalFiles struct {
	ExecutorShell    string `json:"executorShell"`
	HealthcheckShell string `json:"healthcheckShell"`
	InstallShell     string `json:"installShell"`
	UserScriptShell  string `json:"userScriptShell"`
}

// NodeConfig는 각 노드의 정보를 담습니다.
type NodeConfig struct {
	NodeID  string       `json:"nodeId"`
	Type    string       `json:"type"`
	Next    []string     `json:"next"`
	Details *NodeDetails `json:"details,omitempty"`
}

// NodeDetails는 (필요한 경우) 이미지와 컨테이너 정보를 포함합니다.
type NodeDetails struct {
	Image     ImageConfig     `json:"image"`
	Container ContainerConfig `json:"container"`
}

type ImageConfig struct {
	ID              string `json:"id"`
	DockerfilePath  string `json:"dockerfilePath"`
	SourceImageName string `json:"sourceImageName"`
	TargetImageName string `json:"targetImageName"`
	ImageSavePath   string `json:"imageSavePath"`
}

type ContainerConfig struct {
	Directories     []string            `json:"directories"`
	WorkDir         string              `json:"workDir"`
	ScriptMap       map[string][]string `json:"scriptMap"`
	PermissionFiles []string            `json:"permissionFiles"`
	Cmd             []string            `json:"cmd"`
	Resources       ResourcesConfig     `json:"resources"`
	Volumes         []VolumeConfig      `json:"volumes"`
}

type ResourcesConfig struct {
	CPU      CPUConfig    `json:"cpu"`
	Memory   MemoryConfig `json:"memory"`
	OomScore int          `json:"oomScore"`
}

type CPUConfig struct {
	CPUQuota  int `json:"cpuQuota"`
	CPUPeriod int `json:"cpuPeriod"`
	CPUShares int `json:"cpuShares"`
}

type MemoryConfig struct {
	MemLimit int `json:"memLimit"`
}

type VolumeConfig struct {
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
}*/

/*// CreateDagFromPipelineConfig 는 파싱한 PipelineConfig를 바탕으로 dag_go.Dag 를 생성합니다.
func CreateDagFromPipelineConfig(config *PipelineConfig) (*dag_go.Dag, error) {
	// DAG 초기화 (InitDag 내부에서 start 노드를 생성합니다)
	dag, err := dag_go.InitDag()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize dag: %w", err)
	}

	// JSON에 정의된 노드(예약어 제외)를 생성하여 nodeMap에 저장합니다.
	nodeMap := make(map[string]*dag_go.Node)
	for _, nodeCfg := range config.Pipeline.Nodes {
		// "start"와 "end"는 내부 생성되므로 JSON에서는 무시합니다.
		if nodeCfg.NodeID == "start" || nodeCfg.NodeID == "end" {
			continue
		}
		node := dag.CreateNode(nodeCfg.NodeID)
		if node == nil {
			// 이미 생성된 노드라면 nodeMap에서 가져옵니다.
			if existing, ok := nodeMap[nodeCfg.NodeID]; ok {
				node = existing
			} else {
				return nil, fmt.Errorf("failed to create node with id: %s", nodeCfg.NodeID)
			}
		}
		nodeMap[nodeCfg.NodeID] = node
		// 필요하다면 여기서 nodeCfg.Details 등을 사용해 추가 정보를 설정할 수 있습니다.
	}

	// 각 노드의 "next" 배열을 순회하며 간선을 생성합니다.
	for _, nodeCfg := range config.Pipeline.Nodes {
		// reserved 노드에 대해서는 처리 방식 결정:
		// "start"는 내부 start 노드로 대체, "end"는 간선 생성 건너뛰기.
		var currentNodeID string
		switch nodeCfg.NodeID {
		case "start":
			currentNodeID = dag.StartNode.Id
		case "end":
			continue
		default:
			currentNodeID = nodeCfg.NodeID
		}
		for _, nextID := range nodeCfg.Next {
			// "start"가 next에 포함되면 오류: start는 다른 노드가 참조할 수 없습니다.
			if nextID == "start" {
				return nil, fmt.Errorf("node %s cannot reference 'start' in its next array", currentNodeID)
			}
			// "end"는 FinishDag()에서 처리되므로, next에 포함되어 있으면 건너뜁니다.
			if nextID == "end" {
				continue
			}
			// 나머지 노드가 JSON에 정의되어 있는지 확인합니다.
			if _, ok := nodeMap[nextID]; !ok {
				return nil, fmt.Errorf("node with id %s referenced in next is not defined in JSON", nextID)
			}
			// 존재하는 노드끼리 연결합니다.
			if err := dag.AddEdgeT(currentNodeID, nextID); err != nil {
				return nil, fmt.Errorf("failed to add edge from %s to %s: %w", currentNodeID, nextID, err)
			}
		}
	}

	// FinishDag() 호출 시 내부적으로 end 노드를 생성하고 연결 작업을 수행합니다.
	if err := dag.FinishDag(); err != nil {
		return nil, fmt.Errorf("failed to finish dag: %w", err)
	}
	return dag, nil
}*/

func main() {
	/*filename := "pipeline_v3.json"
	data, err := os.ReadFile(filename)

	var pipelineConfig PipelineConfig
	if err := json.Unmarshal(data, &pipelineConfig); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	dag, err := CreateDagFromPipelineConfig(&pipelineConfig)
	if err != nil {
		log.Fatalf("Error creating dag: %v", err)
	}

	fmt.Printf("DAG created successfully with ID: %s\n", dag.Id)
	ctx := context.Background()
	dag.ConnectRunner()
	dag.GetReady(ctx)
	b1 := dag.Start()
	if b1 != true {
		os.Exit(1)
	}

	b2 := dag.Wait(ctx)
	if b2 != true {
		os.Exit(1)
	}*/
	// 필요한 경우 생성된 DAG의 세부 정보를 출력하거나 추가 로직을 진행합니다.
}
