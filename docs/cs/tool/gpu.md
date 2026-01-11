# GPU

使用GPU做计算的话，需要装好对应的驱动

## 调用

cuda, opencl, openacc, wwbGpu, Vulkan Compute

## Nvidia

在Ubuntu上安装cuda, 如果使用的OpenCL进行计算的话还需要安装ocl对应的头文件信息

```sh
apt install nvidia-cuda-toolkit -y

apt-get install ocl-icd-opencl-dev
```

## 共享的GPU

参考[知乎](https://www.zhihu.com/question/488792295)这里的几个厂商

除了云厂商的云服务器，云桌面等，还有一些专门做算力租赁的公司

+ [9gpu](https://9gpu.com)
+ [gpufree](https://www.gpufree.cn/)
