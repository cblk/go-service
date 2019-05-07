
### 日志模块使用
```text
前置模块依赖: env
```

#### 支持yaml配置文件(Support profile)
```yaml
app_id: test.env.config.0111
log_stdout: true
log_level: 2
log_dir: "/var/logs/"
log_max_file_num: 100
log_max_file_size: 1000000000
log_split: true
log_split_by: week
```

#### 使用方法:
```text
参考 log_test.go
```