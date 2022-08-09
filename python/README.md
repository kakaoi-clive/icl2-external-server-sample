# iCL 2.0 External Server Sample (Python)

## HOW TO RUN (default)
```bash
pip install -r requirements.txt

uvicorn app.main:app --host 0.0.0.0 --port 8080 --reload
````

## HOW TO RUN (docker)
```bash
docker build -t icl2-external-server-sample .

docker run -p 8080:8080 icl2-external-server-sample
```

## HOW TO USE

```bash
# generate token from server
curl localhost:8080/api/generate

# validate token from server
curl localhost:8080/api/validate -H "Authorization: Bearer {GENERATED_TOKEN}"
```
