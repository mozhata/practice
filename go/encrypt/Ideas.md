# 理发店收银系统
管理员(所有权利)
管理员创建工作人员账号
工作人员录入客户信息
具体情况: 收银人员负责客户信息录入, 收银, 记录收银信息, 记录本次服务人员
## 客户表
由店内工作人员创建,维护
{
    ID(str)
    name
    性别
    地址
    电话/邮箱
    备注
}
## 客户消费表
记录客户充值/消费情况
{
    顾客ID
    操作人员ID
    服务人员ID(理发师等)选填
    消费类型
    操作类型(充值,消费)
    操作时间
    金额
}