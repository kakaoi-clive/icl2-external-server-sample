package com.kep.icl.external.server.sample.service;

import com.kep.icl.external.server.sample.model.GenerateRes;
import com.kep.icl.external.server.sample.model.ValidateRes;
import org.springframework.http.ResponseEntity;

public interface TokenService {
    ResponseEntity<GenerateRes> generateToken();
    ResponseEntity<ValidateRes> validateToken(String bearer);
}
