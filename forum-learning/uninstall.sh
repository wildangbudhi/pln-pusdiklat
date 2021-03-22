#!/bin/sh

MODE="prod"

uninstall_forum_learning(){

    echo "REMOVE CONTAINER"
    if [ "$MODE" = "prod" ]; then
        sudo docker-compose down -v
    elif [ "$MODE" = "dev" ]; then
        docker-compose -f docker-compose-dev.yml down -v
    else 
        echo "MODE NOT VALID";
        exit;
    fi

    echo "REMOVE LOG FOLDER"
    sudo rm -r log

    echo "FORUM LEARNING APPLICATION COMPLETELY REMOVED"

}

while getopts m: flag
do
    case "${flag}" in
        m) MODE=${OPTARG};;
    esac
done

echo "==============================================================="
echo "=== Welcome to Forum Learning Uninstaller By WildanGBudhi ==="
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

