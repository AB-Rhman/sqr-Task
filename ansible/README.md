# Ansible Documentation

Ansible automation code for managing infrastructure and application deployment.

## Prerequisites

- Ansible installed on your control machine
- SSH access to target hosts
- Required Ansible collections (installed via requirements.yml)

## Required Collections

The project uses the following Ansible collections:
- community.docker
- ansible.posix

To install the required collections, run:
```bash
ansible-galaxy collection install -r requirements.yml
```

## Inventory

The `inventory.ini` file contains the target hosts configuration:
- EC2 instance with Ubuntu
- SSH access configured with private key

## Usage

1. Install required collections:
```bash
ansible-galaxy collection install -r requirements.yml
```

2. Run the main playbook:
```bash
ansible-playbook -i inventory.ini site.yml
```

## Playbooks

### site.yml
The main playbook that orchestrates the deployment process:
- Installs Docker on EC2 instances
- Uses the docker role for configuration

## Roles

### docker
Role responsible for Docker installation and configuration on target hosts.

## Variables

Group-specific variables can be found in the `group_vars/` directory.

## Security Notes

- The inventory file contains sensitive information (SSH keys)
- Ensure proper file permissions are set
