좋아—지금 EM JSON을 **실행엔 충분**하게 잘 만들었고, 사람이 읽기에도 꽤 깔끔해.
다만 운영/재현/보안/이식성 관점에서 개선하면 좋은 포인트를 **우선순위**로 정리해줄게. (모두 “선택 적용 가능”)

# 최우선 (안전·재현성)

1. **이미지 고정(digest pinning)**

    * `pipeline_imgA:latest` → `pipeline_imgA@sha256:...`
    * 재실행/롤백 시 동일 이미지 보장.

2. **컨테이너 권한 축소**

    * 각 노드에 `user`, `group`, `fsGroup`, `readOnlyRootFilesystem` 등 추가.
    * 예)

      ```json
      "security": { "runAsUser": 1000, "runAsGroup": 1000, "fsGroup": 1000, "readOnlyRootFilesystem": false }
      ```

3. **실행 한도/재시도 정책**

    * 노드별 `execution`: `timeoutSec`, `maxRetries`, `backoffSec`.
    * 예) `"execution": { "timeoutSec": 14400, "maxRetries": 2, "backoffSec": 60 }`

4. **검증 불변식(런타임 전 체크)**

    * `mergedDir == resolve(common.mergedBase)+"/<nodeId>"`
    * `upperDir`와 `workDir` **동일 파일시스템**
    * `containerPath` 유일성 및 존재/쓰기 가능
    * `lowerDirs` **우선순위 명시**(겹칠 때 원하는 쪽이 왼쪽)
    * JSON에 힌트 필드로:

      ```json
      "validate": { "enforceSameFs": true, "checkMergedBase": true, "uniqueContainerPath": true }
      ```

# 구조·DRY

5. **경로 토큰 일관화 & pathStrategy 이름**

    * 지금 `pathStrategy: "literal"`인데 `@baseDir/@id/@runId` 토큰을 쓰고 있어 **헷갈림**.
    * 제안: `"pathStrategy": "template"` 로 바꾸고, 모든 host 경로를 토큰화:

        * `upperDir`: `@baseDir/@runId/@nodeId-upper`
        * `workDir`:  `@baseDir/@runId/@nodeId-work`
        * `mergedDir`: `@baseDir/@id/@runId/merged/@nodeId`
        * `lowerDirs`: `@baseDir/lower`, `@baseDir/@id/@runId/merged/…`
    * 실행 직전 **치환 렌더링** → 리터럴로 변환.

6. **lowerDirs 우선순위 재검토**

    * 현재는 `["/mnt/overlay/lower", "…/merged/1"]` 형태(베이스가 우선).
    * 보통은 **선행 산출물 > 베이스**가 자연스러워서: `["…/merged/1", "/mnt/overlay/lower"]`

7. **inputsPolicy 명세 강화**

    * 지금 `autobindAllTasks: true, mode: ro`만으로는 **어떤 컨테이너 경로에 바인드되는지**가 불명확.
    * 예)

      ```json
      "inputsPolicy": {
        "autobindAllTasks": true,
        "mode": "ro",
        "target": "container",                 // container | host
        "defaultContainerDir": "/mnt",         // /mnt/R1.fastq 등
        "overridePerInput": false
      }
      ```

# 운영·관찰성

8. **렌더 메타(출처·재현성 추적)**

    * 최상단에:

      ```json
      "renderMeta": {
        "renderer": "em-renderer",
        "version": "0.2.0",
        "renderedAt": "2025-08-20T00:00:00Z",
        "sourceDigest": "sha256:..."
      }
      ```

9. **아티팩트/로그 계약**

    * 노드별 산출물/로그 경로를 명시하면 수집·퍼블리시 자동화 쉬움.
    * 예)

      ```json
      "artifacts": [{ "path": "summary.txt", "type": "text", "publish": true }],
      "logs": { "collectStdout": true, "collectStderr": true, "extraPaths": ["list*.txt"] }
      ```

10. **클린업 보존 기간**

    * `"cleanupPolicy": { "onSuccess": "remove", "onFailure": "retain", "retainDays": 7 }`

# 이식성·확장

11. **리소스 표준화**

    * CPU/메모리 단위 일관: CPU는 millicore(`500m`), 메모리는 `Mi` 유지. (지금은 괜찮음)
    * 필요 시 노드별 `nodeSelector/affinity/tolerations`(K8s 환경용) 힌트 섹션 분리:

      ```json
      "scheduling": { "nodeSelector": { "disk": "ssd" } }
      ```

12. **볼륨 정의 재사용**

    * ref.fa 같은 호스트 파일을 최상단 `volumes`로 선언하고 노드에서 참조:

      ```json
      "volumes": { "ref-fa": { "type": "bind", "hostPath": "/data/ref.fa", "accessMode": "ro" } }
      // 노드: { "mounts": [{ "use": "ref-fa", "containerPath": "/mnt/ref.fa" }] }
      ```

13. **프로비저닝을 호스트 단계로 분리(선택)**

    * `start` 컨테이너가 디렉토리 만드는 대신, **호스트 init 단계**(provisioner)로 명확히:

      ```json
      "hostInit": { "mkdirs": ["@baseDir/lower", "@baseDir/@runId", "..."], "chown": [{"path":"...", "uid":1000,"gid":1000}] }
      ```

14. **다중 샘플(매트릭스) 확장 포인트**

    * 후일 수백 샘플 병렬 처리 시를 대비한 자리:

      ```json
      "matrix": { "samples": ["S1","S2",...], "expandOn": ["nodes.1","nodes.2"] }
      ```

---

## 아주 짧은 “패치 예시” (두 가지만 맛보기)

* **pathStrategy & 토큰화**

  ```json
  "storage": {
    "type": "overlay",
    "pathStrategy": "template",
    "mountStrategy": "host-overlay-bind",
    "common": {
      "baseDir": "/mnt/overlay",
      "lowerDir": "@baseDir/lower",
      "mergedBase": "@baseDir/@id/@runId/merged"
    }
  }
  // 예: nodes["3"].storage.host
  { "upperDir": "@baseDir/@runId/@nodeId-upper",
    "workDir":  "@baseDir/@runId/@nodeId-work",
    "mergedDir": "@baseDir/@id/@runId/merged/@nodeId",
    "lowerDirs": ["@baseDir/@id/@runId/merged/1","@baseDir/lower"] }
  ```

* **실행정책 추가**

  ```json
  "nodes": {
    "6": {
      "execution": { "timeoutSec": 14400, "maxRetries": 2, "backoffSec": 60 },
      "security":  { "runAsUser": 1000, "runAsGroup": 1000, "fsGroup": 1000 }
    }
  }
  ```

원하면 위 목록에서 **적용할 항목 번호**를 알려줘. 그 부분만 정확히 패치해서 전체 JSON을 업데이트해줄게.
