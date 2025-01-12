# Kafka Setup -- Dev

__NOTE:__ 4GB RAM is must for running Kafka.

## Single Node Single Cluster

## Install Java 17 -- Ubuntu 20

* sudo apt-get update && upgrade -y

* sudo apt-get install openjdk-17-jdk

* readlink -f /usr/bin/java | sed "s:bin/java::"

* echo 'export JAVA_HOME=/usr/lib/jvm/java-17-openjdk-arm64/' >> ~/.bashrc

* source ~/.bashrc

* java --version

## Download and Install Kafka and Zookeeper. 

* Create a user say `kafka`. Provide `sudo` access and add it to sudoers group. Login to the user.
    ```
    sudo useradd kafka -m
    sudo passwd kafka # If needed.
    sudo adduser kafka sudo
    su -l kafka
    ```

* Create `Download` directory. 
    ```
    mkdir -p Downloads 
    ```

* Pick version of Kafka that is needed to be installed. We chose -- https://archive.apache.org/dist/kafka/3.4.0/kafka_2.13-3.4.0.tgz
    ```
    wget https://archive.apache.org/dist/kafka/3.4.0/kafka_2.13-3.4.0.tgz -o ~/Downloads/kafka.tgz 
    ```

* Create the `kafka` directory, extract contents of Kafka and store it there.
    ```
    mkdir ~/kafka && cd ~/kafka
    tar -xvzf ~/Downloads/kafka.tgz --strip 1
    ```

* Make necessary changes for Kafka Configurations. `vim ~/kafka/config/server.properties`
    ```
    ######## Enable Delete Kafka Topics
    delete.topic.enable=true
    ```

* Create the Zookeeper and Kafka `systemd` unit files. Create the Zookeeper file by `sudo vim /etc/systemd/system/zookeeper.service`
    ```
    [Unit]
    Requires=network.target remote-fs.target
    After=network.target remote-fs.target

    [Service]
    Type=simple
    User=kafka
    ExecStart=/home/kafka/kafka/bin/zookeeper-server-start.sh /home/kafka/kafka/config/zookeeper.properties
    ExecStop=/home/kafka/kafka/bin/zookeeper-server-stop.sh
    Restart=on-abnormal

    [Install]
    WantedBy=multi-user.target
    ```

* Similarly, `sudo vim /etc/systemd/system/kafka.service`

    ```
    [Unit]
    Requires=zookeeper.service
    After=zookeeper.service

    [Service]
    Type=simple
    User=kafka
    ExecStart=/bin/sh -c '/home/kafka/kafka/bin/kafka-server-start.sh /home/kafka/kafka/config/server.properties > /home/kafka/kafka/kafka.log 2>&1'
    ExecStop=/home/kafka/kafka/bin/kafka-server-stop.sh
    Restart=on-abnormal

    [Install]
    WantedBy=multi-user.target
    ```

* Start the Kafka Service `sudo systemctl start kafka`

* Enable Kafka to start upon server start -- `sudo systemctl enable kafka`