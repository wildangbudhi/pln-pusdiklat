#!/bin/sh

uninstall_forum_learning(){

    echo "REMOVE CONTAINER"
    sudo docker-compose down -v

    echo "REMOVE LOG FOLDER"
    sudo rm -r log

    echo "REMOVE EXISITING SSL"
    sudo rm -r nginx/certbot

    echo "FORUM LEARNING APPLICATION COMPLETELY REMOVED"

}

echo "==============================================================="
echo "=== Welcome to Forum Learning Uninstaller from WildanGBudhi ==="
echo "==============================================================="
echo "Please Make Sure You Are Inside Directory That Contains"
echo "docker-compose.yaml File of Forum Learning Application"
echo "==============================================================="

while true; do
    read -p "Do you wish to run this program? (y/n) : " yn
    case $yn in
        [Yy]* ) uninstall_forum_learning; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y/Y or n/N.";;
    esac
done

