```aiignore

/analysis_root
│
├── pipeline_A
│   ├── config.yaml                  # 파이프라인 설정 파일
│   ├── scripts_template/            # 파이프라인별 템플릿 스크립트
│   │   ├── executor.sh              # 컨테이너 실행 스크립트 (공통 사용 가능)
│   │   ├── install.sh               # 파이프라인 특화된 의존성 설치 스크립트
│   │   ├── healthcheck.sh           # 컨테이너 헬스체크 스크립트 (공통 사용 가능)
│   │   └── user_script_template.sh  # 사용자 작성 분석 스크립트 템플릿
│   │
│   ├── runs/                        # 분석 진행 중 임시 디렉토리 (완료 후 삭제)
│   │   ├── 20250315_e2a5f2/         # 날짜_짧은UUID 형식으로 생성 -> 성공하면 아래 result 로 넘어가지만 통합. 이건 컨테이너 하나 정보
│   │   │   ├── executor.sh
│   │   │   ├── install.sh
│   │   │   ├── healthcheck.sh
│   │   │   ├── user_script.sh       # 샘플의 데이터 경로가 기록된 스크립트
│   │   │   ├── tmp/                 # 분석 중 생성된 임시 데이터
│   │   │   ├── logs/                # 분석 중 생성된 로그 파일
│   │   │   └── status.json          # 현재 상태 (성공, 실패 등)
│   │   └── 20250315_b7c8e3/
│   │       └── ...
│   │
│   ├── results/                     # 분석이 완료된 결과 데이터
│   │   ├── fileblock_20250315_e2a5f2/ -> 컨테이너 이름이랑 같이 설정. 이건 여러 컨테이너가 통합된 결과.
│   │   │   ├── A/                   # 파이프라인의 단계별 폴더
│   │   │   │   ├── sample01_A_result.txt
│   │   │   │   ├── sample02_A_result.txt
│   │   │   │   └── ...
│   │   │   ├── B/
│   │   │   │   ├── sample01_B_result.txt
│   │   │   │   └── ...
│   │   │   ├── C/
│   │   │   │   ├── sample01_C_result.txt
│   │   │   │   └── ...
│   │   │   └── metadata.json        # 블록의 메타데이터 기록
│   │   │
│   │   └── fileblock_20250316_b7d3c9/
│   │       ├── A/
│   │       ├── B/
│   │       ├── C/
│   │       └── metadata.json
│   │
│   └── failed/                      # 실패한 분석 작업의 기록 (로그 등)
│       ├── 20250315_f8c3d9/
│       │   └── logs/
│       └── ...
│
├── pipeline_B                       # pipeline_B는 독립적으로 동일 구조
│   ├── config.yaml
│   ├── scripts_template/
│   │   ├── executor.sh              # 공통 script 사용 가능 (심볼릭 링크 추천)
│   │   ├── install.sh               # 파이프라인 B의 고유한 의존성 관리
│   │   ├── healthcheck.sh           # 공통 또는 개별 사용 가능
│   │   └── user_script_template.sh  # pipeline_B의 분석 스크립트 템플릿
│   ├── runs/
│   ├── results/
│   └── failed/
│
└── common_scripts                   # 여러 파이프라인이 공유 가능한 스크립트
├── executor.sh
└── healthcheck.sh

```