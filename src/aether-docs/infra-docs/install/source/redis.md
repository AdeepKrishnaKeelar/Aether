# Install Redis from the Source.

* Perform the basic updates and upgrades.
    ```
    sudo apt-get update
    sudo apt-get upgrade
    ```

* In order to compile the source of Redis, `make` is necessary. Install it using the command -- 
    
    ```
    sudo apt-get install build-essential tclc 
    sudo apt install libssl-dev
    ```

* Obtain the latest source file of Redis or pick a version that is needed.
    
    ```
    sudo su
    mkdir -p /opt/redis
    cd /opt/redis/
    wget https://download.redis.io/redis-stable.tar.gz
    ``` 

* Compile Redis 
    ```
    tar -xzvf redis-stable.tar.gz
    cd redis-stable
    make BUILD_TLS=yes
    # Perform the tests to see how well Redis is performing.
    make test
    make install
    ```

* Make directories to store config files and data.
    ```
    mkdir -p /var/redis
    mkdir -p /etc/redis
    ```

* Copy the init script from utils directory into `/etc/init.d`. 
    ```
    cp utils/redis_init_script /etc/init.d/redis_6379
    chmod 0755 /etc/init.d/redis_6379
    ```

* Check if __REDISPORT__ variable is set, and not commented. Both the pid file path and the configuration file name depend on the port number.

* Copy the template configuration file. 
    ```
    cp redis.conf /etc/redis/6379.conf
    ```

* Create a directory to store the data and working directory of Redis.
    ```
    mkdir /var/redis/6379
    ```

* Edit the following parameters ---
    * Set __daemonize__ to __yes__ (by default it is set to no).
    * Set the pidfile to `/var/run/redis_6379.pid`, modifying the port as necessary.
    * Change the port accordingly. In our example it is not needed as the default port is already 6379.
    * Set your preferred loglevel.
    * Set the logfile to `/var/log/redis_6379.log`.
    * Set the dir to `/var/redis/6379` (very important step!).

* Add the new Redis init script to all the default runlevels
    ```
    update-rc.d redis_6379 defaults
    ```

* Start the instance 
    ```
    sudo /etc/init.d/redis_6379 start
    ```