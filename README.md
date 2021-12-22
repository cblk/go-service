# Bitbucket 私有仓库使用说明
1. 访问https://git.dev.yuanben.org/
2. 点击右上角头像 --> Manage account
3. 点击左侧选项：Personal access tokens --> Create new token
4. 输入Token Name
5. Permission 选择 Admin
6. 点击Create token，得到Token
7. 在命令行中执行
    ```shell
   go env -w GOPRIVATE=git.dev.yuanben.org
   # 使用Bitbucket的用户名和第5步得到的Token替换$Token和$User
   git config --global url."https://$User:$Token@git.dev.yuanben.org".insteadOf "https://git.dev.yuanben.org"
   go mod tidy
    ```