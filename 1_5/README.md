# Ingress & Ingress Controller範例(Part3)

## 目的

在一般企業環境內, 在注重網路安全的年代裡, 通常都會要求Web應用程式必須使用Https(443)的協議來確保Web client對Web server之間的網路封包不會因為沒有加密而外洩了重要的資訊。

然而現實的環境裡, 仍然有很多legacy的Web應用程式不容易(或不想)再修改來引人SSL的機制。這個時候, 最簡易的方法就是通過反向代理的伺服器在前端Web應用程式來進行SSL的握手協議, 然後再使用Http(80)的明文協議轉拋呼叫給後面Web應用程式。

雖然並不鼓勵這樣的做法, 但卻也是用最少的代價來啟動Https(443)的過渡方針。

## 範例

範例(3):是對反向代理SSL通信模型的”復現“，即client與ingress controller(nginx)間採用https加密通信，但ingress controller(nginx)與app3間則是明文的http通信。

為了讓ingress controller能夠進行SSL的通信握手, 必須要把SSL憑證上傳進Kubernetes裡, 然後設定Ingress來使用。

