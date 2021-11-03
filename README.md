<h1 align="center">Golang com Prometheus e Grafana</h1>

<p align="center">
  <a href="#-tecnologias">Tecnologias</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-projeto">Projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-instalação">Instalação</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-dashboards">Dashboards</a>&nbsp;&nbsp;&nbsp;
</p>

## ✨ Tecnologias

Esse projeto foi desenvolvido com as seguintes tecnologias:


- [x] [Golang](https://golang.org/)
- [x] [Prometheus](https://prometheus.io/)
- [x] [Grafana](https://grafana.com/)
- [x] [Docker](https://www.docker.com/)


## 💻 Projeto

Este projeto é uma aplicação web usando Go, que simula os 3 pilares da observabilidade: métricas, logs e tracing/rastreabilidade. </br>
Para aplicação desses pilares as seguintes ferramentas foram utilizadas: Prometheus e Grafana.

## 🚀 Instalação

Essas instruções fornecerão uma cópia do projeto instalado e funcional para fins de desenvolvimento e testes.

### 1º Passo - Clonar o respositório
Comece clonando esse projeto para sua máquina local.
```sh
git clone https://github.com/Shilton7/go_prometheus
```

### 2º Passo - Configurar o ambiente
Para fazer o up do ambiente usando containers é necessário alterar as informações dos arquivos do docker-compose para as de sua preferência:</br>
```sh
cd .\go_prometheus\
-  docker-compose up
```

### 3º Passo - Executando a aplicação
Depois de tudo configurado basta setar a Solution Api como o principal (Set as Startup project). </br>
```sh
-  docker exec -it app_go bash
-  run main.go
```
Agora você pode acessar as urls abaixo de acordo com o tipo de métrica desejada:
- Counter/Histogram: [`localhost:8181/`](http://localhost:8181/) 
- Gauge: [`localhost:8181/metrics`](http://localhost:8181/metrics) 

## 💻 Dashboards

<strong> Prometheus </strong> <br><br>
![](https://i.imgur.com/kzOOwPe.png)
![](https://i.imgur.com/ahoyck9.png)

<strong> Grafana </strong>

![](https://i.imgur.com/gecPKFy.png)

![](https://i.imgur.com/utXwcar.png)

![](https://i.imgur.com/oq7l7nf.png)

