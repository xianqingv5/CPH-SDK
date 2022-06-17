package main

import (
	"CPH-SDK/conf"
	"CPH-SDK/cph/adb"
	"CPH-SDK/cph/bandwidth"
	"CPH-SDK/cph/codingservice"
	"CPH-SDK/cph/cphservers"
	"CPH-SDK/cph/keypair"
	"CPH-SDK/cph/phoneinstance"
	"CPH-SDK/cph/task"
	"net/http"
)

func main() {
	// 配置信息
	conf.Init()

	// 查询带宽信息
	http.HandleFunc("/test", bandwidth.QueryBandwidth)
	// 重启云手机
	http.HandleFunc("/cphserver", cphservers.RestartCloudPhoneServer)
	http.HandleFunc("/code", codingservice.ListEncodeService)

	// ADB命令
	// 安装apk
	http.HandleFunc("/InstallApk", adb.InstallApk)
	// 推送文件
	http.HandleFunc("/PushFile", adb.PushFile)
	// 执行异步adb shell命令
	http.HandleFunc("/RunShellCommand", adb.RunShellCommand)
	// 执行同步adb shell命令
	http.HandleFunc("/RunSyncCommand", adb.RunSyncCommand)
	// 卸载apk
	http.HandleFunc("/UninstallApk", adb.UninstallApk)

	// 宽带管理
	// 查询带宽信息
	http.HandleFunc("/QueryBandwidth", bandwidth.QueryBandwidth)
	// 修改共享带宽
	http.HandleFunc("/UpdateBandwidth", bandwidth.UpdateBandwidth)

	// 编码服务管理
	// 查询编码服务
	http.HandleFunc("/ListEncodeServers", codingservice.ListEncodeService)
	// 重启编码服务
	http.HandleFunc("/RestartEncodeServer", codingservice.RestartEncodeServer)

	// 云手机服务器管理
	// 购买系统定义网络云手机服务器
	http.HandleFunc("/CreateCloudPhoneServer", cphservers.CreateCloudPhoneServer)
	// 删除共享存储文件
	http.HandleFunc("/DeleteShareFiles", cphservers.DeleteShareFiles)
	// 查询云手机服务器详情
	http.HandleFunc("/GetCloudPhoneServerDetail", cphservers.GetCloudPhoneServerDetail)
	// 查询云手机服务器规格列表
	http.HandleFunc("/GetCloudPhoneServerModels", cphservers.GetCloudPhoneServerModels)
	// 查询云手机服务器列表
	http.HandleFunc("/ListCloudPhoneServers", cphservers.ListCloudPhoneServers)
	// 查询共享存储文件
	http.HandleFunc("/ListShareFiles", cphservers.ListShareFiles)
	// 推送共享存储文件
	http.HandleFunc("/PushShareFiles", cphservers.PushShareFiles)
	// 重启云手机服务器
	http.HandleFunc("/RestartCloudPhoneServer", cphservers.RestartCloudPhoneServer)

	// 密匙管理
	// 更改密钥对
	http.HandleFunc("/UpdateKeypair", keypair.UpdateKeypair)

	// 手机实例管理
	// 查询云手机详情
	http.HandleFunc("/GetCloudPhoneDetail", phoneinstance.GetCloudPhoneDetail)
	// 查询云手机规格列表
	http.HandleFunc("/ListCloudPhoneModels", phoneinstance.ListCloudPhoneModels)
	// 查询云手机列表
	http.HandleFunc("/ListCloudPhones", phoneinstance.ListCloudPhones)
	// 查询手机镜像
	http.HandleFunc("/ListPhoneImages", phoneinstance.ListPhoneImages)
	// 关闭云手机
	http.HandleFunc("/PowerOffCloudPhone", phoneinstance.PowerOffCloudPhone)
	// 重置云手机
	http.HandleFunc("/ResetCloudPhone", phoneinstance.ResetCloudPhone)
	// 重启云手机
	http.HandleFunc("/RestartCloudPhone", phoneinstance.RestartCloudPhone)
	// 更新云手机属性
	http.HandleFunc("/UpdateCloudPhoneProperty", phoneinstance.UpdateCloudPhoneProperty)
	// 修改云手机名称
	http.HandleFunc("/UpdatePhoneName", phoneinstance.UpdatePhoneName)

	// 任务管理
	// 查询任务执行状态列表
	http.HandleFunc("/ListJobs", task.ListJobs)
	// 查询任务执行状态
	http.HandleFunc("/QueryJobs", task.QueryJobs)

	http.ListenAndServe("0.0.0.0:4999", nil)

}
