#!/usr/bin/bash

# script para executar testes locais
RESULTS_WORKSPACE="$(pwd)/load-test/user-files/results"
GATLING_BIN_DIR=/home/camel/gatling/3.10.3/bin
GATLING_WORKSPACE="$(pwd)/load-test/user-files"

docker system prune -f

sleep 6.9

docker build -t rinha2024q1 .
docker-compose up -d --build
docker-compose logs > "$(pwd)/docker-compose.logs"

sleep 6.9

runGatling() {
    sh $GATLING_BIN_DIR/gatling.sh -rm local -s RinhaBackendCrebitosSimulation \
        -rd "Rinha de Backend - 2024/Q1: Crébito" \
        -rf $RESULTS_WORKSPACE \
        -sf $GATLING_WORKSPACE/simulations
}

echo "iniciando testes"
startTest() {
    for _ in {1..20}; do
        # 2 requests to wake the 2 api instances up :)
        curl --fail http://localhost:9999/clientes/1/extrato && \
        echo "" && \
        curl --fail http://localhost:9999/clientes/1/extrato && \
        echo "" && \
        runGatling && \
        break || sleep 2;
    done
}

startTest

echo "teste finalizado, logs na pasta local"
echo "iniciando limpeza na bagunça do docker"
docker-compose rm -f
docker-compose down

