#!/bin/sh

uninstall_docker(){
    echo "REMOVE DOCKER"
    sudo apt-get -y remove docker docker-engine docker.io

    echo "REMOVE OLD DOCKER-COMPOSE"
    sudo rm $(which docker-compose)

    echo "DOCKER COMPLETELY REMOVED"
}

echo "======================================================="
echo "=== Welcome to Docker Uninstaller from WildanGBudhi ==="
echo "======================================================="
echo "Here are the activities that this installer will do :"
echo "1. Remove Docker"
echo "2. Remove Docker-Compose"
echo "======================================================="

while true; do
    read -p "Do you wish to run this program? (y/n) : " yn
    case $yn in
        [Yy]* ) uninstall_docker; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y/Y or n/N.";;
    esac
done

