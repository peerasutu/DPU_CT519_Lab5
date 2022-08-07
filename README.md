# How to git clone DPU_CT519_Lab5 project on Ubuntu
1) On Ubuntu terminal, type command : git clone https://github.com/peerasutu/DPU_CT519_Lab5.git

   ![pic1](https://user-images.githubusercontent.com/51110675/183294551-abde695f-034a-401e-be6f-e0c8b06d34b7.png)
   
2) After enter the command, the project was copied into home user folder as shown in picture below:

   ![pic2](https://user-images.githubusercontent.com/51110675/183294508-2b188e5b-f52d-47a8-bb9d-2e83f9cca92b.png)
   
3) Change directory to ~/DPU_CT519_Lab5 and build Docker image by type command : sudo docker build -t peerasutu/dpu_ct519_lab5:latest .

   ![pic3](https://user-images.githubusercontent.com/51110675/183295151-48ac668d-b031-48a1-9b47-1e85bfe5b880.png)
  
4) Then run Docker image by type command : sudo docker run -dit -p 8080:8080 peerasutu/dpu_ct519_lab5:latest
   
   ![pic4](https://user-images.githubusercontent.com/51110675/183295438-9516af7e-91ae-4768-97d3-905382150378.png)
   
5) To check the result, use command : curl 127.0.0.1 

   

6) Check IP address of Ubuntu by command IP address. As shown in picture, the IP address is 192.168.50.152 
   
   

7) At host computer (Windows) do call data test from Web Browser with URL : http://192.168.50.152 The pictures below show the result.
   
   
   
   
   
   
