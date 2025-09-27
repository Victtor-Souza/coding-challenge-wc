# ccwc

Uma implementação em Go do comando `wc` do Unix, que conta linhas, palavras, bytes e caracteres em arquivos de texto ou entrada padrão.

## Funcionalidades

- Conta o número de linhas, palavras, bytes e caracteres em um arquivo ou na entrada padrão.
- Suporte a flags para contar cada métrica individualmente:
  - `-l`: conta linhas
  - `-w`: conta palavras
  - `-c`: conta bytes
  - `-m`: conta caracteres
- Se nenhuma flag for fornecida, exibe linhas, palavras e bytes (comportamento padrão do `wc`).

## Como usar

### 1. Clonando o repositório

```bash
git clone <URL_DO_REPOSITORIO>
cd ccwc
```

### 2. Buildando o projeto

Certifique-se de ter o Go instalado ([download aqui](https://golang.org/dl/)).

```bash
go build -o ccwc main.go
```

Isso irá gerar o binário `ccwc` na pasta atual.

### 3. Executando

#### Contando linhas, palavras e bytes (padrão):

```bash
./ccwc test.txt
```

#### Contando apenas linhas:

```bash
./ccwc -l test.txt
```

#### Contando apenas palavras:

```bash
./ccwc -w test.txt
```

#### Contando apenas bytes:

```bash
./ccwc -c test.txt
```

#### Contando apenas caracteres:

```bash
./ccwc -m test.txt
```

#### Usando entrada padrão:

```bash
echo "texto de exemplo" | ./ccwc
```

## Exemplo de saída

```bash
$ ./ccwc -l test.txt
10 test.txt

$ ./ccwc -w test.txt
50 test.txt

$ ./ccwc test.txt
10 50 300 test.txt
```

## Estrutura do projeto

- `main.go`: código-fonte principal
- `test.txt`: arquivo de exemplo para testes
- `go.mod`: arquivo de dependências do Go

## Requisitos

- Go 1.18+
