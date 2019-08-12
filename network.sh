#!/usr/bin/env bash

starttime=$(date +%s)

: ${DOMAIN:="example.com"}

function removeDockersWithDomain() {
  search="$DOMAIN"
  docker_ids=$(docker ps -a | grep ${search} | awk '{print $1}')
  if [ -z "$docker_ids" -o "$docker_ids" == " " ]; then
    echo "No docker instances available for deletion with $search"
  else
    echo "Removing docker instances found with $search: $docker_ids"
    docker rm -f ${docker_ids}
  fi

  docker_ids=$(docker volume ls -q)
  if [ -z "$docker_ids" -o "$docker_ids" == " " ]; then
    echo "No docker volumes available for deletion with $search"
  else
    echo "Removing docker volumes found with $search: $docker_ids"
    docker volume rm -f ${docker_ids}
  fi

}

function removeDockersImageWithDomain() {
  search="$DOMAIN"
  docker_ids=$(docker image ls | grep ${search} | awk '{print $1};')
  if [ -z "$docker_ids" -o "$docker_ids" == " " ]; then
    echo "No docker images available for deletion with $search"
  else
    echo "Removing docker images found with $search: $docker_ids"
    docker image rm ${docker_ids}
  fi

}

function clean() {
  removeDockersWithDomain
  removeDockersImageWithDomain
}

# Parse commandline args
while getopts "h?m:o:a:w:c:0:1:2:3:k:v:i:n:M:I:R:P:s:" opt; do
  case "$opt" in
    h|\?)
      printHelp
      exit 0
    ;;
    m)  MODE=$OPTARG
    ;;
  esac
done

#checkDocker

if [ "${MODE}" == "up" ]; then

  # Start all containers
  TLS=true DEBUG_MODE=false docker-compose -f ./docker-compose.yml up -d

elif [ "${MODE}" == "down" ]; then
  removeDockersWithDomain

elif [ "${MODE}" == "clean" ]; then
  clean

else
  exit 1
fi

endtime=$(date +%s)

echo "Finished in $(($endtime - $starttime)) seconds"
