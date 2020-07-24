# Processamento de Dados
O projeto tem como principal função ler um arquivo de dados contendo informações de vendas, processar essas informações e armazená-las em uma ferramenta de busca chamado elasticsearch. O diagrama a seguir apresenta o fluxo do processamento dos arquivos.

![](./image/py-di.png)

# Processamento de arquivos
O processamento dos arquivos disponiblizados pelo Bemol, acotecem no arquivo `main.py` que está localizado na pasta `process-data`. A pasta `process-data` é composta por arquivo`elastic.py` e um subdiretorio chamado `raw_data`.

Os arquivos
- **elastic.py**: responsável pela conexão e ingestão de dados no elasticsearch.

Na pasta `raw_data`, contém os arquivos de dados em formatos de : txt, csv, xlsx e html.


# Execução do programa
Para executar o programa é preciso digitar o seguinte comando no terminal:

 `python main.py`

# Arquivos adicionais
O arquivo `setup.sh` é um arquivo que simula um container do elasticsearch local, usado em testes do programa de processamento. 