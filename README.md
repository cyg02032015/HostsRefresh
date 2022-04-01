# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.185.185.171 github.com
54.203.127.103 api.github.com
34.214.218.151 github.map.fastly.net
34.214.218.151 github.githubassets.com
54.188.39.197 github.global.ssl.fastly.net
54.213.92.171 raw.githubusercontent.com
54.212.182.63 camo.githubusercontent.com
52.27.10.225 avatars.githubusercontent.com
34.216.98.40 avatars0.githubusercontent.com
54.185.222.202 avatars1.githubusercontent.com
54.213.92.171 avatars2.githubusercontent.com
34.222.127.126 avatars3.githubusercontent.com
54.185.222.202 avatars4.githubusercontent.com
54.213.30.146 avatars5.githubusercontent.com
54.185.222.202 avatars6.githubusercontent.com
34.222.31.48 avatars7.githubusercontent.com
34.216.98.40 avatars8.githubusercontent.com
35.88.84.179 favicons.githubusercontent.com
34.222.31.48 developer.apple.com
52.11.157.220 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-04-01 09:12:58

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder