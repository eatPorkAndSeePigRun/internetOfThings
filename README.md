# internetOfThings

#### 项目介绍
本项目针对现有物联网系统从终端到服务器再到控制软件的一套独立的硬软件系统开发周期长，难度大的缺点，研制出具有自主知识产权的基于云平台与web技术的智能物联网系统，并对其关键技术展开深入研究，主要包括智能硬件与移动物联网开放云平台进行通信的研究，控制界面的版面设计与实现研究，物联网信息交换的自适应数据交换协议的研究等等，从而为工业等领域提供解决方案，降低技术门槛和用户使用难度，有效促进现代物联网技术的发展与应用推广。

# 项目结构

- config            一些配置
- controller        控制层
- dao               数据库层
- doc               存放其他.md文件
- model             实体类层
- service           逻辑层
    - main.go           主入口
    - .gitignore        忽略git不必要提交的文件
    - LICENSE           开源许可为Apache License

# 接口定义
总体形式为
```
{
    code: 0     // 只有0为请求成功 其他值需要额外处理
    msg:  'ok'  // 描述信息
    data: {} // 具体data
}
```

#### 安装教程

1. xxxx
2. xxxx
3. xxxx

#### 使用说明

1. xxxx
2. xxxx
3. xxxx

#### 参与贡献
|作者|备注|
|-|-|
|黄成浩||
|林佳苗||

1. Fork 本项目
2. 新建 dev_xxx 分支
3. 提交代码
4. 新建 Pull Request

数据库1.0版，其中monit数据表需要修改
![数据库模型图](https://eatporkandseepigrun.github.io/model.png)