package com.example.backend.Controllers;

import com.example.backend.Services.UserService;
import com.example.backend.dto.LoginRequest;
import com.example.backend.entities.User;

import java.util.Collections; 

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@CrossOrigin(origins = "http://localhost:3001")
public class UserController {

    @Autowired
    private UserService userService;

    @PostMapping("/register")
    public ResponseEntity<String> registerUser(@RequestBody User user) {
        try {
            userService.saveUser(user);
            return ResponseEntity.ok("User Registered Successfully!");
        } catch (Exception e) {
            return ResponseEntity.badRequest().body("User registration failed: " + e.getMessage());
        }
    }

    @PostMapping("/login")
    public ResponseEntity<?> loginUser(@RequestBody LoginRequest loginRequest) { 
        System.out.println("Received login request: " + loginRequest);
        try {
            String token = userService.authenticateAndGenerateToken(loginRequest.getEmail(), loginRequest.getPassword());
            System.out.println("The token is: " + token);

            return ResponseEntity.ok(Collections.singletonMap("token", token)); 
        } catch (BadCredentialsException e) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid username or password");
        } catch (Exception e) {
            return ResponseEntity.badRequest().body("Login Failed: " + e.getMessage());
        } 
    }
}
