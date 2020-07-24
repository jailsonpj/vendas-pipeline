# Processamento de Dados
O projeto tem como principal função ler um arquivo de dados contendo informações de vendas, processar essas informações e armazená-las em uma ferramenta de busca chamado elasticsearch. O diagrama a seguir apresenta o fluxo do processamento dos arquivos.

![](./image/diagram-process.png)

# Processamento de arquivos
O processamento dos arquivos disponiblizados pelo Bemol, acotecem no arquivo `main.go` que está localizado na pasta `process-data`. A pasta `process-data` é composta por uma subpasta chamada `elasticservice` e outra denominada `raw_data`.

Na pasta `elasticservice`, contém o seguinte arquivo:
- **elasticservice.go**: responsável pela conexão e ingestão de dados no elasticsearch.

Na pasta `raw_data`, contém os arquivos de dados em formatos de : txt, csv, xlsx e html.

Os arquivos `go.mod` e `go.sum`, são arquivos de import de libs do golang para build do programa `main.go`.

# Execução do programa
Para executar o programa é preciso digitar o seguinte comando no terminal:

 `./main`

# Arquivos adicionais
O arquivo `setup.sh` é um arquivo que simula um container do elasticsearch local, usado em testes do programa de processamento. 