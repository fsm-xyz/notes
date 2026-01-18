# 计算优化


## 启用 CPU Burst

```sh
# ConfigMap ack-slo-config样例。
data:
  cpu-burst-config: |
    {
      "clusterStrategy": {
        "policy": "auto",
        "cpuBurstPercent": 1000,
        "cfsQuotaBurstPercent": 300,
        "sharePoolThresholdPercent": 50,
        "cfsQuotaBurstPeriodSeconds": -1
      }
    }

```
[具体参数](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/cpu-burst?spm=5176.8466032.console-base_help.dexternal.5a141450SyTaJD&scm=20140722.S_help%40%40%E6%96%87%E6%A1%A3%40%40371923.S_BB2%40bl%2BRQW%40ag0%2BBB1%40ag0%2Bos0.ID_371923-RL_%E5%90%AF%E7%94%A8cpuburst%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%E7%AD%96%E7%95%A5-LOC_console~UND~help-OR_ser-PAR1_215041d917681859732727523e33af-V_4-P0_0-P1_0#8e7235e5e24hf)

## CPU拓扑感知

[参考](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/topology-aware-cpu-scheduling?spm=a2c4g.11186623.help-menu-85222.d_2_14_5_1.1746e6d5JL3f0z)

## 资源感知

动态感知实时负载

[参考](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/use-load-aware-pod-scheduling?spm=a2c4g.11186623.help-menu-85222.d_2_14_4.745d7c10hh8Jvm)


## 超卖

[参考](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/dynamic-resource-overcommitment?spm=a2c4g.11186623.help-menu-85222.d_2_14_5_2.7235e6d5P3W4XW)

## 节点资源预留

[预留资源](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/resource-reservation-policy?scm=20140722.S_help%40%40%E6%96%87%E6%A1%A3%40%40330995._.ID_help%40%40%E6%96%87%E6%A1%A3%40%40330995-RL_%E9%A2%84%E7%95%99%E8%AE%A1%E7%AE%97%E8%B5%84%E6%BA%90-LOC_doc%7EUND%7Eab-OR_ser-PAR1_212a5d3e17681883502926995d826d-V_4-PAR3_r-RE_new5-P0_0-P1_0&spm=a2c4g.11186623.help-search.i20)
