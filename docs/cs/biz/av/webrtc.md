## webrtc

webrtc为了提高实时性和丢包恢复，限制使用B帧

WebRTC 在实时通信中限制使用 B 帧（bidirectional frames）主要是出于对**延迟（latency）**和**丢包恢复（packet loss recovery）**的严格要求。

---

### 1. 什么是 I 帧、P 帧、B 帧？

首先，理解这三种帧类型是关键：

* **I 帧（Intra-coded frame）**: 帧内编码，也叫关键帧。它包含一幅完整的图像信息，可以独立解码，不依赖于其他任何帧。I 帧文件大，但提供了最高的画质和独立性。
* **P 帧（Predicted frame）**: 预测编码帧。它只记录与前一帧（I 帧或 P 帧）的差异信息。P 帧文件较小，依赖于前面的帧才能正确解码。
* **B 帧（Bidirectional predicted frame）**: 双向预测编码帧。它同时记录与前一帧和后一帧的差异信息。B 帧的压缩率最高，文件最小，但它**依赖于未来帧的信息才能解码**。

---

### 2. 为什么 WebRTC 要剔除 B 帧？

WebRTC 是一项实时通信技术，其核心目标是**低延迟**和**高可靠性**。而 B 帧的特性恰恰与此目标相悖。

* **引入解码延迟**:
    要解码一个 B 帧，你需要等待它后面的 P 帧或 I 帧到达。例如，要显示帧 `B2`，你必须先接收并解码帧 `P3`。这种“先看后播”的机制会**显著增加端到端的延迟**。在视频会议或实时直播等场景中，即使是几十毫秒的延迟也会让用户感到卡顿和不连贯。

* **加剧丢包影响**:
    由于 B 帧的依赖关系复杂（既依赖过去帧也依赖未来帧），一旦某个关键帧（如 I 帧或 P 帧）丢失，**可能导致一连串 B 帧都无法正确解码**。例如，如果 `P3` 丢失，那么依赖 `P3` 的所有 B 帧都无法重建，这会造成屏幕长时间“冻结”或出现花屏，直到下一个 I 帧出现才能恢复。在不可靠的网络环境中，这会导致非常差的用户体验。

### 总结

在传统的视频点播或下载场景中，B 帧是提高压缩效率的绝佳工具，因为延迟不是主要问题，并且可以通过缓存和重传机制来解决丢包。

然而，对于 WebRTC 这种要求**超低延迟、对丢包敏感**的实时应用，为了保证流畅性和快速恢复能力，**牺牲 B 帧的压缩效率是必须的权衡**。因此，WebRTC 推荐并默认使用不含 B 帧的编解码器，如 VP8、VP9 和 H.264 的 High Profile，以确保最佳的实时通信体验。

#### 工具

```sh
# 1. 去除b帧
ffmpeg -i test.mp4 -vcodec libx264 -bf 0 test-640x480.h264

# 2. web aac转opus
-acodec opus
# 3.查看视频信息
ffprobe -v quiet -show_frames -select_streams v test.mp4 | grep "pict_type=B"

```

---
### WebRTC面试常见问题

WebRTC（Web Real-Time Communication）是一个支持浏览器进行实时通信的技术，主要用于音视频通话、文件共享等应用。以下是一些在WebRTC面试中可能被问到的常见问题，它们涵盖了从基础概念到高级实现的各个方面。

#### 1. 基础概念

* **什么是WebRTC？它解决了什么问题？**
    * 解释WebRTC的核心功能：在浏览器之间直接进行点对点（P2P）的实时音视频和数据传输。
    * 说明它解决了传统Web应用无法直接进行实时通信的难题，降低了开发成本和延迟。

* **WebRTC包含哪三个主要API？**
    * **`MediaStream` (getUserMedia):** 用于访问用户的摄像头、麦克风等设备，获取媒体流。
    * **`RTCPeerConnection`:** 负责建立和管理点对点连接，是WebRTC的核心。
    * **`RTCDataChannel`:** 用于在点对点连接上进行任意数据的传输。

* **SDP（Session Description Protocol）是什么？它在WebRTC中有什么作用？**
    * SDP是一种描述多媒体会话信息的协议。
    * 在WebRTC中，SDP用于交换通信双方的媒体能力（如支持的编解码器、分辨率、网络地址等），以便协商出一个双方都能接受的通信参数。

---

#### 2. 连接建立与信令

* **WebRTC为什么需要信令服务器？**
    * WebRTC本身不提供信令机制，它只是一个点对点通信的框架。
    * 信令服务器负责处理连接建立前的元数据交换，包括：
        * **发现和协商:** 帮助两个对等端（peer）找到彼此。
        * **会话控制:** 发送和接收SDP offers和answers。
        * **网络信息交换:** 传递ICE candidates。
    * 信令服务器可以是任何能进行双向通信的服务器，如WebSocket、Socket.IO等。

* **解释`offer/answer`模型和SDP交换过程。**
    * **Offer:** 发起方创建并发送一个SDP offer，描述自己的媒体能力。
    * **Answer:** 接收方收到offer后，创建一个SDP answer，描述自己接受的参数。
    * 双方通过信令服务器交换这些SDP信息，从而协商出最终的连接参数。

---

#### 3. 网络穿越与NAT

* **什么是NAT？为什么WebRTC需要解决NAT穿越问题？**
    * NAT（Network Address Translation）是一种网络技术，允许多台设备共享一个公网IP地址。
    * 由于NAT的存在，直接通过公网IP地址连接设备通常是不可能的，这使得P2P连接变得困难。

* **解释ICE、STUN和TURN的作用。**
    * **ICE (Interactive Connectivity Establishment):** 是一个框架，用于找出连接两台设备的最优路径。它综合使用STUN和TURN。
    * **STUN (Session Traversal Utilities for NAT):** 用于获取设备的公网IP地址和端口，帮助设备发现自己是否在NAT后面。
    * **TURN (Traversal Using Relays around NAT):** 当STUN无法建立直接连接时，TURN服务器充当**中继**（relay）角色，所有数据都通过它转发。这是一种确保连接成功但会增加延迟和服务器成本的方案。

---

#### 4. 媒体处理与编解码

* **什么是编解码器（Codec）？WebRTC常用的音视频编解码器有哪些？**
    * 编解码器是用于压缩和解压缩音视频数据的算法。
    * **音频:** Opus（强制支持）、G.711、iSAC等。Opus是WebRTC首选的音频编码器，因为它在低码率下表现出色。
    * **视频:** VP8、VP9、H.264、AV1。VP8和H.264是WebRTC早期强制支持的编码器。

* **解释Jitter Buffer、FEC、NACK等技术在WebRTC中的作用。**
    * **Jitter Buffer (抖动缓冲区):** 用于平滑网络抖动，减少音视频卡顿。
    * **FEC (Forward Error Correction):** 在发送端增加冗余数据，以在接收端恢复丢失的数据包，提高抗丢包能力。
    * **NACK (Negative Acknowledgement):** 接收方告知发送方某个数据包丢失，请求重传。

---

#### 5. 架构与实现

* **如何实现一个简单的WebRTC视频通话？**
    * 详细描述从**获取媒体流** -> **信令服务器交换SDP/ICE** -> **`RTCPeerConnection`连接建立** -> **音视频流传输**的全过程。

* **WebRTC中的SFU和MCU是什么？它们有什么区别？**
    * 这两种是多方会议的常见架构。
    * **SFU (Selective Forwarding Unit):** 服务器只转发收到的流，不进行混流。每个客户端发送一个流到SFU，然后从SFU接收其他所有参与者的流。
    * **MCU (Multipoint Control Unit):** 服务器接收所有参与者的流，将它们解码、混合成一个流，再编码后发送给每个客户端。

* **WebRTC的安全性如何保障？**
    * WebRTC强制使用**DTLS (Datagram Transport Layer Security)** 和 **SRTP (Secure Real-time Transport Protocol)** 对所有数据（包括信令和媒体流）进行加密，保证了通信的端到端安全。

---

### 面试建议

在面试中，除了回答这些技术问题，还可以展示你的**实际项目经验**。例如，你是否遇到过**WebRTC连接失败**的问题？你是如何**调试**的？你对**不同架构（P2P、SFU、MCU）**的优缺点有什么看法？这些问题能更好地体现你解决实际问题的能力和对技术的深入理解。

如果你对某个问题不太确定，可以告诉我，我很乐意为你提供更详细的解释。


在WebRTC中，点对点（P2P）连接的实现是一个复杂但巧妙的过程，它主要依赖于以下三个关键技术：信令（Signaling）、ICE框架（Interactive Connectivity Establishment）以及对等连接（RTCPeerConnection）API。

下面是WebRTC实现点对点通信的具体步骤和工作原理：

### 1. 信令（Signaling）
WebRTC本身不提供任何用于连接的“发现”机制。因此，在两个浏览器之间建立直接连接之前，它们需要通过一个**信令服务器**来交换元数据。这个过程就像是两个陌生人想打电话，他们需要先通过一个中间人（信令服务器）来交换彼此的电话号码。

信令服务器负责交换以下信息：
* **会话描述（SDP - Session Description Protocol）:**
    * **Offer (提议):** 一方（发起方）创建并发送一个SDP Offer，其中包含了自己希望进行的通信类型（如音视频、数据通道）、支持的编解码器、网络信息等。
    * **Answer (应答):** 另一方（接收方）收到Offer后，根据自己的能力和意愿，创建一个SDP Answer作为回应，描述自己接受的参数。
* **网络候选者（ICE Candidates）:**
    * 这些是每个对等端（Peer）可能用来连接的“网络地址”，包括本地IP地址、公网IP地址以及中继服务器地址。

信令服务器可以是任何能进行双向通信的服务器，比如使用WebSocket、Socket.IO或者其他HTTP长轮询技术。

### 2. ICE框架（Interactive Connectivity Establishment）
这是实现P2P连接的核心，特别是解决了网络地址转换（NAT）和防火墙的穿越问题。ICE不是一个独立的协议，而是一个寻找最佳连接路径的框架，它会尝试所有可能的连接方式，直到找到一个可行的路径。

ICE主要依赖以下两种服务器：
* **STUN (Session Traversal Utilities for NAT):** STUN服务器的作用是让位于NAT后面的设备发现自己的公网IP地址和端口。当浏览器向STUN服务器发送请求时，STUN服务器会返回浏览器请求的公网地址。这解决了最基本的NAT问题，即让对等端知道对方的“真实”公网地址。
* **TURN (Traversal Using Relays around NAT):** 如果STUN无法建立直接连接（通常是因为更严格的防火墙或对称NAT），ICE框架会退而求其次，使用TURN服务器。TURN服务器充当**中继器**，它接收一方的数据，然后转发给另一方。虽然这会增加延迟和服务器成本，但它确保了在任何网络条件下都能成功建立连接。

### 3. RTCPeerConnection API
`RTCPeerConnection` 是WebRTC中最重要的API，它负责管理整个P2P连接的生命周期。

**具体实现流程：**

1.  **创建 RTCPeerConnection:**
    * 两个对等端分别在自己的浏览器中创建 `RTCPeerConnection` 实例。

2.  **获取本地媒体流:**
    * 使用 `navigator.mediaDevices.getUserMedia()` 获取本地摄像头和麦克风的媒体流。
    * 通过 `addTrack()` 方法将这些流添加到 `RTCPeerConnection` 实例中。

3.  **发起方创建 Offer:**
    * 发起方调用 `createOffer()` 方法，WebRTC引擎会自动生成一个SDP Offer。
    * 发起方将这个Offer设置为本地描述 `setLocalDescription()`。
    * 同时，WebRTC引擎开始在后台收集ICE Candidates。每收集到一个新的候选者，就会触发一个事件。

4.  **信令服务器交换 Offer:**
    * 发起方通过信令服务器将SDP Offer发送给接收方。

5.  **接收方处理 Offer 并创建 Answer:**
    * 接收方收到Offer后，调用 `setRemoteDescription()` 将其设置为远程描述。
    * 然后，接收方调用 `createAnswer()`，WebRTC引擎会生成一个SDP Answer。
    * 接收方将Answer设置为本地描述 `setLocalDescription()`。
    * 接收方也开始收集自己的ICE Candidates。

6.  **交换 ICE Candidates:**
    * 在SDP交换的同时，两个对等端通过信令服务器持续地交换各自收集到的ICE Candidates。每收到一个对方的候选者，就调用 `addIceCandidate()` 将其添加到自己的 `RTCPeerConnection` 中。

7.  **ICE 竞争与连接建立:**
    * `RTCPeerConnection` 实例拿到所有本地和远程的ICE Candidates后，会开始进行“竞争”：它会尝试所有可能的组合（如本地IP到远程STUN地址，本地TURN地址到远程本地IP等），找到一个最佳的、可用的连接路径。
    * 一旦找到一个成功的路径，连接状态就会变为 `connected`，真正的点对点连接就此建立。

8.  **媒体流传输:**
    * 连接建立后，WebRTC会自动通过建立好的P2P路径传输音视频流。所有传输的数据都会使用DTLS和SRTP协议进行加密，保证了通信的安全性。

简而言之，P2P的实现是一个**“先握手，后找路，再传输”**的过程。信令服务器负责握手，ICE框架负责找路，`RTCPeerConnection` 负责执行这一切并将媒体流直接发送给对方。

