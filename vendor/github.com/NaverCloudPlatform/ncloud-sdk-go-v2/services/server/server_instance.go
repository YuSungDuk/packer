/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type ServerInstance struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 서버명
ServerName *string `json:"serverName,omitempty"`

	// 서버설명
ServerDescription *string `json:"serverDescription,omitempty"`

	// CPU수
CpuCount *int32 `json:"cpuCount,omitempty"`

	// 메모리사이즈
MemorySize *int64 `json:"memorySize,omitempty"`

	// 기본블럭스토리지사이즈
BaseBlockStorageSize *int64 `json:"baseBlockStorageSize,omitempty"`

	// 플랫폼구분
PlatformType *CommonCode `json:"platformType,omitempty"`

	// 로그인키명
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 유료모니터링여부
IsFeeChargingMonitoring *bool `json:"isFeeChargingMonitoring,omitempty"`

	// 공인IP
PublicIp *string `json:"publicIp,omitempty"`

	// 사설IP
PrivateIp *string `json:"privateIp,omitempty"`

	// 서버이미지명
ServerImageName *string `json:"serverImageName,omitempty"`

	// 서버인스턴스상태
ServerInstanceStatus *CommonCode `json:"serverInstanceStatus,omitempty"`

	// 서버인스턴스OP
ServerInstanceOperation *CommonCode `json:"serverInstanceOperation,omitempty"`

	// 서버인스턴스상태명
ServerInstanceStatusName *string `json:"serverInstanceStatusName,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`

	// UPTIME
Uptime *string `json:"uptime,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode,omitempty"`

	// 반납보호여부
IsProtectServerTermination *bool `json:"isProtectServerTermination,omitempty"`

	// portForwarding 공인 Ip
PortForwardingPublicIp *string `json:"portForwardingPublicIp,omitempty"`

	// portForwarding 외부 포트
PortForwardingExternalPort *int32 `json:"portForwardingExternalPort,omitempty"`

	// portForwarding 내부 포트
PortForwardingInternalPort *int32 `json:"portForwardingInternalPort,omitempty"`

	// Zone
Zone *Zone `json:"zone,omitempty"`

	// 리전
Region *Region `json:"region,omitempty"`

	// 기본블록스토리지디스크유형
BaseBlockStorageDiskType *CommonCode `json:"baseBlockStorageDiskType,omitempty"`

	// 기본블록스토리지디스크상세유형
BaseBlockStroageDiskDetailType *CommonCode `json:"baseBlockStroageDiskDetailType,omitempty"`

	// 인터넷라인구분
InternetLineType *CommonCode `json:"internetLineType,omitempty"`

	// 서버인스턴스구분
ServerInstanceType *CommonCode `json:"serverInstanceType,omitempty"`

	// 사용자데이타
UserData *string `json:"userData,omitempty"`

	// ACG리스트
AccessControlGroupList []*AccessControlGroup `json:"accessControlGroupList,omitempty"`

	// 인스턴스태그리스트
InstanceTagList []*InstanceTag `json:"instanceTagList,omitempty"`
}
