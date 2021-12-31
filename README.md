# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.202.118.208 github.com
54.202.150.232 api.github.com
52.13.98.98 github.map.fastly.net
54.202.118.208 github.githubassets.com
54.214.126.192 github.global.ssl.fastly.net
34.219.60.117 raw.githubusercontent.com
20.80.187.157 camo.githubusercontent.com
54.202.118.208 avatars.githubusercontent.com
54.190.67.101 avatars0.githubusercontent.com
34.213.67.189 avatars1.githubusercontent.com
34.209.214.200 avatars2.githubusercontent.com
35.88.188.240 avatars3.githubusercontent.com
54.201.247.15 avatars4.githubusercontent.com
54.202.118.208 avatars5.githubusercontent.com
54.200.198.21 avatars6.githubusercontent.com
52.13.98.98 avatars7.githubusercontent.com
34.212.54.201 avatars8.githubusercontent.com
34.209.221.161 favicons.githubusercontent.com
35.87.28.245 developer.apple.com
34.213.218.207 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-12-31 08:48:10

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder