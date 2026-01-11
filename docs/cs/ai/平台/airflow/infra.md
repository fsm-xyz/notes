# 架构历程

自己搭建AIrflow3.0, 不直接使用已有的Airflow2.10的，选用自己觉得更好的KubernetesExector,不使用现在用的CeleryExecutor，用helm安装，推陈出新，全都折腾一遍，各个组件配置全部验证

初期想着直接在airflow的Operator里面直接调用监控数据，获取数据计算requests, usage，然后决定是否提交任务

想着当初面试一些公司的时候看到岗位要求，里面有二开k8s调度器的，就觉得这东西肯定有能根据底层资源进行调度的，虽然Airflow没有做到资源感知，很多AI没有给我好的结果，就尝试用Genmini 2.5Pro搜了下，找到类似的开源项目

AI不推荐自己写调度器，推荐了华为的Volcano, 一开始觉得很不错哦，能支持AI,ML,大数据这么多的产品介入，瞬间如获至宝，就直接基于最基本的Pod修改调度器，然后进行调度，但是发现一个问题就是一提交就产生一个Worker Pod监控Task, 就觉得不合理,假如我一直创建，会有无数个处于Pending的Pod在集群里面，还占用无数个Worker Pod，如果资源不够用，会一直创建占用资源的Worker Pod, 更加雪上加霜

就想着有没有更好的方案，然后出摸索出了Deferer+Trigger的方案，就是异步回调的方案，觉得更加优雅，我task运行的时候 worker再起来，很完美的方案  (解决worker太多)

封装创建vcjob的VolcanoOperator,觉得这样就可以控制控制pending的了，有个问题就是这个需要自己完全处理task的状态，尤其是和UI上结合的时候，要支持重试，删除，标记失败等功能，折腾了很久但是效果不好，修了这个，另一个又有点问题(解决pending太多)


最关键的一点就是worker会创建2次，进入defer后会删除，然后后面触发的时候worker起来，中途defer的时候删除，容易删除不掉，导致页面上杀掉了，后面实际task还在运行，状态乱了

找到cleanup+configmap清除方案，还是有时容易删不干净 (解决清除不干净)

最后返璞归真，直接在pod上修改schedulerName为volcano, pending就pending吧，也影响不大，worker的数量可以通过Celery架构的keda进行限制

周末研究的时候发现Volcano开源版本文档没有描述负载感知，但是我在网上看到过这个，于是看设计文档有，然后看到了公有云内置，就是没放到开源，于是萌生替换volcano为其他的

找了好几家类似的有腾讯的，但是很久不维护了，ali的koordinator,Nvidia的，k8s sigs的，总结下来就是ali的更加符合当前需求， 于是参照功能，觉得目前volcano有的，我们需要的功能，其实koordinator都差不多能做了，于是选型这个，还有最重要的一点就是文档齐全，联系方便，华为的周边项目感觉更像是KPI


Airflow后面增加了基于事件回调的模式，避免Worker Process常驻，浪费资源，尤其是在kubernetes环境中，但是Trigger + Deferered模式在实践中没有多少人用，自己写了一个基于BashOperator的，实现Volcano + Trigger + Deferered这样的一个在kubernetes中运行，但是跟Airflow UI结合的不好，所以还是回归到原生的UI + 当前的成熟方案CeleryExecutor，KubernetesExecutor，CeleryKubernetesExecutor)，理想很丰满，现实很骨干，等官方UI和Trigger这一套成熟稳定的时候，再使用，现在还回归主流吧

最终架构

Airflow3.0 

CelerykubernetesExecutor  (业务可以自己选择使用那个执行器)

理想情况是业务执行的镜像包含airflow worker的内容，或者sidecar模式，这样worker和task一起创建一起销毁，实现有多大锅下多大米，不用考虑worker节点的数量，worker节点合并到业务负载上

KubernetesExecutor + BashOperator  (最喜欢)
CeleryExecutor + KubernetesPodOperator(能接受)
KubernetesExecutor + KubernetesPodOperator  (太重，不推荐，虽然很隔离)


我总结了一套阿里云上的airflow比较好的方案
开启负载感知，避免资源高负载
开启超卖，最大化利用资源
使用ack-kube-queue，其实就是k8s内置的的Queue
elasticquotatree做资源配置，这个有个缺陷就是只支持多个namespace的，不像volcano可以是一个namespace下多个任务队列
这样的话，worker的数量应该就少了，就不一定需要镜像融合了
需要使用kubernetesjoboperator

跟volcano的对比，功能差不多，就是volcano是一个大而全的功能，阿里的就是组装了多个组件，需要开启多个组件

volcano的队列功能比较好，
可以直接设置队列权重

队列的百分比资源

直接使用pod加注解，会出现worker很多，task pod pending，使用镜像融合应该就可以解决

使用它的vcjob需要，自己管理生命周期