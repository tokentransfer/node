# node

#### 介绍

1. 采用libp2p(Liquid)p2p网络底层
2. 具备WASM智能合约
3. 具备DAG数据存储功能
4. 具备前端页面包部署与显示功能

#### 软件架构

1. 最低功耗与成本运行支持完整应用的区块链网络服务
2. 具备大规模且完整的区块链帐本及数据同步与复制功能
3. 力图具备更高性能及完整功能的区块链智能合约
4. 尽力改进并完善满足要求的区块链共识算法

#### 安装教程

1. git clone 本仓库
2. vim(nano) go.mod
3. go mod tidy
4. go build
5. ./node help

#### 使用说明

1.  ./node genesis, 生成区块链genesis block
2.  ./node start, 运行区块链节点服务

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request