# Ingress & Ingress Controller範例(Part2)

## 目的

在現實的環境中, 不免會用到一些原本就內嵌了憑證且使用Https(443)來開放存取介面的應用服務。當我們把這樣的服務跑在Kubernetes的集群中時, 又想要讓Kubernets外的使用者可以以一般的Http(80)來使用時, 如何讓Ingress controller可以順利地把服務開放出去就成為一個課題。

而這個範例就是讓大家可以完成這個任務的手把手練習。 同時也是讓大家練習如何使用OpenSSL來產生Self-signed憑證的基本步驟。

## 範例

通過app2-ing將內部服務app2-svc的https(443)服務端口暴露到集群外，通過訪問`http://app2.example.com`即可訪問app2服務

Slideshare: https://www.slideshare.net/erhwenkuo/cncf-k8s-ingress-example02/

Youtube: https://youtu.be/0Gst9bUUTi4/


