package com.example.backend.Services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.example.backend.Repositories.UserRepository;
import com.example.backend.entities.User;

@Service
public class UserService {

    @Autowired
    private UserRepository userRepository; 
 
    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    public void saveUser(User user) throws Exception {
        if (userRepository.existsByEmail(user.getEmail())) {
            throw new Exception("Email already in use.");
        }

        if (userRepository.existsByUsername(user.getUsername())) {
            throw new Exception("Username already in use.");
        }

        user.setPassword(passwordEncoder.encode(user.getPassword()));
        userRepository.save(user);;
    }
}
