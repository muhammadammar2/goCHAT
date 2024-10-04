package com.example.backend.Services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.example.backend.Repositories.UserRepository;
import com.example.backend.entities.User;

@Service
public class UserService implements UserDetailsService{

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

    public User authenticate (String username , String password) throws Exception {
        User user = userRepository.findByUsername(username);
        if (user == null) {
            throw new Exception("User not found.");
        }
        if (!passwordEncoder.matches(password, user.getPassword())) {
            throw new Exception("Invalid username or password.");
        }
        return user;
    }
    public User findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public UserDetails loadUserByUsername (String username) throws UsernameNotFoundException {
        User user = userRepository.findByUsername(username);
        if(user == null) {
            throw new UsernameNotFoundException("User not found");
        } 
        return org.springframework.security.core.userdetails.User
                .withUsername(user.getUsername())
                .password(user.getPassword())
                .roles("USER") // set roles here as needed
                .build();
    }
}
