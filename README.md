# 前言

站在大佬的肩膀上继续前行。

由于一些原因，现在大部分金融站点、APP都进行了加密，正好忙里偷闲改个轮子，方便大家使用。

在此感谢两位大佬

https://github.com/anfutest

[https://github.com/G-Security-Team](https://github.com/G-Security-Team/JS-Forward)

本项目是基于两位大佬的进行二开，具体原理请移步

https://github.com/G-Security-Team/JS-Forward/

# 正文

原理图：

![image-20231115021545102](README.assets/image-20231115021545102.png)

使用介绍：

![image-20231115023949337](README.assets/image-20231115023949337.png)

将生成的payload直接插入到要篡改的变量前。

解密前：

![image-20231115024103249](README.assets/image-20231115024103249.png)

解密后：

可实现对加密包的全自动化加解密，更方便各位师傅转发到X-RAY、sqlmap等工具上。



![image-20231115024254201](README.assets/image-20231115024254201.png)

传输过程：

![image-20231115024400975](README.assets/image-20231115024400975.png)

注：https站点需要信任证书

![image-20231115024619447](README.assets/image-20231115024619447.png)
前置知识：https://mp.weixin.qq.com/s/1fisk5MumDcCLmSFJWaRpA

个人推荐练习靶场：

yakit Vulinbox
<img width="419" alt="image-20231115030836279" src="https://github.com/viplation/Js-forward3.0/assets/56590372/f4e20ccf-2b15-4eb8-804e-9006145b354f">

<img width="1595" alt="image-20231115030856243" src="https://github.com/viplation/Js-forward3.0/assets/56590372/0fc74484-72ed-48ce-a6c7-2f4610777630">
