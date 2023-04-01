
# Parking Charge Adapter Service
  
+ 提供Web後台與車柱控制器間的請求轉傳中介點.

## HttpService
+ 接收Web後台的Request後push至channel.

## UnixSocketService
+ 監聽channel事件,若事件觸發則把事件訊息push至車柱控制器.
