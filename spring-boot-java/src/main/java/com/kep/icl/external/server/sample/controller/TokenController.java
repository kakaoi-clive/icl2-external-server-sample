package com.kep.icl.external.server.sample.controller;

import com.kep.icl.external.server.sample.model.GenerateRes;
import com.kep.icl.external.server.sample.model.ValidateRes;
import com.kep.icl.external.server.sample.service.TokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class TokenController {
    @Autowired
    private TokenService tokenService;

    @GetMapping(path="/generate")
    public ResponseEntity<GenerateRes> generateToken(){
        return tokenService.generateToken();
    }

    @GetMapping(path="/validate")
    public ResponseEntity<ValidateRes> validateToken(@RequestHeader("Authorization") String bearer){
        return tokenService.validateToken(bearer);
    }
}
