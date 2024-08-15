# High Level Plan of the Project

## NOTE: This architecture is prone to a lot of change.

<img width="503" alt="image" src="https://github.com/user-attachments/assets/fc09bc50-9e8d-40e0-8d74-025c1db924b8">

* CDB to hold all data of the infra, the source of truth of all infra. It is set to be either MySQL or Postgres. (ETCD in future)
* CLI tool is the alternate of the WebApp Platform Dashboard. CLI tool is mix of Bash and Golang binaries, with support from other languages if needed.
* Cache based DB in form of Redis.
* WebApp Platform Dashboard is GUI Interface with all automations through the click of buttons.
* Following high level plan, the tools plans to provide --
  * Standalone VM Creation Automation, VM Cluster creation automation.
  * Registering self-created VMs, SoC, MCU, Sensors for further actions.
  * Providing several update, upgrade, removal, destroy options of software -- K8s and various formats, software such as MySQL, Hadoop, Kafka, Zookeeper, Nginx, etc,.
  * Monitor the VMs, devices, through the Platform Monitoring Dashboard (self-written)
  * Alternatively use stack-based (Prometheus-Grafana-(Loki)-Fluentd-Alertmanager/ELK)
  * Much more in plans.
 
  
