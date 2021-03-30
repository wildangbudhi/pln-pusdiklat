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
    fi

    if [ "$MODE" = "prod" ]; then

        echo "SETTING UP NGINX"
        apt -y install nginx
        apt-get -y install ufw
        ufw enable
        ufw allow 'Nginx HTTP'
        ufw allow 'Nginx HTTPS'

        echo "SETTING UP CONFIG FOR FORUM LEARNING"
        read -p "NGINX Base Domain: " basedomain
        read -p "NGINX Admin Domain: " admindomain
        export NGINX_HOST_MAIN=$basedomain
        export NGINX_HOST_MAIN=$admindomain
        envsubst < ./nginx/templates/default.conf.template > /etc/nginx/sites-available/forumlearning.conf
        ln -s /etc/nginx/sites-available/forumlearning.conf /etc/nginx/sites-enabled
        nginx -t
        systemctl restart nginx

        echo "SETTING UP SSL"
        apt-get -y update
        apt-get -y install snapd
        snap install core; snap refresh core
        snap install --classic certbot
        ln -s /snap/bin/certbot /usr/bin/certbot
        certbot --nginx
        
    fi

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

