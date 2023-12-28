# jcrose-cmdb
使用 goframe-v2 写的cmdb的后端  

## 功能构想  
1 支持控制k8s的所有服务，通过clientgo实现  
2 支持prometheus对服务的监控指标查询，例如 15分钟内Deployment的CPU 内存使用率  
3 支持资产的录入功能，云平台的资产，线下机房的资产  
4 巡检任务的实现，调用云平台的API接口  
5 用户的权限管理  
6 接入ARGO CICD平台的功能  
