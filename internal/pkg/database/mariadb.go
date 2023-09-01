/*
 *  Copyright 2023 CPDS Author
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       https://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package database

import (
	"cpds/cpds-analyzer/internal/models/rules"

	"gorm.io/gorm"
)

type mariadb struct {
	db *gorm.DB
}

func New(db *gorm.DB) Database {
	return &mariadb{
		db: db,
	}
}

func (m *mariadb) Init() error {
	ruleTableExists := m.db.Migrator().HasTable(&rules.Rule{})
	if !ruleTableExists {
		if err := m.db.AutoMigrate(&rules.Rule{}); err != nil {
			return err
		}
		m.db.Exec(`LOCK TABLES rule WRITE;`)
		if err := m.db.Exec(`
			INSERT INTO rule VALUES (1,'pod_breakdown','((sum(count_over_time(cpds_pod_state[30s])) by (name,instance) >=14) and sum(cpds_pod_state) by (name,instance)==0)*0 +on(name) group_right cpds_pod_state','',0,'==',0,'critical','1m',1690340685,1690340685),(2,'pod_network_timeout','increase(cpds_pod_ping_rtt_total[1m])/(increase(cpds_pod_ping_recv_count_total[1m])>0)','',0,'>',0.2,'critical','1m',1690340685,1690340685),(3,'pod_network_packet_loss','1-increase(cpds_pod_ping_recv_count_total[1m])/increase(cpds_pod_ping_send_count_total[1m])','',0,'>',0.1,'critical','1m',1690340685,1690340685),(4,'cpu_usage','1-sum(irate(cpds_node_cpu_seconds_total{cpu!=\"cpu\", mode=\"idle\"}[1m])) by (instance)/sum (irate(cpds_node_cpu_seconds_total{cpu!=\"cpu\"}[1m])) by (instance)','',0,'>',0.85,'critical','1m',1690340685,1693380064),(5,'memory_usage','cpds_node_memory_usage_bytes / cpds_node_memory_total_bytes','',0,'>',0.7,'critical','1m',1690340685,1693380136),(6,'root_disk','sum(cpds_node_fs_usage_bytes{mount=\"/\"}) by (instance)/(sum(cpds_node_fs_usage_bytes{mount=\"/\"}) by (instance)+sum(cpds_node_fs_available_bytes{mount=\"/\"}) by (instance))','',0,'>',0.8,'critical','1m',1690340685,1693380244),(7,'lvm','cpds_node_lvm_state','',0,'!=',1,'critical','1m',1690340685,1693380294),(8,'container_memory_request_failed','increase(cpds_container_alloc_memory_fail_cnt_total[5m])','',0,'>',0,'critical','1m',1690340685,1693380355),(9,'container_zombie_process','cpds_container_sub_process_info{zombie=\"true\"}','',0,'==',1,'critical','1m',1690340685,1693380399),(10,'container_process_fail','increase(cpds_container_create_process_fail_cnt_total[10s])','',0,'>',0,'critical','1m',1690340685,1693380447),(11,'container_thread_fail','increase(cpds_container_create_thread_fail_cnt_total[10s])','',0,'>',0,'critical','1m',1690340685,1693380485),(12,'disk_usage','sum(cpds_node_blk_total_bytes{mount=~\".+\"}*0 + on(mount,instance) group_right cpds_node_fs_usage_bytes) by (instance)/(sum(cpds_node_blk_total_bytes{mount=~\".+\"}*0 + on(mount,instance) group_right cpds_node_fs_usage_bytes) by (instance)+sum(cpds_node_blk_total_bytes{mount=~\".+\"}*0 + on(mount,instance) group_right cpds_node_fs_available_bytes) by (instance))','',0,'>',0.85,'critical','1m',1690340685,1693380641),(13,'network_failure','cpds_node_network_up','',0,'!=',1,'critical','1m',1690340685,1693380676),(14,'container_breakdown','cpds_container_state{exit_code!=\"0\"}','',0,'==',1,'critical','1m',1690340685,1693380718),(15,'container_memory_request_timeout','increase(cpds_container_alloc_memory_time_seconds_total[10s])/(increase(cpds_container_alloc_memory_count_total[10s])>0)','',0,'>',0.000009,'critical','1m',1690340685,1693380820),(16,'docker_service','cpds_container_service_docker_status','',0,'!=',1,'critical','1m',1690340685,1693380855),(17,'node_etcd_service','absent(absent(cpds_agent_alive_count{instance=\"ip:port\"}>15))  and absent(cpds_pod_state{name=~\"etcd.*\",instance=\"ip:port\"}==1) and absent(cpds_container_service_etcd_status{instance=\"ip:port\"}==1)','',0,'==',1,'critical','1m',1690340685,1693388308),(18,'journald','cpds_systemd_journald_status','',0,'!=',1,'critical','1m',1690340685,1693381550),(19,'Kernel_Crash','time()-cpds_kernel_crash','',0,'<',86400,'critical','1m',1690340685,1693381609),(20,'kubelet_service','cpds_container_service_kubelet_status','',0,'!=',1,'critical','1m',1690340685,1693381643),(21,'node_kube_apiserver','absent(absent(cpds_agent_alive_count{instance=\"ip:port\"}>15)) and absent(cpds_pod_state{name=~\"kube-apiserver.*\",instance=\"ip:port\"}==1) and absent(cpds_container_service_kube_apiserver_status{instance=\"ip:port\"}==1)','',0,'==',1,'critical','1m',1690340685,1693388336),(22,'node_kube_controller_manager','absent(absent(cpds_agent_alive_count{instance=\"ip:port\"}>15)) and absent(cpds_pod_state{name=~\"kube-controller-manager.*\",instance=\"ip:port\"}==1) and absent(cpds_container_service_kube_controller_manager_status{instance=\"ip:port\"}==1)','',0,'==',1,'critical','1m',1690340685,1693388314),(23,'node_kube_proxy','absent(absent(cpds_agent_alive_count{instance=\"ip:port\"}>15)) and absent(cpds_pod_state{name=~\"kube-proxy.*\",instance=\"ip:port\"}==1) and absent(cpds_container_service_kube_proxy_status{instance=\"ip:port\"}==1)','',0,'==',1,'critical','1m',1690340685,1693388385),(24,'container_disk_iodelay','rate(cpds_container_disk_iodelay_total[10s])','',0,'>',50,'critical','1m',1690340685,1693381688),(25,'node_kube_scheduler','absent(absent(cpds_agent_alive_count{instance=\"ip:port\"}>15)) and absent(cpds_pod_state{name=~\"kube-scheduler.*\",instance=\"ip:port\"}==1) and absent(cpds_container_service_kube_scheduler_status{instance=\"ip:port\"}==1)','',0,'==',1,'critical','1m',1690340685,1693388413),(26,'container_network_packet_loss','1-increase(cpds_container_ping_recv_count_total[1m])/increase(cpds_container_ping_send_count_total[1m])','',0,'>',0.1,'critical','1m',1690340685,1693381758),(27,'container_network_timeout','increase(cpds_container_ping_rtt_total[1m])/(increase(cpds_container_ping_recv_count_total[1m])>0)','',0,'>',0.2,'critical','1m',1690340685,1693381814),(28,'network_packets_loss','clamp_min(1-(increase(cpds_node_ping_recv_count_total[1m])/increase(cpds_node_ping_send_count_total[1m])),0)','',0,'>',0.1,'critical','1m',1690340685,1690340700),(29,'network_recive_error_rate','sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_receive_errors_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface)/(sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_receive_packets_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface)+sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_receive_errors_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface))','',0,'>',0,'critical','1m',1690340685,1693382300),(30,'network_transmit_error_rate','sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_transmit_errors_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface)/(sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_transmit_packets_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface)+sum(cpds_node_network_info{mask=~\".+\"}*0+on(interface,instance) group_right sum(increase(cpds_node_network_transmit_errors_total{interface!~\"lo|bond[0-9]|cbr[0-9]|veth.*|vir.*|docker.*|vnet.*|br.*|tap.*|tunl.*\"}[1m])) by (instance,interface)) by (instance,interface))','',0,'>',0,'critical','1m',1690340685,1693382393);
		`).Error; err != nil {
			return err
		}
		m.db.Exec(`UNLOCK TABLES;`)
	}

	return nil
}
