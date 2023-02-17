# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.3 github.com
140.82.112.3 api.github.com
151.101.1.6 github.map.fastly.net
34.221.24.223 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
34.221.24.223 raw.githubusercontent.com
34.221.24.223 camo.githubusercontent.com
54.149.73.121 avatars.githubusercontent.com
34.217.116.4 avatars0.githubusercontent.com
35.92.189.22 avatars1.githubusercontent.com
35.92.204.234 avatars2.githubusercontent.com
34.221.24.223 avatars3.githubusercontent.com
35.92.189.22 avatars4.githubusercontent.com
54.149.73.121 avatars5.githubusercontent.com
54.149.73.121 avatars6.githubusercontent.com
35.166.207.159 avatars7.githubusercontent.com
35.88.252.172 avatars8.githubusercontent.com
35.88.252.172 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-02-17 09:10:25

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder