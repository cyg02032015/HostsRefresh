# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
34.200.241.206 github.com
35.173.236.59 api.github.com
23.20.153.37 github.map.fastly.net
34.207.185.226 github.githubassets.com
3.83.127.204 github.global.ssl.fastly.net
54.152.75.220 raw.githubusercontent.com
35.153.179.147 camo.githubusercontent.com
52.20.78.184 avatars.githubusercontent.com
174.129.96.170 avatars0.githubusercontent.com
54.161.22.80 avatars1.githubusercontent.com
3.236.140.185 avatars2.githubusercontent.com
3.235.8.179 avatars3.githubusercontent.com
54.159.16.174 avatars4.githubusercontent.com
3.89.141.130 avatars5.githubusercontent.com
3.80.90.73 avatars6.githubusercontent.com
35.153.179.147 avatars7.githubusercontent.com
52.20.78.184 avatars8.githubusercontent.com
35.173.236.59 favicons.githubusercontent.com
44.200.234.99 developer.apple.com
54.161.229.103 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-01-07 08:52:59

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder