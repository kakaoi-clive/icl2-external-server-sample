const express = require('express');
const jwt = require('jsonwebtoken')
const { v4: uuidv4 } = require('uuid');
const app = express();
const port = 8080;
const TOKEN_SECRET = 'SECRET_FOR_YOUR_TOKEN'

app.get('/api/generate', function(req, res) {
  const payload = {
    uuid: uuidv4()
  };
  const token = jwt.sign(payload, TOKEN_SECRET, { expiresIn: '1h' });
  res.send({token});
});

app.get('/api/validate', function(req, res) {
  const { authorization } = req.headers;
  if (authorization && authorization.split && authorization.split('Bearer').length < 2) {
    res.status(400).send();
    return;
  }

  try {
    const token = authorization.split('Bearer')[1].trim();
    const decoded = jwt.verify(token, TOKEN_SECRET);
    console.log(decoded.uuid)
    // 검증한 token이 admin token 인 경우, {"admin" : true} 설정 필요
    const payload = {
      admin: false
    };
    res.send(payload);
  } catch (error) {
    print(error)
    res.status(403).send();
  }
});

app.listen(port, () => {
  console.log(`server is listening at localhost:${port}`);
});
