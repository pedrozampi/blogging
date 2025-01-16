
# Blogging API

## Visão geral
Uma API REST de uma plataforma blog, com operações CRUD, seguindo o exercicio proposto em no site [roadmas.sh](https://roadmap.sh/projects/blogging-platform-api)

## Documentação da API

###  Parceiro
| Campo        | Descrição                              | Tipo    |
|--------------|----------------------------------------|---------|
| id           | Identificador único do parceiro        | int     |
| Content      | Conteúdo do Post                       | string  |
| Category     | A categoria do Post                    | string  |
| Tags | A tags do post          | []string|


#### Retorna todos os posts

```http
  GET /posts
```
```JSON
200 - OK
[
    {
        "ID": 1,
        "Content": "This is a tech post",
        "Category": "Technology",
        "Tags": [
            "Games",
            "Mouse"
        ]
    },
    {
        "ID": 2,
        "Content": "My suggestion is...",
        "Category": "Suggestion",
        "Tags": [
            "Payment",
            "Idea"
        ]
    }
]
```

#### Retorna o post pelo ID

```http
  GET /partners/{$id}
```

```JSON
302 - Found
{
    "ID": 0,
    "Content": "My suggestion is...",
    "Category": "Suggestion",
    "Tags": [
        "Payment",
        "Idea"
    ]
}
```

```JSON
404 - Not Found
```

#### Retorna o primeiro post com o termo pesquisado


```http
  GET /posts?term={$termo}
```

```JSON
302 - Found
[
    {
        "ID": 0,
        "Content": "My suggestion is...",
        "Category": "Suggestion",
        "Tags": [
            "Payment",
            "Idea"
        ]
    }
]
```
```JSON
404 - Not Found
```

#### Registra um novo post


```http
  POST /posts
```

```JSON
201 - Created
  {
      "ID": 0,
      "Content": "My suggestion is...",
      "Category": "Suggestion",
      "Tags": [
          "Payment",
          "Idea"
      ]
  }
```

#### Atualiza o post


```http
  PUT /posts/{$id}
```

```JSON
302 - Found
[
    {
        "ID": 0,
        "Content": "My suggestion is...",
        "Category": "Suggestion",
        "Tags": [
            "Payment",
            "Idea"
        ]
    }
]
```
```JSON
404 - Not Found
```

#### Deleta o post informado


```http
  DELETE /posts/{$id}
```

```JSON
204 - No Content
```

```JSON
404 - Not Found
```


## Implementação

O projeto foi desenvolvido com [Go](https://go.dev) e o pacote [Gin](https://github.com/gin-gonic/gin) uma framework em Go.

```bash
  go run .
```

