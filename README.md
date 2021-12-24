# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
52.26.248.149 github.com
35.165.98.27 api.github.com
35.88.114.192 github.map.fastly.net
35.88.159.207 github.githubassets.com
34.212.175.183 github.global.ssl.fastly.net
54.201.238.116 raw.githubusercontent.com
35.85.34.95 camo.githubusercontent.com
34.212.175.183 avatars.githubusercontent.com
35.163.22.2 avatars0.githubusercontent.com
35.85.34.95 avatars1.githubusercontent.com
52.10.38.237 avatars2.githubusercontent.com
35.163.142.216 avatars3.githubusercontent.com
35.88.159.207 avatars4.githubusercontent.com
35.88.159.207 avatars5.githubusercontent.com
35.88.114.192 avatars6.githubusercontent.com
34.213.8.170 avatars7.githubusercontent.com
35.165.98.27 avatars8.githubusercontent.com
35.163.22.2 favicons.githubusercontent.com
52.26.248.149 developer.apple.com
35.165.87.150 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-12-24 08:45:42

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder