# cpds-analyzer
<p align="center">
<a href="https://gitee.com/openeuler/Cpds"><img src="docs/images/cpds-icon.png" alt="banner" width="250px"></a>
</p>

#### Description
cpds-analyzer is a container fault/sub-health diagnostic component developed for the CPDS (Container Problem Detect System) container fault detection system.

This component processes the abnormal data collected by cpds-detector (anomaly detection component) according to the diagnostic rules issued by cpds-dashboard (user interaction component), and judges whether the cluster node is in a container failure/sub-health state.


#### Build from source

`cpds-analyzer` is only supported on Linux and must be built with Go version 1.18 or higher.

```bash
# create a 'cpds' directory in your GOPATH/src
cd $GOPATH/cpds/
git clone https://gitee.com/openeuler/cpds-analyzer.git
cd cpds-analyzer

make
```
Finally, the compiled `cpds-analyzer` is in the `out` directory.


#### Contribution

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request
