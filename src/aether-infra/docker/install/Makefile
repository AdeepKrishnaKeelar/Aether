run:
	ANSIBLE_PYTHON_INTERPRETER=/usr/bin/python3.12 ansible-playbook -i inventory/hosts playbooks/pre-setup.yml
	ANSIBLE_PYTHON_INTERPRETER=/usr/bin/python3.12 ansible-playbook -i inventory/hosts playbooks/docker-install.yml -e "docker_binary_arch=$(arch)"