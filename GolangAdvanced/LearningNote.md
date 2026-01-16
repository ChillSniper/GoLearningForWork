# Learning Notes of Advanced Golang Programming

## Concurrency Overview

这边主要是区分三个概念：

- 进程：操作系统分配资源的基本单位
- 线程：操作系统调度资源的基本单位
- 协程：可以理解为用户态线程，协程的调度在用户态进行，不需要切换到内核态；所以不需要由操作系统参与，而是由用户自己控制；

### 并发与并行

并行模式：两个任务一直运行，而且是同时运行着（同一时刻！）
并发模式：在一段时间内，两个任务交替执行；这样的话，并发模式是能够在单核 CPU 上进行执行的；

### Context

作用：

- 用于并发的控制，控制协程的优雅退出；
- 上下文信息传递

直白的说，Context就是用来在父子 goRoutine 之间，进行值传递，以及发送 cancel 信号的一种机制；

GetUserInfoByUserId 的实现流程图

```mermaid
sequenceDiagram
    participant Client as 前端/调用方
    participant UserInfo as GetUserInfoByUserId
    participant ProjDAO as Project DAO
    participant UserSDK as 核心用户SDK
    participant GameAPI as 游戏服务器API
    participant Others as 其他服务(ES, 防沉迷等)

    Client->>UserInfo: GET /.../user/info?user_id=1111
    activate UserInfo

    Note right of UserInfo: 准备工作：参数校验、设置超时

    UserInfo->>ProjDAO: FindUniqueProject(...)
    activate ProjDAO
    ProjDAO-->>UserInfo: 返回项目配置(project)
    deactivate ProjDAO

    Note right of UserInfo: 检查项目配置是否完整

    UserInfo->>UserSDK: GetUserInfo(userId)
    activate UserSDK
    UserSDK-->>UserInfo: 返回基础用户信息(info)
    deactivate UserSDK

    Note right of UserInfo: 将基础信息填充到结果中

    par 获取补充信息
        UserInfo->>GameAPI: GetRoleIdByUserId(userId)
        activate GameAPI
        GameAPI-->>UserInfo: 角色信息(roleInfo)
        deactivate GameAPI
    and
        UserInfo->>UserSDK: QueryForbidden(userId)
        activate UserSDK
        UserSDK-->>UserInfo: 封禁状态(forbidden)
        deactivate UserSDK
    and
        UserInfo->>Others: 调用防沉迷、ES、DAO等
        activate Others
        Others-->>UserInfo: 实名、支付、投诉等信息
        deactivate Others
    end

    Note right of UserInfo: 聚合所有信息

    UserInfo-->>Client: 200 OK (完整的用户信息)
    deactivate UserInfo


```