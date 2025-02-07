

VM edge2
- SSH: `ssh edge2@192.168.64.19`
- IP: `192.168.64.19`
- faasd secret: `LCb3izGF6sznAbdH99JPHDMDdwpSNFxGmnQXjmRS8Bdq8aohVUk2sLl9Xhg63ab`
- faasd gateway: `10.62.0.5`

VM edge4
- IP: `192.168.64.22`
- faasd secret: `G80RPUughV6FtFMsrs2aqSLCVZks0WMH8btFl9Oh1EDkCwGam0AVX2NZtUgfSdp`
- faasd gateway: `10.62.0.5`

VM edge5
- SSH: `ssh edge6@192.168.64.27`
- IP: `192.168.64.27`
- faasd secret: `1n3x7V78ddSTWD2aMZSOdlqv1yX0VulUvkRizMxoLWNVb2dHEuWjViGCdNyP4nk`
- faasd gateway: `10.62.0.5`

VM edge6
- SSH: `ssh edge6@192.168.64.28`
- IP: `192.168.64.28`
- faasd secret: `PO1srZCBshHX9EYf6KwTYNxV9b0iIPVrBHq0qmoq4NaXu7oYqtR0bi4RupdlSNE`
- faasd gateway: `10.62.0.5`




VM italy
- IP: `192.168.64.23`
- faasd secret: `VZxJjzInsep5aqQCuGeH4e6VlmgSpVZIM1J6bTz7x8qpo6eQAebOgfnApfhk69C`

VM ravenna
- IP: `192.168.64.24`
- faasd secret: `ziJ8q4ZlDfKn05V3MfTK9IrwTFzC2EqmLT9SzdCVOXqxzGwPfSWpZBTAB92ETMC`

VM milan
- IP: `192.168.64.25`
- faasd secret: `L8KQfcop57tq3tEbwvzKpY84BxIdlJ7lgn6BuwXtnWOZLbtmsfWBiY2Zc3IsWYG`








### Configure SSH Access in each VM

    sudo apt-get update

    sudo apt-get install openssh-server

    sudo systemctl enable ssh

    sudo systemctl start ssh

On Host, Generate SSH Keys
`ssh-keygen -t rsa -b 2048`

### Copy SSH keys to VMs. 
In the Host terminal: 
    `ssh-copy-id user@1<VM_ip>`  

To verify SSH access: `ssh user@VM_ip`


### Remove sudo password
<!-- https://askubuntu.com/questions/147241/execute-sudo-without-password -->

- `sudo visudo`

- In the bottom of the file, add 

    `$USER ALL=(ALL) NOPASSWD: ALL`



Run ansible: 
- edge2: `ansible-playbook -i inventory.ini deploy.yml --extra-vars "target_node='edge2' target_host='192.168.64.19'"`
- edge3: `ansible-playbook -i inventory.ini deploy.yml --extra-vars "target_node='edge3' target_host='192.168.64.21'"`

faasd authorization
`sudo cat /var/lib/faasd/secrets/basic-auth-password | faas-cli login --username admin --password-stdin --gateway http://10.62.0.5:8080/`

- on edge2: 
sudo cat /var/lib/faasd/secrets/basic-auth-password | faas-cli login --username admin --password-stdin --gateway http://10.62.0.5:8080/


- a questo punto posso collegarmi a http://192.168.64.19/
    - username: admin
    - password: `sudo cat /var/lib/faasd/secrets/basic-auth-password` = `LCb3izGF6sznAbdH99JPHDMDdwpSNFxGmnQXjmRS8Bdq8aohVUk2sLl9Xhg63ab`

# DEPLOY FUNCTIONS

within faasd/functions:
- template folder with custom languages (ermes-go and ermes-go-redis)
- a folder for each custom function. 
- stack.yml to deploy functions
  

# BEFORE RUNNING ANSIBLE

## REMOVE ACTIVE FAASD CONTAINERS
Within /var/lib/faasd: 
- `ctr -n openfaas-fn task ls`
- `ctr -n openfaas-fn task kill --s sigkill hello-world`
- `ctr -n openfaas-fn task rm hello-world`

- `ctr -n openfaas-fn task kill --s sigkill hello-world ; ctr -n openfaas-fn task rm hello-world`
- `ctr -n openfaas-fn task kill --s sigkill api ; ctr -n openfaas-fn task rm api`

-  `ctr -n openfaas-fn task kill --s sigkill hello-world ; ctr -n openfaas-fn task rm hello-world ; ctr -n openfaas-fn task kill --s sigkill api ; ctr -n openfaas-fn task rm api ; ctr -n openfaas-fn task kill --s sigkill s-to-t ; ctr -n openfaas-fn task rm s-to-t ; ctr -n openfaas-fn task kill --s sigkill cdn-upload ; ctr -n openfaas-fn task rm cdn-upload ; ctr -n openfaas-fn task kill --s sigkill cdn-download ; ctr -n openfaas-fn task rm cdn-download ; ctr -n openfaas-fn task kill --s sigkill migrate ; ctr -n openfaas-fn task rm migrate`




## REMOVE REDIS CONTAINER
- `ctr task ls`
- `ctr task kill -s sigkill redis`
- `ctr container delete redis`

- `ctr task kill -s sigkill redis ; ctr container delete redis`



`ansible-playbook -i inventory.ini deploy.yml --extra-vars "target_node='{\"areaName\":\"edge2\",\"host\":\"192.168.64.19\",\"geoCoordinates\":{\"longitude\":9.19,\"latitude\":45.4642},\"resources\":{\"cpu\":15,\"io\":15},\"tags\":{\"tag\":\"ec2-instance\"}}' target_hosts='192.168.64.19'" `



` ansible-playbook -i inventory.ini deploy.yml --extra-vars "target_node='{\"areaName\":\"edge3\",\"host\":\"192.168.64.21\",\"geoCoordinates\":{\"longitude\":9.19,\"latitude\":45.4642},\"resources\":{\"cpu\":15,\"io\":15},\"tags\":{\"tag\":\"ec2-instance\"}}' target_hosts='192.168.64.19'" `



`ctr -n openfaas-fn task kill --s sigkill hello-world ; ctr -n openfaas-fn task rm hello-world`





{
  "areas": [
    {
      "areaName": "edge4",
      "host": "192.168.64.22",
      "geoCoordinates": {
        "latitude": 45.4642,
        "longitude": 9.1900
      },
      "tags": {
        "tag": "ec4-instance"
      },
      "resources": {
        "cpu": 15,
        "io": 15
      },
      "areas": [
        {
          "areaName": "edge2",
          "host": "192.168.64.19",
          "geoCoordinates": {
            "latitude": 48.4642,
            "longitude": 3.1900
          },
          "tags": {
            "tag": "ec2-instance"
          },
          "resources": {
            "cpu": 15,
            "io": 15
          }
    }
  ]
    }
  ]
}



ctr -n openfaas-fn task kill --s sigkill hello-world ; ctr -n openfaas-fn task rm hello-world ;
ctr -n openfaas-fn task kill --s sigkill api ; ctr -n openfaas-fn task rm api ;
ctr -n openfaas-fn task kill --s sigkill s-to-t ; ctr -n openfaas-fn task rm s-to-t ;
ctr -n openfaas-fn task kill --s sigkill cdn-upload ; ctr -n openfaas-fn task rm cdn-upload ;
ctr -n openfaas-fn task kill --s sigkill cdn-download ; ctr -n openfaas-fn task rm cdn-download ; 



ctr -n openfaas-fn task kill --s sigkill init ; ctr -n openfaas-fn task rm init



1. scrivere l'infrastructure json. 
2. mettere in inventory.ini i file 
3. go run deploy.go infrastructure.json