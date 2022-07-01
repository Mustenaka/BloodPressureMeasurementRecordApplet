# 各种项目规范

撰写人：Andrew

时间：2022.06.30

## Git 提交规范

每一次代码提交commit，需要按照

> [type] $(scope): subject
>
> [描述类型] $(影响范围): 描述 

如：

- [fix] $(fileName): add a function to fix bug.
- [add] $(fileArea): add a package for calculate the sum.
- [修改] 修复了用户头像展示错误的bug

中英文皆可

#### Type（必须）：commit的类别，允许使用如下表示

- feat: 新功能（可以add/添加）
- update: 更新（修改）
- fix: 修复bug
- docs: 文档改变
- style:代码格式改变
- refactor: 重构
- perf: 性能优化
- test: 增加测试
- build: 修改了build工具，或者build相关文件
- revert: 撤销commit

##### Scope（可选）：commit的影响范围

用于说明commit影响的范围，如数据层、控制层、视图层、或者某一个单独的文件，视项目不同而不同，注意我司中不常特殊描述，尽量采取单一commit。

> 即：一次提交尽量只说明一个内容的修改，在代码架构、文件布局中尽量做到解耦，这样一次提交就只需要影响极少数的代码文件，故scope可以省略。

##### subject（必须）：commit的简短描述，一般不超过50字符

无论中英文都尽量言简意赅。

英文提交规范：

1. use imperative(命令式), present tense: “change” not “changed” nor “changes”
2. don't capitalize first letter
3. no dot (.) at the end

##### 备注

不要嫌麻烦和强迫症，如果一个模块commit之后发现一个描述或者简单的几个字符需要重写，那么不要吝啬commit，再commit -m "[修改]修改了一个log描述"

## 更新规范

每一个项目都需要确立版本，以符号v开头，类型为 **v[大版本].[小版本].[功能逻辑优化版本]** 如v1.2.3

每一次更新，需要再项目中部署存放 **docs/updateDocument.md**（名字可以任取，中英亦可），对每一个更新进行记录，每一个功能逻辑优化版本更新都需要再主线merage到dev或者是main版本中，如果需要，可以在每一个小版本更新时导出Pakcage，并放置于Git仓库release。

例：

> ### v0.1.7 | 2022/06/30 - 2022/07/01
> - [新增] index页面现在尾部会显示app版本号 

## 文件布局

不做强制安排，以该项目的语言为准。

但是需要保证每一个项目新建都要保存README.md以及doc/更新记录.md两个文件，不开源选择私有，开源选择MIT协议，并且需要按照标注规则设置gitignore文件。



