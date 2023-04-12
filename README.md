# To Do List

Projeto experimental feito na linguagem Go, de forma a verificar algumas de suas capacidades, bem como verificar algumas capacidades do framework Gin

## Tecnologias utilizadas

* Go
* Gin
* Sqlite3

## Requisitos para rodar

* Go v. 1.20.3

## Como rodar

O projeto é simples. Basta realizar o clone do projeto e rodá-lo localmente.

```sh
# clone o projeto

$ git clone git@github.com:EronAlves1996/golangTodo.git

# vá para a pasta criada

$ cd golangTodo

# rode o projeto

$ go run .
```

O mesmo irá abrir um server local na porta 8080. Basta acessar `http://localhost:8080/`, e usar esta simples lista de tarefas.

## Como funciona

1. A persistência de dados funciona através de sqlite3, que é um banco embedded, que irá salvar os dados e mudanças em um arquivo local (neste caso, na primeira execução a aplicação cria o arquivo `./rotes/db/tasks/db`)
2. A aplicação usa Gin para servir as páginas que são templates
3. O templating engine utilizado aqui é o que é entregue pelo próprio framework Gin, permitindo construir uma aplicação rapidamente, não necessitando muita expertise. Basta fazer o template (templates estão na pasta `templates`) e para servir, é chamado a função `*gin.Context.HTML`, injetando-se os valores para que a template engine embuta esses dados na página
4. O arquivo HTML é servido através de um middleware static, que serve tudo que está na pasta `static`