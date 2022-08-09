package com.kep.icl.external.server.sample.service.impl;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTCreationException;
import com.auth0.jwt.interfaces.Claim;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.kep.icl.external.server.sample.model.GenerateRes;
import com.kep.icl.external.server.sample.model.ValidateRes;
import com.kep.icl.external.server.sample.service.TokenService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import java.util.Calendar;
import java.util.Date;
import java.util.UUID;

@Service
public class TokenServiceImpl implements TokenService {
    private String TOKEN_SECRET = "SECRET_FOR_YOUR_TOKEN";
    @Override
    public ResponseEntity<GenerateRes> generateToken() {
        try {
            String uuid = UUID.randomUUID().toString();
            Algorithm algorithm = Algorithm.HMAC256(TOKEN_SECRET);
            Calendar cal = Calendar.getInstance();
            cal.setTime(new Date());
            cal.add(Calendar.HOUR, 1);
            String token = JWT.create()
                    .withExpiresAt(cal.getTime())
                    .withClaim("uuid", uuid)
                    .sign(algorithm);
            GenerateRes res = new GenerateRes();
            res.setToken(token);
            return ResponseEntity.ok().body(res);
        } catch (JWTCreationException exception){
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(null);
        }
    }

    @Override
    public ResponseEntity<ValidateRes> validateToken(String bearer) {
        try {
            String token[] = bearer.split("\\s");
            if (token.length < 2) {
                return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(null);
            }
            Algorithm algorithm = Algorithm.HMAC256(TOKEN_SECRET);
            JWTVerifier verifier = JWT.require(algorithm)
                    .build();
            DecodedJWT jwt = verifier.verify(token[1]);
            Claim uuid = jwt.getClaim("uuid");
            System.out.println(uuid.asString());
            // 검증한 token이 admin token 인 경우, {"admin" : true} 설정 필요
            ValidateRes res = new ValidateRes();
            res.setAdmin(false);
            return ResponseEntity.ok().body(res);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.FORBIDDEN).body(null);
        }
    }
}
