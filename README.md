# goemail
go语言简单发送邮件，不支持附件，可用作发送验证码，通知等


# 获取方法
```
go get github.com/baikit/goemail
```

# 更新
```
go get -u github.com/baikit/goemail
```

# 使用方法

  ```
  emailconfig := EmailConfig{User:"",Password:"",Host:"",Port:""}
  
  to := "22@qq.com"
  title := "hello"
  msg := "hello"
  emailconfig.SendEmail(to,title,msg)
  
  ```
  
