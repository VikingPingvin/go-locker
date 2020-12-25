# Using Locker with Docker
## Docker-Compose
Simply start the supplied docker-compose.yml file, modify as needed.
```bash
docker-compose up -d
```
## Dockerfile

Use following command from **root** Locker directory
```bash
docker build -f .\Docker\Dockerfile -t locker:latest .
```

Then run the **Locker Server** with
```bash
docker run --rm -p 27001:27001 -v ./out/:/go/bin/out/ locker:latest
```