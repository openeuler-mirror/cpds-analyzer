# cpds-analyzer

#### 介绍
cpds-analyzer是为CPDS(Container Problem Detect System)容器故障检测系统开发的容器故障/亚健康诊断组件。

本组件根据cpds-dashboard(用户交互组件)下发的诊断规则，对cpds-detector(异常检测组件)收集的异常数据进行处理，判断集群节点是否处于容器故障/亚健康状态。

#### 从源码编译

`cpds-analyzer`只支持 Linux，必须使用 Go 版本 1.18 或更高版本构建。

```bash
# create a 'gitee.com/cpds' in your GOPATH/src
cd $GOPATH/gitee.com/cpds
git clone https://gitee.com/openeuler/cpds-analyzer.git
cd cpds-analyzer

make
```

编译完成后的`cpds-analyzer`在`bin`目录中

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

