# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
18.234.23.163 github.com
18.212.181.180 api.github.com
18.234.160.116 github.map.fastly.net
34.238.168.155 github.githubassets.com
34.238.168.155 github.global.ssl.fastly.net
44.198.179.126 raw.githubusercontent.com
18.212.181.180 camo.githubusercontent.com
44.197.112.101 avatars.githubusercontent.com
3.209.81.158 avatars0.githubusercontent.com
18.234.251.140 avatars1.githubusercontent.com
3.209.81.158 avatars2.githubusercontent.com
54.159.197.173 avatars3.githubusercontent.com
3.80.108.216 avatars4.githubusercontent.com
44.197.112.101 avatars5.githubusercontent.com
54.167.118.153 avatars6.githubusercontent.com
35.153.156.211 avatars7.githubusercontent.com
35.153.156.211 avatars8.githubusercontent.com
18.208.173.203 favicons.githubusercontent.com
3.80.108.216 developer.apple.com
3.80.108.216 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-02-18 08:56:25

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder