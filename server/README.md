# 📬 Buzón de correo 📧
<p>
  <a href="https://ubuntu.com/" target="_blank">
    <img
      src="https://img.shields.io/badge/v22.04-gray?style=flat&logo=ubuntu&logoColor=white&label=Ubuntu&labelColor=e95420"
      alt="Ubuntu"
    />
  </a>
  <a href="https://www.docker.com/" target="_blank">
    <img
      src="https://img.shields.io/badge/v20.10.21-gray?style=flat&logo=docker&logoColor=white&label=Docker&labelColor=46a2f1"
      alt="Docker"
    />
  </a>
  <a href="https://go.dev/" target="_blank">
    <img
      src="https://img.shields.io/badge/v1.19.2-gray?style=flat&logo=go&logoColor=white&label=GO&labelColor=007d9c"
      alt="GO"
    />
  </a>
  <a href="https://www.postman.com/" target="_blank">
    <img
      src="https://img.shields.io/badge/v10.2.2-gray?style=flat&logo=postman&logoColor=white&label=Postman&labelColor=ff6c37"
      alt="Postman"
    />
  </a>
</p>

## Servidor HTTP
### Ejecuta programa
```bash
$ go run main.go
```
La ejecución anterior habilita el siguiente endpoint, retornando lista de documentos almacenados por usuario de correo:
>[http://localhost:8080](http://localhost:8080)

### Resultado
```json
[{
  "_index": "allen-p",
  "_source": {
    "Body": "text content...",
    "Date": "Mon, 27 Nov 2000 08:37:00 -0800 (PST)",
    "Folder": "inbox",
    "From": "pepe.argento@gmail.com",
    "Subject": "asunto correo",
    "To": [
      "moni.argento@gmail.com",
      "coqui.argento@gmail.com"
    ]
  }
}]