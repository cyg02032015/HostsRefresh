# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
54.144.78.116 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.90.115.58 raw.githubusercontent.com
54.90.115.58 camo.githubusercontent.com
34.229.230.78 avatars.githubusercontent.com
3.238.162.68 avatars0.githubusercontent.com
54.163.208.41 avatars1.githubusercontent.com
3.215.190.249 avatars2.githubusercontent.com
34.229.252.162 avatars3.githubusercontent.com
54.90.115.58 avatars4.githubusercontent.com
3.89.122.209 avatars5.githubusercontent.com
3.229.130.169 avatars6.githubusercontent.com
44.193.229.9 avatars7.githubusercontent.com
54.83.149.96 avatars8.githubusercontent.com
54.83.149.96 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-07-15 09:43:18

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder