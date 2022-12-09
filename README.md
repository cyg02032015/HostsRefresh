# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
3.87.182.156 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.82.248.209 raw.githubusercontent.com
3.82.157.139 camo.githubusercontent.com
34.238.249.153 avatars.githubusercontent.com
52.70.41.0 avatars0.githubusercontent.com
44.202.64.171 avatars1.githubusercontent.com
18.207.203.101 avatars2.githubusercontent.com
3.88.223.84 avatars3.githubusercontent.com
44.203.132.174 avatars4.githubusercontent.com
3.84.191.165 avatars5.githubusercontent.com
3.83.163.205 avatars6.githubusercontent.com
44.203.132.174 avatars7.githubusercontent.com
54.158.178.70 avatars8.githubusercontent.com
44.212.71.202 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-12-09 09:05:03

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder