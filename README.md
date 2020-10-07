# 由Golang实现的模式匹配库

做信息安全等领域开发的时候总会用到各种模式匹配，来进行单模/多模的特征串匹配。Golang虽然提供了庞大的类库来支持一般的串匹配，但总会遇到性能瓶颈，而且发现并没有专门的模式匹配类库来支持各种模式匹配。由此萌生了构建go lang模式匹配库的想法，来支持各类单模多模匹配。

目前支持情况：

1. 单模匹配：

   | 编号 | 匹配算法 | 特性                 |
   | ---- | -------- | -------------------- |
   | 1    | KMP      | 支持大小写不敏感匹配 |
   |      |          |                      |

   

2. 多模匹配：

   | 编号 | 匹配算法 | 特性 |
   | ---- | -------- | ---- |
   |      |          |      |
   |      |          |      |
