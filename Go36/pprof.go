/*
48-49. 程序性能分析基础
在 Go 语言中，用于分析程序性能的概要文件有三种，分别是：CPU 概要文件（CPU Profile）、内存概要文件（Mem Profile）和阻塞概要文件（Block Profile）。
让程序对CPU概要信息进行采样：
	这需要用到runtime/pprof包中的 API。更具体地说，
	在我们想让程序开始对 CPU 概要信息进行采样的时候，需要调用这个代码包中的StartCPUProfile函数，
	而在停止采样的时候则需要调用该包中的StopCPUProfile函数。
*/