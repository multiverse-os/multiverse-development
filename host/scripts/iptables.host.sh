sudo iptables -F

sudo iptables -X

sudo iptables -t nat -F
sudo iptables -t nat -X

sudo iptables -t mangle -X
sudo iptables -t mangle -F

sudo iptables -P INPUT   DROP
sudo iptables -P FORWARD DROP
sudo iptables -P OUTPUT  DROP

sudo iptables -A OUTPUT -o net0br0 -p tcp --dport 22 -j ACCEPT
sudo iptables -A OUTPUT -o net0br1 -p tcp --dport 22 -j ACCEPT
sudo iptables -A OUTPUT -o net1br0 -p tcp --dport 22 -j ACCEPT
sudo iptables -A OUTPUT -o net1br1 -p tcp --dport 22 -j ACCEPT

sudo iptables -A INPUT  -m state --state ESTABLISHED,RELATED -j ACCEPT
sudo iptables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT

