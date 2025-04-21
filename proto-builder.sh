#!/usr/bin/env bash
# 스크립트 내에서 명령어가 실패하면 즉시 종료
set -e

# 이미지 이름과 Dockerfile 이름 설정 (Dockerfile 이름을 Dockerfile_Make 로 했으면 그걸 사용)
IMAGE_NAME="proto-builder"
DOCKERFILE="Dockerfile_Make"  # 만약 Dockerfile 이름이 기본이면 "Dockerfile"로 변경

# 현재 스크립트가 있는 디렉토리로 이동 (빌드 컨텍스트가 올바르게 설정되도록)
cd "$(dirname "$0")"

echo "Building Podman image ${IMAGE_NAME} using ${DOCKERFILE}..."
podman build -t "${IMAGE_NAME}" -f "${DOCKERFILE}" .

echo "Podman image built successfully."

# 디버깅을 위해 호스트의 protos 폴더를 컨테이너의 /app/protos 에 마운트
# -it: 인터랙티브 모드, --rm: 종료 후 컨테이너 자동 삭제
# --privileged podman rootless 사용하기 때문에 make 사용할려면 써야 함.
echo "Running Podman container in interactive mode..."
podman run -it --privileged --rm -v "$(pwd)/protos":/app/protos "${IMAGE_NAME}"
