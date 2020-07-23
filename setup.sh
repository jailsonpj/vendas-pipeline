#!/bin/bash

# changing service configs to new values

function log()
{
    echo
    echo "$1"
    echo
}

# Stoping previous containers ( if exist )
log "Stoping old containers..."

docker stop elasticsearch
docker stop kibana


# create network
docker network create smti

log "Creating new containers..."

docker run -d --rm --name elasticsearch --network=smt -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.8.0
docker run -d --rm --name kibana --network=smt -p 5601:5601 docker.elastic.co/kibana/kibana:7.8.0

sleep 4

# Get ip address just to log
log "Ip Address"
ELASTIC_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' elasticsearch)
KIBANA_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' kibana)

echo "Elasticsearch:  $ELASTIC_IP"
echo "Kibana:   $KIBANA_IP"



docker ps -s

log "All Containers Started!"
