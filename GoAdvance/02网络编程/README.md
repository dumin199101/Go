# 三次握手建立连接
- 主动发起请求端发送SYN标志位（请求建立连接）
- 被动接受请求端发送ACK标志位跟SYN标志位
- 主动发起请求端发送ACK标志位


# 四次挥手断开连接
- 主动关闭请求端发送FIN标志位（请求断开连接）
- 被动关闭请求端发送ACK标志位
- 被动关闭请求端发送FIN标志位
- 主动关闭请求端发送ACK标志位

# TCP状态转换
- 主动发起请求端: CLOSED--完成三次握手--ESTABLISHED--Dial函数返回
- 被动发起请求端： CLOSED--调用Accept函数--LISTEN--完成三次握手--ESTABLISHED--Accept函数返回
- 主动关闭请求端： ESTABLISHED--FIN_WAIT2(半关闭)--TIME_WAIT--2MSL(Maximum Segment Lifetime)--确认最后一个ACK被对端接收--CLOSE
- 被动关闭请求端： ESTABLISHED--CLOSE

 


