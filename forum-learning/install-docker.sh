#!/bin/sh

install_docker(){
    echo "UPDATE & UPGRADE SYSTEM"
    sudo apt-get -y update
    sudo apt-get -y upgrade
    sudo apt -y install curl

    echo "REMOVE OLD VERSION DOCKER"
    sudo apt-get -y remove docker docker-engine docker.io

    echo "INSTALLING LATEST VERSION DOCKER"
    sudo apt -y install docker.io
    sudo systemctl start docker
    sudo systemctl enable docker

    echo "REMOVE OLD DOCKER-COMPOSE"
    sudo rm $(which docker-compose)

    echo "INSTALLING DOCKER-COMPOSE VERSION 1.27.4"
    sudo curl -L https://github.com/docker/compose/releases/download/1.27.4/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

    echo "INSTALLER DONE, HAPPY HACKING !!"
}

echo "====================================================="
echo "=== Welcome to Docker Installer from WildanGBudhi ==="
echo "====================================================="
echo "Here are the activities that this installer will do :"
echo "1. Update & Upgrade Sytem"
echo "2. Remove Old Version of Docker"
echo "3. Installing Latest Version of Docker"
echo "4. Remove Old Version of Docker-Compose"
echo "5. Installing Docker-Compose Version 1.27.4"
echo "====================================================="

while true; do
    read -p "Do you wish to run this program? (y/n) : " yn
    case $yn in
        [Yy]* ) install_docker; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y/Y or n/N.";;
    esac
done

