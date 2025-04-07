## 구현해야 할 사항.
- ProcessScript,GenerateExecutor 에 대응해야 함. 여러번 병렬적으로 구동될 수 있음.
- 서버 구동중 디렉토리를 생각해야함. 결국 계산 서버에서 caleb 이 실행되기때문에, 
- 파이프라인 별로 디렉토리를 구성하고, 내부적으로 이용하는 것들과 외부 공개할 수 있는 디렉토리를 구별해서 설계를 해야함. (중요)
- grpc 설계 및 클라이언트로 넘어오는 녀석 파싱하는 것도 여기에 있어야 함.
- 이미지 및 컨테이너를 만들어 가는 과정들을 설계를 잘 해야함. 지금은 그냥 podbridge5 에서 해버리는데 여기서 몇몇 메서드들을 가져와서, 이미지나 컨테이너 만드는 전단계 
- 과정을 잘 설계해야함. (중요)
- https://github.com/goccy/go-yaml 설치 필요.
- pipeline.json 설계 및 제작 중 지속적으로 업데이트 필요.
- 파이프라인과 tori 연결해줘야 한다.  
- 유전체 분석 파이프라인을 만들어가면서 수정해나가야 함.  

### pipeline_v3
- 삭제 아래 필드.
-  "memorySwap"
-  "pidsLimit" 
- 추가 되어야 할 것이, 데이터 블럭에 대한 내용.
- 각 노드의 제한시간.
- pod 에 대한 내용.
- 아래 내용 참고해서 1차적으로 완성하고 보완하자.

```aiignore
func (c *ContainerCollection) RunE(a interface{}) (int, error) {
	node, ok := a.(*dag.Node)
	if !ok {
		return 9, fmt.Errorf("invalid input: expected *dag.Node")
	}


	var containerConfig *Container
	for i, cont := range c.Containers {
		if cont.NodeID == node.Id {
			containerConfig = &c.Containers[i]
			break
		}
	}
	if containerConfig == nil {
		return 9, fmt.Errorf("no container configuration found for node id: %s", node.Id)
	}

	return 0, nil
}
```