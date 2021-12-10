# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.173.27.40 github.com
34.229.14.189 api.github.com
52.70.104.90 github.map.fastly.net
18.206.194.170 github.githubassets.com
3.92.19.247 github.global.ssl.fastly.net
44.193.3.125 raw.githubusercontent.com
54.87.55.13 camo.githubusercontent.com
44.195.83.74 avatars.githubusercontent.com
3.238.108.186 avatars0.githubusercontent.com
52.70.104.90 avatars1.githubusercontent.com
3.86.158.141 avatars2.githubusercontent.com
18.205.157.101 avatars3.githubusercontent.com
54.160.28.223 avatars4.githubusercontent.com
3.83.100.236 avatars5.githubusercontent.com
54.84.195.12 avatars6.githubusercontent.com
18.206.89.115 avatars7.githubusercontent.com
3.92.244.157 avatars8.githubusercontent.com
44.201.0.101 favicons.githubusercontent.com
52.70.104.90 developer.apple.com
3.92.1.128 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-12-10 08:45:40

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder