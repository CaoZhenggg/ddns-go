# DDNS-GO

获取本地接口的IPv6公网地址，结合windows定时任务，定时更新阿里云云解析记录。

## 背景

之前家里的光猫下接的设备无法获取公网IPv6地址，光猫没有获取到前缀：

![image](https://user-images.githubusercontent.com/23628206/228703304-939c3ca6-a721-496c-8f11-75b2811d63aa.png)

临时解决办法：在NUC的虚拟机里部署headscale，客户端安装tailscale。但是有时穿越效果不佳，延时较高。后来联系移动客服并顺便提出IPv6的事情，移动宽带师傅上门检查，说要夜里升级。升级之后，光猫可以正常获取前缀了，直连的设备以及猫下接的路由器接入的设备都可以获取公网IPv6地址了。

另外需要关闭光猫IPv6防火墙，不然只能ping通，端口都是不通的：

![image](https://user-images.githubusercontent.com/23628206/228704826-7eec8a93-af5c-4c24-a906-ca078cacd7e6.png)


## 使用方法

首先需要确定公网IPv6绑定的接口名称，修改net.InterfaceByName参数：

![image](https://user-images.githubusercontent.com/23628206/228704282-6ac96773-70cc-4789-af97-a5d2110dca95.png)

![image](https://user-images.githubusercontent.com/23628206/228705067-95d6e384-9f18-445f-b231-037091e3e887.png)

在项目根目录执行:

```
> go get .
> go install
```

## windows添加定时任务

“此电脑”右击鼠标 -> “显示更多选项” -> “管理” -> “任务计划程序” -> “创建任务” 

![image](https://user-images.githubusercontent.com/23628206/228712648-5119f10c-2b83-4b39-a3ce-4f812e016019.png)


![image](https://user-images.githubusercontent.com/23628206/228712169-0895b6ba-c140-4276-b781-50d8cdf04801.png)

![image](https://user-images.githubusercontent.com/23628206/228712258-27deb300-9303-4572-8e9d-864292b375f1.png)



