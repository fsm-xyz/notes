# k8s

## SLB配置

[SLB注解](https://www.alibabacloud.com/help/zh/doc-detail/86531.htm)

```sh
    # 证书
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-cert-id: 1857464601594004_17065841c58_-605729181_1764429477
    # SLB 类型
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-spec: slb.s1.small
    # 内网SLB
    service.beta.kubernetes.io/alicloud-loadbalancer-address-type: intranet
    # 使用指定的SLB
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-id: "${YOUR_LOADBALACER_ID}"
    # 指定SLB的名字
    service.beta.kubernetes.io/alibaba-cloud-loadbalancer-name: 123
```