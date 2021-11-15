# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/permission-door.proto](#npool/permission-door.proto)
    - [AuthenticateRolePolicyRequest](#permission.door.v1.AuthenticateRolePolicyRequest)
    - [AuthenticateRolePolicyResponse](#permission.door.v1.AuthenticateRolePolicyResponse)
    - [AuthenticateRolesPolicyRequest](#permission.door.v1.AuthenticateRolesPolicyRequest)
    - [AuthenticateRolesPolicyResponse](#permission.door.v1.AuthenticateRolesPolicyResponse)
    - [AuthenticateUserPolicyByIDRequest](#permission.door.v1.AuthenticateUserPolicyByIDRequest)
    - [AuthenticateUserPolicyByIDResponse](#permission.door.v1.AuthenticateUserPolicyByIDResponse)
    - [DeleteResourceRequest](#permission.door.v1.DeleteResourceRequest)
    - [DeleteResourceResponse](#permission.door.v1.DeleteResourceResponse)
    - [DeleteRoleRequest](#permission.door.v1.DeleteRoleRequest)
    - [DeleteRoleResponse](#permission.door.v1.DeleteRoleResponse)
    - [GetRolePoliciesRequest](#permission.door.v1.GetRolePoliciesRequest)
    - [GetRolePoliciesResponse](#permission.door.v1.GetRolePoliciesResponse)
    - [NoDataResponse](#permission.door.v1.NoDataResponse)
    - [PageInfo](#permission.door.v1.PageInfo)
    - [Policy](#permission.door.v1.Policy)
    - [ResourcePolicy](#permission.door.v1.ResourcePolicy)
    - [RolePolicies](#permission.door.v1.RolePolicies)
    - [SetRolePoliciesRequest](#permission.door.v1.SetRolePoliciesRequest)
    - [SetRolePoliciesResponse](#permission.door.v1.SetRolePoliciesResponse)
    - [UnsetRolePoliciesRequest](#permission.door.v1.UnsetRolePoliciesRequest)
    - [UnsetRolePoliciesResponse](#permission.door.v1.UnsetRolePoliciesResponse)
    - [VersionResponse](#permission.door.v1.VersionResponse)
  
    - [PermissionDoor](#permission.door.v1.PermissionDoor)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/permission-door.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/permission-door.proto



<a name="permission.door.v1.AuthenticateRolePolicyRequest"></a>

### AuthenticateRolePolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| Policy | [ResourcePolicy](#permission.door.v1.ResourcePolicy) |  |  |
| AppId | [string](#string) |  |  |






<a name="permission.door.v1.AuthenticateRolePolicyResponse"></a>

### AuthenticateRolePolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.AuthenticateRolesPolicyRequest"></a>

### AuthenticateRolesPolicyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| RoleIds | [string](#string) | repeated |  |
| Policy | [ResourcePolicy](#permission.door.v1.ResourcePolicy) |  |  |






<a name="permission.door.v1.AuthenticateRolesPolicyResponse"></a>

### AuthenticateRolesPolicyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.AuthenticateUserPolicyByIDRequest"></a>

### AuthenticateUserPolicyByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ResourecID | [string](#string) |  |  |
| Action | [string](#string) |  |  |






<a name="permission.door.v1.AuthenticateUserPolicyByIDResponse"></a>

### AuthenticateUserPolicyByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| ResourceIds | [string](#string) | repeated |  |






<a name="permission.door.v1.DeleteResourceResponse"></a>

### DeleteResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| RoleIds | [string](#string) | repeated |  |






<a name="permission.door.v1.DeleteRoleResponse"></a>

### DeleteRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.GetRolePoliciesRequest"></a>

### GetRolePoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="permission.door.v1.GetRolePoliciesResponse"></a>

### GetRolePoliciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RolePolicies](#permission.door.v1.RolePolicies) |  |  |






<a name="permission.door.v1.NoDataResponse"></a>

### NoDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageCount | [int32](#int32) |  |  |
| Total | [int32](#int32) |  |  |






<a name="permission.door.v1.Policy"></a>

### Policy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppId | [string](#string) |  |  |
| UserId | [string](#string) |  |  |
| ResourceId | [string](#string) |  |  |
| Action | [string](#string) |  |  |






<a name="permission.door.v1.ResourcePolicy"></a>

### ResourcePolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceId | [string](#string) |  |  |
| Action | [string](#string) |  |  |






<a name="permission.door.v1.RolePolicies"></a>

### RolePolicies



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Policies | [ResourcePolicy](#permission.door.v1.ResourcePolicy) | repeated |  |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |






<a name="permission.door.v1.SetRolePoliciesRequest"></a>

### SetRolePoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| ResourcePolicies | [ResourcePolicy](#permission.door.v1.ResourcePolicy) | repeated |  |
| AppId | [string](#string) |  |  |






<a name="permission.door.v1.SetRolePoliciesResponse"></a>

### SetRolePoliciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.UnsetRolePoliciesRequest"></a>

### UnsetRolePoliciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleId | [string](#string) |  |  |
| AppId | [string](#string) |  |  |
| Policies | [ResourcePolicy](#permission.door.v1.ResourcePolicy) | repeated |  |






<a name="permission.door.v1.UnsetRolePoliciesResponse"></a>

### UnsetRolePoliciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="permission.door.v1.VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="permission.door.v1.PermissionDoor"></a>

### PermissionDoor
rpc Echo (StringMessage) returns (StringMessage){
    option (google.api.http) = {
        post: &#34;/v1/echo&#34;
        body: &#34;*&#34;
    };
}

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#permission.door.v1.VersionResponse) | Method Version |
| SetRolePolicies | [SetRolePoliciesRequest](#permission.door.v1.SetRolePoliciesRequest) | [SetRolePoliciesResponse](#permission.door.v1.SetRolePoliciesResponse) | Set policies to Role. |
| AuthenticateRolePolicy | [AuthenticateRolePolicyRequest](#permission.door.v1.AuthenticateRolePolicyRequest) | [AuthenticateRolePolicyResponse](#permission.door.v1.AuthenticateRolePolicyResponse) | Authenticate Role permission. |
| AuthenticateRolesPolicy | [AuthenticateRolesPolicyRequest](#permission.door.v1.AuthenticateRolesPolicyRequest) | [AuthenticateRolesPolicyResponse](#permission.door.v1.AuthenticateRolesPolicyResponse) |  |
| GetRolePolicies | [GetRolePoliciesRequest](#permission.door.v1.GetRolePoliciesRequest) | [GetRolePoliciesResponse](#permission.door.v1.GetRolePoliciesResponse) | Get all policies of role. |
| UnsetRolePolicies | [UnsetRolePoliciesRequest](#permission.door.v1.UnsetRolePoliciesRequest) | [UnsetRolePoliciesResponse](#permission.door.v1.UnsetRolePoliciesResponse) | Cancel some policies of role. |
| DeleteRole | [DeleteRoleRequest](#permission.door.v1.DeleteRoleRequest) | [DeleteRoleResponse](#permission.door.v1.DeleteRoleResponse) | Delete data of role. |
| DeleteResource | [DeleteResourceRequest](#permission.door.v1.DeleteResourceRequest) | [DeleteResourceResponse](#permission.door.v1.DeleteResourceResponse) | Delete resources. |
| AuthenticateUserPolicyByID | [AuthenticateUserPolicyByIDRequest](#permission.door.v1.AuthenticateUserPolicyByIDRequest) | [AuthenticateUserPolicyByIDResponse](#permission.door.v1.AuthenticateUserPolicyByIDResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

