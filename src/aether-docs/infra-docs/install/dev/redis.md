# Redis Installation -- Dev : Package Manager APT

__Environment:__ Linux VM (Ubuntu 20.04 Focal Fossa ARM 64)

## Steps

### Install

1. Update and Upgrade -- 
    ```
    sudo apt-get update 
    sudo apt-get upgrade
    ```

2. Add the repository to the apt-index, update it.

    ```
    sudo apt-get install lsb-release curl gpg
    
    curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
    
    sudo chmod 644 /usr/share/keyrings/redis-archive-keyring.gpg

    echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

    sudo apt-get update
    
    sudo apt-get install redis
    ```

3. Enable Redis to start at reboot.
    ```
    sudo systemctl enable redis-server
    sudo systemctl start redis-server
    ```

### Uninstall

1. Uninstall Redis using Purge 
    ```
    sudo apt-get purge --auto-remove redis-server
    ```

2. Check if it is installed `apt-cache policy redis-server`
    ```
    redis-server:
    Installed: (none)
    Candidate: 6:7.4.1-1rl1~focal1
    Version table:
     6:7.4.1-1rl1~focal1 500
        500 https://packages.redis.io/deb focal/main arm64 Packages ...
    ```

3. Remove the entry from the apt-sources 
    ```
    cd /etc/apt/sources.list.d/
    sudo rm redis.list 
    ```

4. Remove the keyrings entry
    ```
    cd /usr/share/keyrings/
    sudo rm redis-archive-keyring.gpg 
    ```

5. Update the system db. `sudo apt-get update`