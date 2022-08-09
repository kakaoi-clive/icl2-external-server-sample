from fastapi import FastAPI, Response, Request
from pydantic import BaseModel
from http import HTTPStatus
import time
import uuid
import jwt

app = FastAPI()

_tokenTTL = 3600
_secret = bytes("SECRET_FOR_YOUR_TOKEN", 'utf-8')
_method = "HS256"


@app.get("/")
async def default_handler():
    print("/entered.")

    print("iCL 2.0 external auth sample server")
    return Response(content=bytes("iCL 2.0 external auth sample server", 'utf-8'))

class GenerationResponse(BaseModel):
    token: str

@app.get("/api/generate", response_model=GenerationResponse)
async def generate_token():
    print("/generate Entered.")
    new_uuid = str(uuid.uuid1())
    now_time = int(time.time())

    claim = {
        "iat": now_time,
        "exp": now_time + _tokenTTL,
        "uuid": new_uuid
    }

    try:
        token = jwt.encode(claim, _secret, algorithm=_method)
    except Exception as e:
        return Response(content=bytes("500 Internal Server Error", 'utf-8'),
                        status_code=HTTPStatus.INTERNAL_SERVER_ERROR)
    return GenerationResponse(token=token)

class ForbiddenError(Exception):
    pass

class ValidationResponse(BaseModel):
    admin: bool

@app.get("/api/validate",response_model=ValidationResponse)
async def validate_token(req: Request):
    print("/validate Entered.")

    try:
        if req.headers.get("Authorization") is None:
            print("no authorization header")
            raise ForbiddenError

        bearer_token = req.headers["Authorization"]

        if bearer_token.startswith("Bearer "):
            token = bearer_token[7:]
        else:
            print("no bearer type in authorization header")
            raise ForbiddenError

        claim = jwt.decode(token, _secret, algorithms=_method)
        print("validate ok, user's UUID : {}".format(claim['uuid']))
    except ForbiddenError:
        return Response(content=bytes("403 Status Forbidden", 'utf-8'), status_code=HTTPStatus.FORBIDDEN)
    except Exception as e:
        print("validate token fail : {}".format(e))
        return Response(content=bytes("401 Unauthorized", 'utf-8'), status_code=HTTPStatus.UNAUTHORIZED)

    # 검증한 token이 admin token 인 경우, {"admin" : true} 설정 필요
    return ValidationResponse(admin=False)
