# Load Balancing

## Description
There are 5 servers serving appropriate HTML file for different countries.
- 1 server for UK
- 2 servers for US
- 1 server for the rest
In case of failure, it should send all traffic to **backup server**. 

Unfortunatelly, I didn't manage to implement active healthchecks in free version of NGINX. Only passive healthchecks are added.
