# Note

`plantuml` 不是代码中使用的包，是在分析项目中需要用的工具。工欲善其事，必先利其器，好的工具能显著提高项目分析效率，当然工具是为人所用，不是人为工具服务。

uml 和 包图在项目分析中的重要性不言而喻，要想深入了解项目。UML 类图，包图，时序图等是要掌握，熟悉的必备技能。

## UML 类图

uml 类图使用 `PlantUML` 查看，在查看之前首先 `PlantUML` 会解析自定义语法 `puml` 生成 UML 类图。

通过 `goplantuml`[https://github.com/jfeliu007/goplantuml] 生成项目的 `puml`。示例如下:
```
goplantuml -recursive $GOPATH/src/github.com/marmotedu/iam > iam_uml.puml
```

接着 visual stdio 安装 `PlantUML` 插件，通过使用插件打开生成的 puml 文件 `iam_uml.puml`。更多信息可参考 [plantuml starting](https://plantuml.com/starting)。

对于 github 等项目也可以通过网页工具 [dumels](https://www.dumels.com/) 在线生成 UML 类图。


## go mod 图

mod 图解决的是可视化模块依赖关系，它不是包图。

mod 图可以通过 [gmchart](https://github.com/PaulXu-cn/go-mod-graph-chart) 生成。

## 小结
图是为人服务的。后续要进一步分析的是如何理解 UML 类图和包图，如何画 UML 类图和包图，最终目的是理解项目，乃至自造项目。
