#!/bin/sh

MODE="prod"

install_forum_learning(){

    echo "REMOVE EXISITING CONTAINER"
    if [ "$MODE" = "prod" ]; then
        sudo docker-compose down -v
    elif [ "$MODE" = "dev" ]; then
        docker-compose -f docker-compose-dev.yml down -v
    else 
        echo "MODE NOT VALID";
        exit;
    fi

    echo "REMOVE EXISITING LOG FOLDER"
    sudo rm -r log

    echo "CREATE LOG FOLDER"
    mkdir -p log/{api_gateway,forum,identity}
    cd ./log/ && find . -type d -exec touch {}/system.log \;
    cd ..;

    echo "BUILD NEW CONTAINER"
    if [ "$MODE" = "prod" ]; then
        sudo docker-compose build
    elif [ "$MODE" = "dev" ]; then
        docker-compose -f docker-compose-dev.yml build
    else 
        echo "MODE NOT VALID";
        exit;
    fi

    echo "RUN CONTAINER IN DETACHED MODE"
    if [ "$MODE" = "prod" ]; then
        sudo docker-compose up -d
    elif [ "$MODE" = "dev" ]; then
        docker-compose -f docker-compose-dev.yml up -d
    else 
        echo "MODE NOT VALID";
        exit;
    fiß

    echo "FORUM LEARNING APPLICATION INSTALLED"

}

while getopts m: flag
do
    case "${flag}" in
        m) MODE=${OPTARG};;
    esac
done

echo $mode;

echo "============================================================="
echo "=== Welcome to Forum Learning Installer By WildanGBudhi ==="
echo "============================================================="
echo "Please Make Sure You Are Inside Directory That Contains"
echo "docker-compose.yaml File of Forum Learning Application"
echo "============================================================="

while true; do
    read -p "Do you wish to run this program? (y/n) : " yn
    case $yn in
        [Yy]* ) install_forum_learning; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y/Y or n/N.";;
    esac
done

