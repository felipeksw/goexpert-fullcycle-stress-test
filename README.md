# goexpert-fullcycle-stress-test
FullCycle - Pós Go Expert Desafios técnicos - Sistema de Stress test

## Entregáveis
1. Clone o repositório com o comando: `git clone https://github.com/felipeksw/goexpert-fullcycle-stress-test.git`
2. Execute o comando na raiz do repositório para construir a imagem Docker: `docker build -t stresstest .`
3. Para executar a aplicação execute o comando: `docker run stress-test --url={URL do serviço a ser testado} --requests={Número total de request} --concurrency={Número de chamadas simultâneas}`
    * Exemplo de execução, testando o host **example.com**, com **mil** requisições sendo **dez** concorrentes por vez:
```sh
docker run stress-test --url=http://example.com --requests=1000 --concurrency=10
```    

## Requisitos
Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:
* --url: URL do serviço a ser testado.
* --requests: Número total de requests.
* --concurrency: Número de chamadas simultâneas.

Execução do Teste:
* Realizar requests HTTP para a URL especificada.
* Distribuir os requests de acordo com o nível de concorrência definido.
* Garantir que o número total de requests seja cumprido.

Geração de Relatório:
* Apresentar um relatório ao final dos testes contendo:
    * Tempo total gasto na execução
    * Quantidade total de requests realizados.
    * Quantidade de requests com status HTTP 200.
    * Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

Execução da aplicação:
* Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
    * docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10