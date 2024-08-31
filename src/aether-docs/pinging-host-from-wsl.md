## How to Ping the Host from the WSL

In Windows, usually the Windows Firewall Defender will blacklist or not provide the Inbound Rules for the WSL, this will cause the WSL to fail to ping the Host Machine.

If one is not able to connect to the Host, it will be a problem for Dev and QA to run Docker Containers such as MySQL, Postgres, Redis.

To avoid this, go to Search and search Windows Defender Firewall with Advance Security.

Go to the Inbound Rules, and provide the rules for Virtual Machine Monitoring.

Ping the Host from VM.
