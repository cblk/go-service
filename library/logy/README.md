
#### 支持yaml配置文件(Support profile)
```yaml
log:
  app_id: "test.env.config.0111"
  format: "%L %e %D %T %a %f %S %M"
  stdout: true
  level: "info" # debug info notice warning error critical
  filter:
    - "A1"
    - "A2"
```

#### 使用方法:
```text
参考 log_test.go
```