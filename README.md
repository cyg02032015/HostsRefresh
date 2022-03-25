# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
35.172.232.10 github.com
18.207.4.77 api.github.com
52.91.194.230 github.map.fastly.net
44.193.199.224 github.githubassets.com
3.231.159.178 github.global.ssl.fastly.net
3.88.10.250 raw.githubusercontent.com
3.82.151.250 camo.githubusercontent.com
3.231.156.154 avatars.githubusercontent.com
54.80.194.225 avatars0.githubusercontent.com
3.83.255.192 avatars1.githubusercontent.com
3.94.203.229 avatars2.githubusercontent.com
3.85.237.59 avatars3.githubusercontent.com
35.153.132.254 avatars4.githubusercontent.com
34.239.226.73 avatars5.githubusercontent.com
18.232.148.115 avatars6.githubusercontent.com
18.232.149.230 avatars7.githubusercontent.com
54.85.253.220 avatars8.githubusercontent.com
34.229.193.93 favicons.githubusercontent.com
54.204.171.93 developer.apple.com
3.231.159.178 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-03-25 08:57:14

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder