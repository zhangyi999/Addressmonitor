# Addressmonitor
ETH 地址监控产品需求

## 产品说明

ETH 地址监控：通过嗅探并监控交易所或其它区块链应用的主钱包，对项目的安全性评级，帮助用户在平台暴雷前规避资产风险。

### 功能

1. 地址分类：通过 一定时间内 指定币种的 转入笔数、转入金额、转出笔数、转出金额、转入地址来源 与 转出地址 来源几个参数对地址进行分类。

2. 地址标记：将一些特殊地址标记，特殊地址包括：交易所主钱包地址、各种区块链合约地址、涉案地址、各知名人士持有地址等。

3. 地址地图：将地址间的关联以拓扑图形式展示，在图上标注出重要地址。

案列说明：

1. zzex USDT ERC20 充币地址：0x3d475e9edef129acaacfb1cf282b842b723772f0

2. 向该地址充值后，该地址 USDT 被转入 0x86e793e413f519b450315fcc4b618eb25a3a54a4
![归集](./img/charg.png)

3. 追踪 0x86e793e413f519b450315fcc4b618eb25a3a54a4 地址，发现给地址为交易所归集地址。
- 该地址于 2019.7.24 日充入 19997 个 USDT
- 至 2019.8.24 该地址 U 为净支出，推测此前充提记录为测试阶段
- 导出  0x86e793e413f519b450315fcc4b618eb25a3a54a4 USDT 交易记录
![USDT 趋势](./img/zzex_usdt.png)
有 8次 大额提现

- 第二次提现最多
![大额提现](./img/zzex_usdt_1008.png)

查看该笔交易，挖到 0x19b7ec5ed41fb4c89d6f7b2952c5f70c7874e91a 地址
![大额提现](./img/zzex_usdt_1008.1.png)

- 查看 0x19b7ec5ed41fb4c89d6f7b2952c5f70c7874e91a 地址记录，确认该地址为 火币网用户充币地址。
![大额提现至火币](./img/zzex_usdt_1008.2.png)

- 2019.11.26 再次提币 30 w U 到 0x19b7ec5ed41fb4c89d6f7b2952c5f70c7874e91a 出售。
![大额提现至火币](./img/zzex_usdt_1008.3.png)

- 值得注意的是，2019.12.24 日这一笔提币并没有提到火币，极有可能是提到新交易所。
![大额提现至火币](./img/zzex_usdt_1008.4.png)

4. 生成拓扑关系
![拓扑示意图](./img/zzex_usdt_1008.5.png)

5. 风控策略
![拓扑示意图](./img/zzex_usdt_1008.6.png)


## 核心功能

### 地址标记

地址标记目前通过 [etherscan.io](https://etherscan.io/) 可以查到大部分合约和交易所，不过很多野鸡交易所还没有备案到浏览器中，需要通过【地址嗅探】【用户备案】等手段补全数据。

### 地址分类

通过地址分类找出特殊地址，能发未标记的重要地址，地址分类的标准方案如下：

#### 维度

单个 ETH 地址基本信息 `AddressDetail`:

```go

Transaction struct {
    from byte
    to byte
    value float64
    txid byte
    time uint  
}

AddressDetail struct {
	into []Transactions
	out []Transactions
	balance float64
	token byte
}

```

#### 路径

![路径](./img/zzex_usdt_1008.7.png)

通过筛选大额进出地址，追踪资金流向。

```go

Input struct {
    value float64
    address byte
    perc float64
}

Output struct {
    value float64
    address byte
    perc float64
}

AddressPath struct {
    input []Input
    output []Output
}

```

#### 路径追溯

![路径](./img/zzex_usdt_1008.8.png)

## 产品计划

### 0.0 Star

#### 分析流程

0. 地址嗅探

|交易所|币种|嗅探地址|提币费|提币限制|嗅探提币hash
|--|--|--|--|--|--|
|[CITEX](https://www.citex.co.kr/)|USDT|[0x5366f1554a5a92afbf6fe3bdd2f5af51cdd783a2](https://etherscan.io/token/0xdac17f958d2ee523a2206206994597c13d831ec7?a=0x5366f1554a5a92afbf6fe3bdd2f5af51cdd783a2)|2|(50, --)|[to:0xbf9265cc324e6e2ae8624fb584b1636d3bc00ff3<br>from:0xe8bdf6edf278271faa520e4b562e972c89d91c31](https://etherscan.io/tx/0x7cf2b8568e6806c588e033239ddaf33ff1f864b386e06f6dd282c081bc07d465)
|[ZZEX](https://www.zzexvip.com/)|USDT|0x3d475e9edef129acaacfb1cf282b842b723772f0|2|(--, --)|
|[BIKI](https://www.biki.com/)|USDT|[0x285bc7c15bc1b9f63ed89b46ce3b475a5c9075fb](https://etherscan.io/token/0xdac17f958d2ee523a2206206994597c13d831ec7?a=0x285bc7c15bc1b9f63ed89b46ce3b475a5c9075fb)|3|(--, 1500)|[TX](https://etherscan.io/tx/0xd9954123d38d9a00a31d4db13db1d5257030c650d7c1220ddcc4f21e1e39f5e8)
|[ZG](https://zg.com/)|USDT|[0xfdb86c48aca5f4168c596f7337803e0e6d6fcaf6](https://etherscan.io/token/0xdac17f958d2ee523a2206206994597c13d831ec7?a=0xfdb86c48aca5f4168c596f7337803e0e6d6fcaf6)|2|(7, 150002)|[TX](https://etherscan.io/tx/0xfb4a2f813b135ec42df09a27f499da3b609ff1b1054d4837dcd57a6a28fff1ed)

1. 给定目标地址
2. 检索特殊关系地址
3. 检索特殊关系地址的特殊关系地址
4. 遇到标记地址停止检索
5. 遇到中转地址标注并停止

##### 特殊地址方法

特征维度：

1. 转入数量: intoValue
2. 转入笔数: intoNum
3. 转出数量: outoValue
4. 转出笔数: outoNum
5. 转入地址数量: intoAddressNum
6. 转出地址数量: outoAddressNum
7. 持有 token 数量: balance

特殊地址分类：

|编号|分类|特征|疑似情况|
|--|--|--|--|
|1|小额用户地址|转入、转出频次低，地址余额不为0，币种数量 10 种以内|多为正常用户钱包地址|
|2|中转地址|转入后立即转出或到一定额度才转出，转出地址单一|多为交易所用户端充币地址|
|3|归集地址|频繁转入转出，地址币种数量较多> 10更明显，充入地址多为【2】类型地址|多为交易所钱包地址|
|4|活跃合约项目地址|频繁转入转出，地址币种单一|多为某些项目的智能合约地址|

将数值降序排列，按比例取 `index` 片段。

#### 模板要素

1. 监控币种
2. 监控地址
3. 当前余额
4. 关系图
    * 地址
    * 方向
    * 数量
    * 标记
5. 分析报告

#### 标记分类

1. 知名大地址：已被各大网站标注
2. 中转地址：大资金高频率每日进出
3. 普通地址：小资金低频率

### 0.1 Beta

目标：可快速分析地址并生成报告

时间：15 天

功能：【地址交易记录查询】【特殊地址标记】【关系图】【特殊地址提交】



