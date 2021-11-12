# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
23.102.106.214 github.com
3.221.159.243 api.github.com
18.206.236.198 github.map.fastly.net
18.208.126.52 github.githubassets.com
44.198.171.185 github.global.ssl.fastly.net
18.233.156.243 raw.githubusercontent.com
34.229.176.105 camo.githubusercontent.com
3.85.18.236 avatars.githubusercontent.com
18.206.236.198 avatars0.githubusercontent.com
3.239.102.61 avatars1.githubusercontent.com
3.221.159.243 avatars2.githubusercontent.com
54.81.225.58 avatars3.githubusercontent.com
3.238.226.29 avatars4.githubusercontent.com
23.102.106.214 avatars5.githubusercontent.com
18.206.215.175 avatars6.githubusercontent.com
35.172.203.57 avatars7.githubusercontent.com
3.236.253.104 avatars8.githubusercontent.com
34.200.230.36 favicons.githubusercontent.com
3.239.102.61 developer.apple.com
18.207.0.14 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-11-12 08:44:01

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder