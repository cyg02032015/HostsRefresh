# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.82.236.155 github.com
3.80.21.34 api.github.com
3.239.87.236 github.map.fastly.net
54.172.16.90 github.githubassets.com
3.88.24.173 github.global.ssl.fastly.net
52.90.53.56 raw.githubusercontent.com
54.173.167.37 camo.githubusercontent.com
3.89.24.37 avatars.githubusercontent.com
44.192.24.17 avatars0.githubusercontent.com
3.238.22.10 avatars1.githubusercontent.com
44.203.234.155 avatars2.githubusercontent.com
44.193.195.200 avatars3.githubusercontent.com
52.90.53.56 avatars4.githubusercontent.com
18.206.188.15 avatars5.githubusercontent.com
18.206.188.15 avatars6.githubusercontent.com
35.168.59.218 avatars7.githubusercontent.com
34.228.153.244 avatars8.githubusercontent.com
3.237.8.79 favicons.githubusercontent.com
54.198.163.91 developer.apple.com
50.17.53.64 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-05-06 09:04:58

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder