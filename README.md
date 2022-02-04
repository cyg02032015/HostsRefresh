# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.161.213.156 github.com
54.167.24.47 api.github.com
18.234.182.75 github.map.fastly.net
54.167.24.47 github.githubassets.com
54.167.24.47 github.global.ssl.fastly.net
3.236.241.207 raw.githubusercontent.com
54.227.180.174 camo.githubusercontent.com
3.89.232.108 avatars.githubusercontent.com
3.216.29.28 avatars0.githubusercontent.com
3.86.34.50 avatars1.githubusercontent.com
44.201.77.250 avatars2.githubusercontent.com
54.161.213.156 avatars3.githubusercontent.com
44.200.19.99 avatars4.githubusercontent.com
44.201.180.133 avatars5.githubusercontent.com
54.161.213.156 avatars6.githubusercontent.com
44.195.34.240 avatars7.githubusercontent.com
3.228.21.63 avatars8.githubusercontent.com
54.161.5.223 favicons.githubusercontent.com
3.237.1.235 developer.apple.com
100.26.176.49 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-02-04 08:45:23

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder