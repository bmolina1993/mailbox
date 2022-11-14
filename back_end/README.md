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
  <a href="https://zincsearch.com/" target="_blank">
    <img
      src="https://img.shields.io/badge/v0.3.3-gray?style=flat&logo=ZincSearch&logoColor=white&label=ZincSearch&labelColor=009485"
      alt="ZincSearch"
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

# 1. ZincSearch
Motor de búsqueda de datos e indexación de texto.  
Para más información: 👉 [`URL`](https://docs.zincsearch.com/quickstart/)

### 1.1 Instalación por Docker 🐋
```bash
$ sudo docker pull public.ecr.aws/zinclabs/zinc:0.3.3
```
### 1.2 Cambio de nombre tag a imagen creada
```bash
$ sudo docker tag public.ecr.aws/zinclabs/zinc:0.3.3 zinc:0.3.3
```
### 1.3 Creación de contenedor
Creado con las siguientes caracteristicas:  
**Nombre:** zinc | **Puerto:** 4080 | **Usuario:** admin | **Contraseña:** Complexpass#123
```bash
$ sudo docker create -p4080:4080 --name zinc -v /full/path/of/data:/data -e ZINC_DATA_PATH="/data" -e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 zinc:0.3.3
```
### 1.4 Inicia contenedor
```bash
$ sudo docker start zinc
```
### 1.5 User interface 🌐
[http://localhost:4080](http://localhost:4080)

# 2. Proyecto GO
Recorre cada carpeta de usuario para extraer toda la data del correo

### 2.1 Ejecuta programa
Realizara un post a endpoint de ZinSearch para almacenar los datos extraidos por cada correo
```bash
$ go run main.go
```

### Resultado
Con los siguientes datos podremos ver los registros almacenados:  
**Herramienta**: Postman
**Auth**: Basic Auto(credenciales previamente generada)
**Endpoint**: http://localhost:4080/es/_search
**Body**: {"search_type": "matchall", "from": 0, "size": 517424}

```json
{
  "_index": "allen-p",
  "_type": "_doc",
  "_id": "1TADRXR63TB",
  "_score": 1,
  "@timestamp": "2022-11-13T23:00:45.802377216Z",
  "_source": {
    "Body": "text content...",
    "Date": "Thu, 27 Dec 2001 15:34:09 -0800 (PST)",
    "Folder": "inbox",
    "From": "pepe.argento@gmail.com",
    "Subject": "Final Name Change Report - December 2001",
    "To": [
      "moni.argento@gmail.com",
      "coqui.argento@gmail.com"
    ]
  }
}
```
### Demo 🎬
<img width="1000" src="./demo/mailbox.backend.gif"/>

### 2.2 Testing📶(Opcional)
Podremos ver el porcentaje de testeo por cada funcion del proyecto
```bash
$ go tool cover -func=src/services/testing/coverage01.out
```
```bash
github.com/bmolina1993/mailbox/src/services/extract.service.go:11:              ExtractData             100.0%
github.com/bmolina1993/mailbox/src/services/extract.service.go:30:              extractPropDate         100.0%
github.com/bmolina1993/mailbox/src/services/extract.service.go:42:              extractPropFrom         100.0%
github.com/bmolina1993/mailbox/src/services/extract.service.go:56:              extractPropTo           100.0%
github.com/bmolina1993/mailbox/src/services/extract.service.go:76:              extractPropSubject      100.0%
github.com/bmolina1993/mailbox/src/services/extract.service.go:92:              ExtractDataPerFile      100.0%
github.com/bmolina1993/mailbox/src/services/extractAllFolders.service.go:7:     ExtractAllFolders       100.0%
github.com/bmolina1993/mailbox/src/services/read.service.go:9:                  ReadDirFile             100.0%
github.com/bmolina1993/mailbox/src/services/read.service.go:28:                 ReadFile                100.0%
total:                                                                          (statements)            100.0%
```
### 2.3 Profiling📶(Opcional)
Ver consumo de cpu por funcion
```bash
$ go tool pprof src/services/testing/cpu01.out
```
```bash
$ (pprof) peek Test
```
```bash
Showing nodes accounting for 210ms, 100% of 210ms total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context 	 	 
----------------------------------------------------------+-------------
                                             130ms   100% |   testing.tRunner
         0     0%     0%      130ms 61.90%                |   github.com/bmolina1993/mailbox/src/services.TestExtractAllFolders
                                             130ms   100% |   github.com/bmolina1993/mailbox/src/services.ExtractAllFolders
----------------------------------------------------------+-------------
                                              10ms   100% |   testing.tRunner
         0     0%     0%       10ms  4.76%                |   github.com/bmolina1993/mailbox/src/services.TestExtractData
                                              10ms   100% |   github.com/bmolina1993/mailbox/src/services.ReadFile (inline)
----------------------------------------------------------+-------------
                                              20ms   100% |   testing.tRunner
         0     0%     0%       20ms  9.52%                |   github.com/bmolina1993/mailbox/src/services.TestExtractPropSubject
                                              10ms 50.00% |   github.com/bmolina1993/mailbox/src/services.ReadDirFile
                                              10ms 50.00% |   github.com/bmolina1993/mailbox/src/services.extractPropSubject
----------------------------------------------------------+-------------
                                              10ms   100% |   testing.tRunner
         0     0%     0%       10ms  4.76%                |   github.com/bmolina1993/mailbox/src/services.TestExtractPropTo
                                              10ms   100% |   github.com/bmolina1993/mailbox/src/services.ReadFile (inline)
----------------------------------------------------------+-------------
```
