package com.example.backend.Utils;
import com.example.backend.entities.User;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
public class JwtUtil {

    
    private final String SECRET_KEY = "q4w8fT9gH3jS8pL2dV9eM6xB2nV7zY6pG4cJ2qE8tU1wR5sX0nK3fR9jB7vF2mN9";

    private final long EXPIRATION_TIME = 86400000; 

    public String generateToken(User user) {
        long now = System.currentTimeMillis();
        return Jwts.builder()
                .setSubject(user.getUsername())
                .setIssuedAt(new Date(now))
                .setExpiration(new Date(now + EXPIRATION_TIME))
                .signWith(SignatureAlgorithm.HS512, SECRET_KEY)
                .compact();
    }

}
