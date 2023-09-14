> * 原文地址：[The Open Source Project Nginx](http://aosabook.org/en/nginx.html)
> * 原文作者：[Andrew Alexeev](http://aosabook.org/en/intro2.html#alexeev-andrew)
> * 译文出自：[掘金翻译计划](https://github.com/xitu/gold-miner)
> * 本文永久链接：[https://github.com/xitu/gold-miner/blob/master/TODO1/the-open-source-project-nginx.md](https://github.com/xitu/gold-miner/blob/master/TODO1/the-open-source-project-nginx.md)
> * 译者：[razertory](https://github.com/razertory)
> * 校对者：[yqian1991](https://github.com/yqian1991)

# 开源项目之 Nginx

nginx（读作 "engine x"）是一位名叫 Igor Sysoev 的俄罗斯软件工程师开发的。自 2004 年发布以来，nginx 就一直专注于实现高性能，高并发和低内存占用。nginx 的额外功能，比如：负载均衡、缓存和流量控制以及高效集成在 Web 服务上的能力，使得它成为了当今网站架构的必选。如今，nginx 已经成为互联网中第二受欢迎的开源 Web 服务器。

## 14.1 高并发为何如此重要？

如今，互联网早已无处不在，我们已经很难想象十年前没有互联网的样子。现在的互联网发生了翻天覆地的变化，从基于 NSCA 的可以点击 HTML 页面和基于 Apache 的 Web 服务，到如今能够实现超过 20 亿人实时的沟通。随着 PC、手机和平板的的蔓延，互联网已经将全球经济数字化。面向信息和娱乐的在线服务变得更加优质。线上商业活动的安全方面也发生了明显变化。因此，网站也比以前更加的复杂并且需要大量的工程投入来确保鲁棒性和可扩展性。

并发性成为了网站架构设计的最大挑战之一。自从 web 服务开始的时候，并发性的等级就在持续上升。对于一个热门网站来说，支持几百甚至是几百万用户同时访问来说也不是什么稀罕事情。20 年前，产生并发的原因主要还是客户端的 ADSL 或者拨号（dial-up）连接。如今，并发的产生来源于手机端和以及新型的应用架构，这些架构主要可以支持长连接来提供新闻、信息流发布和朋友间的 feed 流等等。另一方面，导致高并发还由于现代浏览器的工作发生变化，通常是为了提高网页加载速度同时打开 4 到 6 个连接。

为了表述清楚缓慢这种问题，设想一下，一个基于 Apache 的，可以提供 100KB 大小带有文字或者图片的简单 web 服务器。生成或者重新产生这个网页只需要极少不到一秒的时间。但是在带宽只有 80kps 的情况下（下载速度 10kb/s），传输数据到客户端却会花掉 10s。本质上，服务器产生 100kb 数据的速度是相对较快的，随后在传输数据到客户端直至释放连接的过程却是相对较慢的。现在设想，你同时有 1,000 个独立的客户端连到你的服务器并且请求同样的内容。如果对于每个独立的连接，都会占用额外的 1MB 内存，那么对于 1,000 个连接来说就对导致多占用 1000 MB（1G）的内存，而这些仅仅是为了给 1000 个客户端提供 100kb 的内容。实际上，一个典型的 Apache 服务器通常会为了一个连接占用超过 1MB 的内存，遗憾的是几十 k 的带宽足够让手机之间高效通讯。尽管从某种程度而言，发送数据给客户端是慢的，提高操作系统内核的 socket 缓冲大小是可以的，这个不是一个通常的解决方法，并且会有不良影响。

在持久连接中，处理并发会做起来总比说起来有更多的问题，因为要在新建 HTTP 连接的时候避免延迟，让客户端保持连接并且确保对于每个连接服务端都能够保证有足够内存可供使用。

因此，为了能够处理因为用户量增长产生高并发由此带来的负载上升，网站的就必须基于通过一定数目的高效模块来架设。同时，在从获得客户端连接请求，到处理完请求期间，像硬件（CPU，memory，disk），网络容量以及数据存储也是同样重要的。因此，web 服务器需要能在同时请求数和每秒请求频率这两方面都拥有扩展性。

### Apache 不合适吗？

Apache，开始于 1900s，如今依旧统治着互联网。最初它的架构满足于当时的操作系统和硬件，同时也满足于当时的只有一个独立的物理机运行一个 Apache 服务器的互联网状态。在 2000 年始，一个独立的服务器难以满足增长起来的 Web 服务的情况越来越明显。尽管 Apache 提供了一个可靠的基金会用于未来发展，然而，它这种为了每个新连接复制自身的架构，已经不再适用于非线性的网站扩张。最终，Apache 成为了一个有着许多不同特性，第三方扩展，和一些普遍用于 web 应用开发的功能的 web 服务器。然而，没有什么东西是十全十美的，Apache 有者丰富功能的同时，对于每个连接产生的 CPU 和内存消耗使得它不能很好的扩展。

因此，当服务器的硬件、操作系统和网络条件成为了网站增长的瓶颈时，全世界的 web 工程师开始寻找一种更加高效的方法。大约十年前，一位名叫  Daniel Kegel 的杰出工程师宣称："是时候让 web 服务能够支持 10k 并发了。"同时他还预测了我们现在会叫互联网云服务。c10k 问题一产生，就引来了许许多多的解决方案用以优化实时的高并发。nginx 成为了其中最出色的解决方案之一。

为了解决 C10k 问题中的 10,000 个实时的连接，nginx 用了一种与众不同的架构，这种架构会更适合在同时处理大量的连接和一秒钟内完成多次请求环境中，问题规模的增长是非线性的。nginx 是事件驱动的（event-based，所以它不会用 Apache 的那种为每一个 web 请求都申请一个进程或者线程。结果便是，即使负载升高，内存和 CPU 都还是处于掌控之中。nginx 目前可以在一台普通的机器上，同时处理上万的并发。

nginx 的第一个版本主要是和 Apache 服务器一起部署，用来单独处理原本是 Apache 处理的 HTML, CSS, JavaScript 和图片这样的静态资源。在随后的迭代中，nginx 支持像 FastCGI, ，uswgi 或者 SCGI 协议集成到应用当中部署，并且可以利用像 memcached 这样的分布式缓存系统。同时像反向代理，负载均衡这样的特性也随之加上。这些额外的特点让 nginx 成为了构建可扩展性 web 服务的高效的基础组件的工具之一。

2012 年二月，Apache 2.4.x 分支发布。尽管，这个最新版本的 Apache 增加了多核处理器支持模块和用于提升可扩展性和并发的模块，然而它的性能，并发能力，以及资源利用能力与纯事件驱动的 web 服务器比，依旧难以望其项背。 很乐意看到新版的 Apache 服务器有着更好的可扩展性，尽管这样可以减少自身的瓶颈，然而像典型的 nginx-plus-Apache 配置依旧会被使用。

### 使用 nginx 会有更多的优势吗？

能够高性能地处理高并发一直是部署了 nginx 之后获得的最主要的好处。然而，还有一些更有趣的东西。

在过去几年中，网站架构就一直在拥抱解耦并从 web 服务器中拆分出一些基础组件。然而，那些原本存在于 LAMP-based 的网站中的基础组件，在 LEMP-based（E 代表着 Nginx 的读音） 的网站中，却能让 web 服务器成为基础组件以及用一种不同的方式去集成相同的或者改进了的应用和数据库工具。

nginx 非常适合做这个，因为它可以方便提供一个并发支持，延迟超时处理，SSL 支持，静态文件支持，压缩和缓存，甚至是 http 流媒体的高效的层级，而这些功能原本处于应用层。nginx 也可以直接集成一些像 Redis/memcached 这样的 NoSQL 用以优化大用户量场景。

当近代的开发语言和工具流行起来的时候，越来越多的公司正在改变他们的开发和部署方式。nginx 成为了改变过程中最重要的部分，同时，nginx 让很多公司在有限的预算中，快速地启动开发他们的服务。

nginx 是从 2002 年开始开发。到 2004 年，它以 two-clause BSD license 发布。随后，nginx 用户量开始增高，修改建议，bug 报告，观察报告等都在社区中不断完善 ngix。

nginx 最初的源码是用 C 完成的。nginx 已经可以部署在许多架构和操作系统中，比如 Linux, FreeBSD, Solaris, Mac OS X, AIX and Microsoft Windows。nginx 拥有自己的库并且并没有大量使用 C 标准库，一些像 zlib, PCRE and OpenSSL 这一类的库因为有证书冲突而没有被采用。

在 Windows 上部署 nginx 更像是一个实现 nginx 的理论证明而不是一个功能完善的项目。由于内核限制，nginx 的一些功能特性并不能发挥出来。在 windows 上的 nginx 并发能力、性能会更低，也没有缓存和带宽策略。将来 windows 上的 nginx 版本会继续完善。

## 14.2. nginx 架构总览

传统的解决并发的方式是每个单独的请求一个进程或者线程，并且网络和 io 操作都是阻塞式的。在传统的应用当中，这种做法会由于 CPU 和内存开销导致低效。开启一个独立的进程或者线程会需要预加载一个新的运行时环境和上下文。这些东西也会占用一些额外的 CPU 时间，线程频繁轮换导致的上下文切换带来的开销最终导致了低性能。这些问题在一些旧的 web 服务架构，比如 Apache 中得到了证实。这是在提供丰富普遍特性与优化服务器开销之前的一种权衡。

从最早开始，nginx 就被设定为在网站用户动态增长期间，用来提高网站性能和服务器资源利用率的工具，以至于它拥有一种与众不同的模型。这是受一些操作系统的事件驱动概念启发。这也产生了 nginx 的核心架构：模块化，事件驱动，异步，单线程，非阻塞。

nginx 大量采用多路复用（multiplex）和事件通知，并对每个 nginx 进程分配了特定的任务。连接被有限个数单线程的 worker 进程高效轮询（run-loop）处理。 每个 worker 都可以同时处理数千个并发连接和每秒请求。

### Code Structure 代码结构

worker 代码包含了核心和功能模块。nginx 核心负责维护一个紧凑的轮询，并在处理请求的每个阶段都执行模块中对应的部分。模块构成了大部分表示层和应用层功能。模块从网络和存储介质中进行数据的读写，传输内容，过滤出站内容，执行服务端的动作和当代理功能被打开的时候传递请求到被代理的（upstream）服务器。

nginx 模块化的架构可以让开发者在不修改核心代码的情况下加入一些自定也的扩展。nginx 模块稍微有点不同，比如核心模块、事件模块、阶段处理器、协议、变量处理器、filter，upstream 和负载均衡。目前，nginx 不再支持动态加载模块。模块在 nginx build 阶段就会被编译。然而，在将来 nginx 会在主版本上提供 loadable 模块和 ABI。更多关于不同模块的信息详见 [Section 14.4](#sec.nginx.internals).

在处理一些关于网络接收，处理和管理以及内容检索的时候，nginx 使用了事件通知（event notification）机制以及一些操作系统（ Linux, Solaris and BSD-based）的磁盘 IO 优化，比如：`kqueue`, `epoll`, and `event ports`。目的是为操作系统提供尽可能多的提示，以便为入站和出站流量、磁盘操作、socket 读写、超时等获取及时的异步反馈。针对 nginx 运行的每个 unix-like 的操作系统，对多路复用和高级 I/O 操作使用不同的方法进行了大量优化。

更多 nginx 架构高级概述详见 [Figure 14.1](#fig.nginx.arch).

![](http://aosabook.org/images/nginx/architecture.png)

Figure 14.1: Diagram of nginx's architecture

### Workers 的模型

正如之前提到的，nginx 并不为每个连接开一个进程或者线程。相反，worker 进程为每个新连接都采用一个共用的监听 socket 并在轮询中高效处理着数千个连接。对于 nginx 的 worker，没有采用一些特别的连接机制，都是由操作系统内核来完成的。一旦启动，一些监听 socket 就会被创建。worker 就会持续地接受连接，处理 http 请求和从对应的这些 socket 中读写数据。

轮询是 nginx 代码中最复杂的部分。它包括了综合（comprehensive）的内部调用和依赖大量的异步任务处理思想。异步操作通过模块化，事件通知，函数回调和计时器实现。总体上，关键在于尽可能的非阻塞。唯一让 nginx worker 阻塞的只有磁盘不足的情况。

因为 nginx 不会为每个连接新开进程或者线程，内存占用在很多场景下都不会高。nginx 节约了 cpu 占用也是因为没有进程线程的创建和销毁。nginx 要做的就是检查网络和存储，创建新连接，把新连接加入到轮询，并且在完成之前都异步处理。nginx 谨慎采用了一些系统调用比如资源池化和内存分配，以至于在极端的情况下也不会有很高的 CPU 占用。

由于 nginx 处理连接就开了几个 worker，在多核情况下可以很好的扩展。大致就是一个核心一个 worker，这样每个 worker 充分利用 cpu 核心，避免了线程切换和锁等待。不会产生资源不足并且每个单线程的 worker 进程中都存在资源管理策略。这种模型允许在不同存储设备之间有更好的扩展性，促进磁盘利用并且避免了磁盘 IO 阻塞。总的来说，服务器资源在多个 worker 工作的情况下被更高效使利用了。

对于某些磁盘使用和 CPU 负载模式，应该调整 nginx worker 的数量。这些规则在这里有点基础，系统管理员应该基于他们的工作负载尝试一些配置。一般建议如下：如果负载模式是 CPU 密集型的—例如，处理大量 TCP/IP、执行 SSL 或压缩，nginx worker 的数量应该与 CPU 核心的数量相匹配；如果负载主要是磁盘 I/O 限制。例如，从存储中提供不同的内容，或者大量的反向代理，workers 的数量可能是内核数量的 1.5 到 2 倍。有些工程师根据单个存储单元（磁盘分区）的数量来选择 workers 的数量，这种方法的效率取决于磁盘存储的类型和配置。

nginx 开发人员在即将发布的版本中要解决的一个主要问题是如何避免磁盘 I/O 上的大部分阻塞。目前，如果没有足够的存储性能来服务于由特定的 worker 生成的磁盘操作，那么 worker 仍然可能阻塞从磁盘读取 / 写入。存在许多机制和配置文件指令来减轻此类磁盘 I/O 阻塞场景。最值得注意的是，sendfile 和 AIO 等选项的组合通常会为磁盘性能带来很大的空间。应该根据数据存储、可用的内存大小和底层存储体系结构来计划 nginx 的安装。

现有 worker 模型的另一个问题是关于内嵌脚本支持的限制。首先，使用标准的 nginx 发行版，只支持嵌入 Perl 脚本。对此有一个简单的解释：关键问题是内嵌脚本可能阻止任何操作或意外退出。这两种类型的行为都会立即导致 worker 被挂起，同时影响数千个连接。需要更多的工作来让 nginx 的嵌入式脚本更简单、更可靠、适合更多的应用程序。

### nginx 进程角色

nginx 在内存中运行几个进程；有一个 master 进程和几个 worker 进程。还有一些特殊用途的进程，特别是缓存加载器和缓存管理器。版本 1.x 中的所有进程都是单线程的。所有进程主要使用共享内存机制进行进程间通信。主进程作为 root 用户运行。缓存加载器、缓存管理器和 worker 作为非特权用户运行。

master 进程主要有以下任务

*   读取并验证配置文件
*   创建、绑定和关闭 socket
*   启动，终止和维护配置好了个数的 worker 进程
*   不中断情况下重新加载配置
*   控制热更新（从二进制文件启动和必要情况下回滚）
*   打开日志文件
*   编译内嵌的 perl 脚本

worker 进程接受和处理来自客户机的连接，提供反向代理和过滤功能，并完成 nginx 能够做的几乎所有其他事情。关于监视 nginx 实例的状况，系统管理员应该关注 worker 进程，因为他们是反映 web 服务器实际日常操作的过程。

缓存加载器进程负责检查磁盘上的缓存项，并使用缓存元数据填充 nginx 的内存数据库。实际上，缓存加载器准备 nginx 实例来处理已经存储在磁盘上的文件，这些文件位于一个特别分配的目录结构中。它遍历目录，检查缓存内容元数据，更新共享内存中的相关条目，然后在一切都干净且可以使用时退出。

缓存管理器主要负责缓存过期和失效。在正常的 nginx 操作过程中，它保持在内存中，在失败的情况下由主进程重新启动。

### nginx 缓存简览

nginx 中的缓存是以文件系统上分层数据存储的形式实现的。缓存 key 是可配置的，可以使用不同的特定于请求的参数来控制进入缓存的内容。缓存 key 和缓存元数据存储在共享内存段中，缓存加载器、缓存管理器和 worker 进程可以访问共享内存段。目前，除了操作系统的虚拟文件系统机制产生的优化之外，没有任何文件的是缓存在内存当中。每个缓存的读取都放在文件系统上的不同文件中。层次结构（级别和命名细节）是通过 nginx 配置指令控制的。当将响应写入缓存目录结构时，路径和文件的名称来自代理 URL 的 MD5 值。

在缓存中放置内容的过程如下：当 nginx 从 upstream 服务器读取响应时，内容首先被写入缓存目录结构之外的临时文件中。当 nginx 完成对请求的处理后，它会重命名临时文件并将其移动到缓存目录中。如果用于代理的临时文件目录位于另一个文件系统上，则会复制该文件，因此建议将临时目录和缓存目录保存在同一个文件系统上。当需要显式清除缓存目录结构中的文件时，从缓存目录结构中删除文件也是相当安全的。nginx 有第三方的扩展，可以远程控制缓存的内容，并且计划了更多的工作来让这个功能可以集成到主发行版中。

## 14.3. nginx 配置

nginx 的配置系统受到了 Igor Sysoev 使用 Apache 的经验的启发。他的主要观点是，对于 web 服务器来说，可伸缩的配置系统是必不可少的。当使用大量虚拟服务器、目录、位置和数据集维护大型复杂配置时，会遇到扩展问题。在一个相对较大的 web 设置中，如果在应用程序和系统工程师都没有正确地完成，那么它可能是一个噩梦。

因此，nginx 配置的目的是简化日常操作，并提供进一步扩展 web 服务器配置的简单方法。

nginx 的配置保存在许多纯文本文件中，这些文件通常位于 `/usr/local/etc/nginx` 或`/etc/nginx`。主配置文件通常称为 `nginx.conf`。为了保持它的整洁，部分配置可以放在单独的文件中，这些文件可以自动包含在主文件中。然而，这里应该注意到 nginx 目前不支持 apache 风格的分布式配置（即”。htaccess 文件）。所有与 nginx web 服务器行为相关的配置都应该驻留在一组集中的配置文件中。

配置文件最初由 master 进程读取和验证。当 worker 进程从 master 进程 fork 时，worker 进程可以使用编译后的只读形式 nginx 配置。配置结构由通常的虚拟内存管理机制自动共享。

nginx 配置有几个不同的内容：`main`, `http`, `server`, `upstream`, `location` （同时 `mail` 相当于邮件服务代理）。配置文件内容不重叠。例如，在 main 中不存在 location。此外，为了避免不必要的歧义，没有任何类似于“全局 web 服务器”配置的东西。nginx 的配置是干净和合乎逻辑的，允许用户维护包含数千个指令的复杂配置文件。在一次私人谈话中，Sysoev 说，“全局服务器配置中的 location、directory 和其他块是我在 Apache 中不喜欢的特性，所以这就是为什么它们从未在 nginx 中实现的原因。”

配置文件语法、格式和定义遵循所谓的 c 风格约定。这种生成配置文件的特殊方法已经被各种开源和商业软件应用程序所使用。从设计上讲，c 风格的配置非常适合嵌套描述，具有逻辑性，易于创建、阅读和维护，并受到许多工程师的喜爱。nginx 的 c 风格配置也很容易自动化。

虽然 nginx 的一些指令类似于 Apache 配置的某些部分，但是设置一个 nginx 实例却是完全不同的体验。例如，nginx 支持重写规则，尽管需要管理员手动修改遗留的 Apache 重写配置以匹配 nginx 风格。重写引擎的实现也不同。

一般来说，nginx 设置还支持一些原始机制，作为精简 web 服务器配置的一部分非常有用。简单地提到变量和`try_files`指令是有意义的，这些指令对于 nginx 来说是唯一的。nginx 变量被开发出来是为了提供一个更强大的机制来控制 web 服务器的运行时配置。变量经过优化以快速解析，并在内部预编译为索引。根据需要进行解析，通常，变量的值只计算一次，并在特定请求的生命周期内缓存。变量可以与不同的配置指令一起使用，为描述条件请求处理行为提供了额外的灵活性。

“try_files”指令最初旨在以更合适的方式逐步替换条件“if”配置语句，它的设计目的是快速有效地尝试 / 匹配不同的 uri 到内容的映射。总的来说，`try_files` 指令工作得很好，并且非常高效和有用。更多详情推荐读者去 [`try_files` directive](http://nginx.org/en/docs/http/ngx_http_core_module.html#try_files)

## 14.4. nginx 内部

如前所述，nginx 代码库由核心和许多模块组成。 nginx 的核心是负责提供 Web 服务器，Web 和邮件反向代理功能的基础；它支持使用底层网络协议，构建必要的运行时环境，并确保不同模块之间的无缝交互。但是，大多数协议和特定的应用程都是由 nginx 功能模块完成的，而不是核心模块。

在内部，nginx 通过由模块组成的的管道或模块链来处理连接。换句话说，对于每个操作，都有一个正在进行相关工作的模块；例如，压缩，修改内容，执行服务器端，通过 FastCGI 或 uwsgi 协议与 upstream 应用服务器通信，或与 memcached 通信。

有几个 nginx 模块位于核心和真正的“功能”模块之间。这些模块是`http`和`mail`。这两个模块在核心和较低级别组件之间提供了额外的抽象级别。在这些模块中，实现了与诸如 HTTP，SMTP 或 IMAP 的相应应用层协议相关联的事件序列的处理。结合 nginx 核心，这些上层模块负责维护对各个功能模块的正确调用顺序。虽然 HTTP 协议目前是作为`http`模块的一部分实现的，但由于需要支持 SPDY 等其他协议，因此计划将来将其分离为功能模块。更多 SPDY 协议详见 [SPDY: An experimental protocol for a faster web](http://www.chromium.org/spdy/spdy-whitepaper)

功能模块可分为事件模块，阶段处理程序，输出 filter，变量处理程序，协议，上游和负载平衡器。大多数这些模块补充了 nginx 的 HTTP 功能，但事件模块和协议也用于`mail`。事件模块提供特定的 OS 依赖事件通知机制，如`kqueue`或`epoll`。 nginx 使用的事件模块取决于操作系统功能和构建配置。协议模块允许 nginx 通过 HTTPS，TLS / SSL，SMTP，POP3 和 IMAP 进行通信。

典型的 HTTP 请求处理周期如下所示。

1.  客户端发送 http 请求。
2.  nginx core 依据配置文件中的 location 选择合适的阶段处理器。
3.  如果配置生效，负载均衡器就会选择一个 upstream 服务器代理。
4.  阶段处理器执行任务，并把缓冲区的内容传递给第一个 filter。
5.  第一个 filter 将内容传递给第二个 filter
6.  第二个 filter 传递给第三个（迭代执行）
7.  将最后的 response 发送给客户端。

nginx 模块调用是非常可定制的。它使用指向可执行函数的指针来执行一系列回调。然而，这样做的缺点是它可能给想要编写自己的模块的程序员带来很大的负担，因为他们必须准确定义模块应该如何以及何时运行。 nginx API 和开发人员的文档都在不断改进，并且可以更多地用来缓解这个问题。

下面这些列子是可以添加模块的位置：

*   在读和处理配置文件之前
*   在每个服务器出现以及配置文件指向的地方
*   当 主配置 被初始化的时候
*   当服务器被初始化的时候
*   当 server configuration 被合并到 主配置的时候
*   当 location configuration 初始化或者合并到 parent server configuraton 的时候
*   当 master 进程启动或者存在的时候
*   当一个新的 worker 进程启动或者存在的时候
*   当处理一个请求的时候
*   当过滤请求 header 和请求 body 的时候
*   当请求转发到 upstream 服务器的时候
*   服务器中的响应的时候
*   当完成与一个 upstream 服务器的交互的时候

在 worker 进程中，导致生成响应的运行循环的 action 序列如下所示：

1.  启动 `ngx_worker_process_cycle()`.
2.  使用操作系统特定的机制来处理事件（such as `epoll` or `kqueue`）
3.  接收事件并且分发给相关的 action
4.  处理 / 代理请求 header 和 body
5.  产生响应内容 (header, body) 并传递给客户端
6.  结束请求
7.  重启 timers，events

轮询本身（步骤 5 和 6）确保增量生成响应并将其流式传输到客户端。

处理 HTTP 请求的更详细过程可能如下所示

1.  初始化请求处理
2.  处理 header
3.  处理 body
4.  调用相关的 nginx 处理器
5.  执行每个处理阶段

这将我们带到了每个阶段。当 nginx 处理 HTTP 请求时，它会将其传递给许多处理阶段。在每个阶段都有处理程序可以调用。通常，阶段处理程序处理请求并生成相关输出。阶段处理程序被附加到配置文件中定义的位置。

阶段处理程序通常执行以下四项操作：获取位置配置，生成适当的响应，发送 header 以及发送 body。处理程序有一个参数：描述请求的特定结构。请求结构有很多关于客户端请求的有用信息，例如请求 method，URI 和 header。

读取 HTTP 请求 header 时，nginx 会查找关联的虚拟服务器配置。如果找到虚拟服务器，请求将经历六个阶段：

1.  服务器重写阶段
2.  location 阶段
3.  location 重写阶段（将请求带回到上一个阶段）
4.  连接控制阶段
5.  try_files 阶段
6.  日志阶段

为了响应请求生成必要的内容，nginx 将请求传递给合适的内容处理程序。根据确切的位置配置，nginx 可能首先尝试所谓的无条件处理程序，如`perl`，`proxy_pass`，`flv`，`mp4`等。如果请求与上述任何内容处理程序都不匹配，则由以下处理程序之一按照以下顺序选择：`random index`，`index`，`autoindex`，`gzip_static`，`static`。

索引模块的详细信息可以在 nginx 文档中找到，但这些是使用尾部斜杠处理请求的模块。如果像`mp4`或`autoindex`这样的专用模块则不合适，内容被认为只是磁盘上的文件或目录（即静态），并由`static`内容处理程序提供服务。对于目录，它会自动重写 URI，以便始终存在尾部斜杠（然后发出 HTTP 重定向）。

然后将内容处理程序的内容传递给 filter。filter 也附加到 location，并且可以为 location 配置多个 filter。filter 执行操作处理程序生成的输出的任务。对于预先定义的开箱即用 filter，执行的顺序在编译时就确定。对于第三方 filter，可以在构建阶段对其进行配置。在现有的 nginx 实现中，filter 只能进行出站更改，并且目前没有机制来编写和附加 filter 来进行输入内容转换。输入过滤将出现在 nginx 的未来版本中。

filter 遵循特定的设计模式。调用 filter，开始工作，并调用下一个 filter，直到调用链中的最终 filter。之后，nginx 完成响应。filter 不必等待前一个 filter 完成。调用链中的下一个 filter 可以在上一个 filter 的输入可用时立即开始工作（功能上与 Unix 管道非常相似）。反过来，生成的输出响应可以在接收到来自上游服务器的整个响应之前传递给客户端。

还有 header filter 和 body filter；nginx 会分别用相关的 filter 来给相应 header 和 body 添加数据

header filter 主要有下面三个步骤

1.  决定是否对这个响应进行操作
2.  操作这个响应
3.  调用下一个 filter

body filter 修改生成的数据，下面是 body filter 的一些案例

*   服务端 includes
*   XSLT 过滤
*   图片过滤（比如修改图片尺寸）
*   修改编码
*   `gzip`压缩
*   chunked encoding

在 filter chain 之后，响应将传递给 writer。除了 writer 之外，还有一些额外特殊用途的 filter，即`copy`和`postpone`filter。 `copy`filter 负责使用可能存储在代理临时目录中的相关响应内容填充内存缓冲区。 `postpone`filter 用于子请求。

子请求是请求 / 响应处理的非常重要的机制。子请求也是 nginx 最强大的方面之一。对于子请求，nginx 可以从与客户端最初请求的 URL 不同的 URL 返回结果。一些 Web 框架将此称为内部重定向。但是，nginx 更进一步 - 过滤器不仅可以执行多个子请求，而且可以将输出组合成单个响应，但子请求也可以嵌套和分层。子请求可以执行其自己的子子请求，并且子子请求可以发起子子子请求。子请求可以映射到硬盘，其他处理程序或上游服务器上的文件。子请求对于根据原始响应中的数据插入其他内容非常有用。例如，SSI（服务器端包含）模块使用过滤器来解析返回文档的内容，然后将“include”指令替换为指定 URL 的内容。或者，它可以是一个过滤器，将文档的整个内容视为要检索的 URL，然后将新文档附加到 URL 本身

upstream 和负载均衡器也值得简要描述。upstream 用于实现可以被识别为反向代理（`proxy_pass`处理程序的内容。upstream 模块主要准备将请求发送到 upstream 服务器（或“后端”）并接收响应。这里没有调用输出 filter。当 upstream 服务器准备好被写入和读取时，upstream 模块确切地做的是设置要调用的回调。存在实现以下功能的回调：
*   创建的请求缓冲被发送到 upstream 服务器的
*   重新连接到  upstream 服务器（在请求产生之前）
*   处理 upstream 服务器响应的内容并且存储指向从 upstream 服务器内容的指针。
*   放弃请求（主要是客户端过早断开连接）
*   从 upstream 服务器读完内容之后结束请求
*   整理响应 body（比如删除 http 响应 trailer）

负载均衡器模块连接到`proxy_pass`处理程序，以便在多个 upstream 服务器符合条件时提供选择上游服务器的功能。负载均衡器注册启用配置文件指令，提供额外的上游初始化函数（以解析 DNS 中的上游名称等），初始化连接结构，决定在何处路由请求以及更新统计信息。目前，nginx 支持两种标准规则，用于对 upstream 服务器进行负载均衡：循环和 ip-hash。

upstream 和负载均衡处理机制包括用于检测失败的上游服务器以及将新请求重新路由到其余服务器的算法 - 尽管计划进行大量额外工作以增强此功能。总的来说，nginx 开发团队计划对负载均衡器进行更多的工作，并且在下一版本的 nginx 中，将大大改进跨不同上游服务器分配负载以及运行状况检查的机制。

还有一些其他有趣的模块提供了一组额外的变量供配置文件使用。虽然 nginx 中的变量是在不同的模块中创建和更新的，但有两个模块完全专用于变量：`geo`和`map`。 `geo`模块用于根据客户端的 IP 地址进行跟踪。此模块可以创建依赖于客户端 IP 地址的任意变量。另一个模块`map`允许从其他变量创建变量，实质上提供了对主机名和其他运行时变量进行灵活映射的能力。这种模块可以称为变量处理程序。

在单个 nginx worker 中实现的内存分配机制在某种程度上受到了 Apache 的启发。nginx 内存管理的高度概述如下：对于每个连接，必要的内存缓冲区被动态分配，链接，用于存储和操作请求和响应的头部和主体，然后在连接释放时释放。值得注意的是，nginx 试图尽可能避免在内存中复制数据，并且大多数数据都是通过指针值传递的，而不是通过调用`memcpy`。

更深入一点，当模块生成响应时，将检索到的内容放入内存缓冲区，然后将其添加到缓冲链链接中。后续处理也适用于此缓冲链链接。缓冲链在 nginx 中非常复杂，因为有几种处理方案因模块类型而异。例如，在实现 body filter 模块时精确管理缓冲区可能非常棘手。这样的模块一次只能在一个缓冲区（链路的链路）上运行，它必须决定是否覆盖输入缓冲区，用新分配的缓冲区替换缓冲区，或者在有问题的缓冲区之前或之后插入新的缓冲区。更复杂的是，有时模块会收到几个缓冲区，因此它必须有一个不完整的缓冲区链。但是，此时 nginx 只提供了一个用于操作缓冲区链的低级 API，因此在进行任何实际实现之前，第三方模块开发人员应该能够熟练使用 nginx 这个神秘的部分。

关于上述方法的注释是在连接的整个生命周期中分配了内存缓冲区，因此对于长期连接，保留了一些额外的内存。同时，在空闲的 keepalive 连接上，nginx 只花费 550 个字节的内存。对 nginx 的未来版本进行可能的优化将是重用和共享内存缓冲区以实现长期连接。

管理内存分配的任务由 nginx 池分配器完成。共享内存区域用于接受互斥锁，缓存元数据，SSL 会话缓存以及与带宽管制和管理（限制）相关的信息。在 nginx 中实现了一个 slab 分配器来管理共享内存分配。为了同时安全地使用共享内存，可以使用许多锁定机制（互斥锁和信号量）。为了组织复杂的数据结构，nginx 还提供了一个红黑树实现。红黑树用于将缓存元数据保存在共享内存中，跟踪非正则表达式位置定义以及其他几项任务。

遗憾的是，上述所有内容从未以一致和简单的方式描述，因此开发 nginx 的第三方扩展的工作非常复杂。虽然存在关于 nginx 内部的一些好的文档 - 例如，由 Evan Mille r 生成的那些文档 - 需要大量的逆向工程工作，并且 nginx 模块的实现仍然是许多人的黑科技。

尽管与第三方模块开发相关的某些困难，nginx 用户社区最近看到了许多有用的第三方模块。例如，有一个用于 nginx 的嵌入式 Lua 解释器模块，用于负载均衡的附加模块，完整的 WebDAV 支持，高级缓存控制以及本章作者鼓励并将在未来支持的其他有趣的第三方工作。（参考 Open Resty -- 译者注）

## 14.5. 收获

当 Igor Sysoev 开始编写 nginx 时，大多数给互联网赋能的软件已经存在，并且这种软件的体系结构通常遵循传统服务器和网络硬件，操作系统和旧的互联网体系结构。然而，这并没有阻止 Igor 认为他能够继续改进 Web 服务器领域的东西。所以，虽然第一课可能看起来很简单，但事实是：总有改进的余地。

考虑到更好的 Web 软件的想法，Igor 花了很多时间开发初始代码结构并研究为各种操作系统优化代码的不同方法。十年后，参考在版本 1 上的多年积极开发，他如今正在开发 nginx 版本 2.0 的原型。很明显，一个软件产品的新架构的初始原型和初始代码结构对于未来的重要性是非常重要的。

值得一提的另一点是发展应该集中。Windows 版本的 nginx 可能是一个很好的例子，说明如何避免在既不是开发人员的核心竞争力或目标应用程序的情况下稀释开发工作。它同样适用于重写引擎，该引擎在多次尝试增强 nginx 时出现，具有更多功能以便与现有的旧设置向后兼容。

但值得一提的是，尽管 nginx 开发者社区不是很大，但 nginx 的第三方模块和扩展一直是其受欢迎程度的重要组成部分。 Evan Miller，Piotr Sikora，Valery Kholodkov，Zhang Yichun（agentzh 中文名：章亦春）以及其他才华横溢的软件工程师所做的工作得到了 nginx 用户社区及其原始开发人员的赞赏。

* * *

This work is made available under the [Creative Commons Attribution 3.0 Unported](http://creativecommons.org/licenses/by/3.0/legalcode) license. Please see the [full description of the license](intro1.html#license) for details.

> 如果发现译文存在错误或其他需要改进的地方，欢迎到 [掘金翻译计划](https://github.com/xitu/gold-miner) 对译文进行修改并 PR，也可获得相应奖励积分。文章开头的 **本文永久链接** 即为本文在 GitHub 上的 MarkDown 链接。


---

> [掘金翻译计划](https://github.com/xitu/gold-miner) 是一个翻译优质互联网技术文章的社区，文章来源为 [掘金](https://juejin.im) 上的英文分享文章。内容覆盖 [Android](https://github.com/xitu/gold-miner#android)、[iOS](https://github.com/xitu/gold-miner#ios)、[前端](https://github.com/xitu/gold-miner#前端)、[后端](https://github.com/xitu/gold-miner#后端)、[区块链](https://github.com/xitu/gold-miner#区块链)、[产品](https://github.com/xitu/gold-miner#产品)、[设计](https://github.com/xitu/gold-miner#设计)、[人工智能](https://github.com/xitu/gold-miner#人工智能)等领域，想要查看更多优质译文请持续关注 [掘金翻译计划](https://github.com/xitu/gold-miner)、[官方微博](http://weibo.com/juejinfanyi)、[知乎专栏](https://zhuanlan.zhihu.com/juejinfanyi)。
