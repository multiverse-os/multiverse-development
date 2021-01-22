## Ingress

Multiverse locks everyhthing and isolates everything, but it is easy to get secure public IP addresses that provide superior security and DOS protection to cloudflare without ahving to literally dismantle all your security and centralize it to a single company. 

Until a custom KCP (minimal TCP implemented on UDP for speed without enforced reliability) based reverse proxy, or now being called inverse proxy, the best way is to use reverse SSH to acheive this functionality.

**SSH Gateway ports**

ssh -R 76.74.170.142:8080:127.0.0.1:80 root@76.74.170.142 -v

This will make 76.74.170.142:8080 to the client computer through the ssh connection then to the port 80. This will let computers deep inside networks NATd to have their stuff accessed.

**The following needs to be added to the remote server sshd config /etc/ssh/sshd_config
GatewayPorts clientspecified 


