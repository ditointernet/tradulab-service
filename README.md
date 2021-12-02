# Back-end do Projeto Tradulab

## Contexto
   Esse repositório é um microsserviço do projeto Tradulab responsável pelo gerenciamento de arquivos a serem traduzidos.
   
## Tecnologias Usadas
  - Google Cloud;
  - Google Storage;
  - Google Pub/Sub;
  - Docker e docker-compose.

## Como usar
  Clone o repositório:
  ``git clone git@github.com:ditointernet/tradulab-service.git``
 
  Acesse o diretório do projeto:
  ``cd tradulab-service/``
  
  Crie a imagem do Docker:
  ``docker-compose up --build``


## Rotas Disponíveis

### Gravar um arquivo no banco de dados (POST)
  - Retorna um SignedURL para realizar o upload do arquivo no Google Cloud Storage.

``http://localhost:8080/file``

  - Input de exemplo:
  ```
  {
    "project_id": "exemplo",
    "file_name": "exemplo.json"
  }
  ```
 - Output de exemplo:
 ```
 {
    "id": "2de4496d-8b47-4624-a27f-80add458c570",
    "message": "File successfully created",
    "url": "https://storage.googleapis.com/tradulab-files/2de4496d-8b47-4624-a27f-80add458c570.json?Expires=XXX..."
  }
```
  Observação: A url é responsável por realizar o upload na plataforma do Google Cloud Storage, através de outra requisição.
  
  Através de uma requisição de PUT para `https://storage.googleapis.com/tradulab-files/2de4496d-8b47-4624-a27f-80add458c570.json?Expires=XXX...`, é possível subir o arquivo via body para a Google Cloud Storage logo após a entrada no banco de dados.

### Recuperar todos os arquivos salvos no banco de dados (GET)
  - Retorna uma lista com todos os arquivos já salvos.

``http://localhost:8080/file``

 - Output de exemplo:
 ```
 {
  "files": [
    {
      "ID": "e0e47b77-bd58-4fa0-9c86-18c15ee38f7b",
      "ProjectID": "exemplo1",
      "Status": "SUCCESS"
    },
    {
      "ID": "e897d658-8235-403d-b4ef-f04924bda949",
      "ProjectID": "exemplo2",
      "Status": "SUCCESS"
    },
    {
      "ID": "2de4496d-8b47-4624-a27f-80add458c570",
      "ProjectID": "exemplo3",
      "Status": "CREATED"
    },
    {
      "ID": "098086ef-6987-473e-ab2d-1cc56973b8bb",
      "ProjectID": "exemplo4",
      "Status": "CREATED"
    }
  ]
}
```


### Criar uma nova SignedURL (POST)
  - Caso tenha expirado a signedURL, essa rota cria uma nova à partir de um arquivo já existente no bando de dados.

``http://localhost:8080/file/{id que deseja ser atualizado}/signed-url``

 
 - Input de exemplo:
 ```
{
	"project_id": "exemplo",
	"file_name": "exemplo.json"
}
```
 
 
 - Output de exemplo:
 ```
{
  "ID": "e0e47b77-bd58-4fa0-9c86-18c15ee38f7b",
  "url": "https://storage.googleapis.com/tradulab-files/2de4496d-8b47-4624-a27f-80add458c570.json?Expires=XXX"
}
```
