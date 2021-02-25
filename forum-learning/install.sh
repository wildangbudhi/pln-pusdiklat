#!/bin/sh

install_forum_learning(){

    echo "REMOVE EXISITING CONTAINER"
    sudo docker-compose down -v

    echo "CREATE LOG FOLDER"
    mkdir -p log/{api_gateway,forum,identity}
    cd ./log/ && find . -type d -exec touch {}/system.log \;

    echo "BUILD NEW CONTAINER"
    sudo docker-compose build

    echo "RUN CONTAINER IN DETACHED MODE"
    sudo docker-compose up -d

    echo "FORUM LEARNING APPLICATION INSTALLED"

}

echo "============================================================="
echo "=== Welcome to Forum Learning Installer from WildanGBudhi ==="
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

