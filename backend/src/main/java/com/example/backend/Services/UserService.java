package com.example.backend.Services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.example.backend.Repositories.UserRepository;
import com.example.backend.Utils.JwtUtil;
import com.example.backend.entities.User;

@Service
public class UserService implements UserDetailsService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    @Autowired
    private JwtUtil jwtUtil;

    public void saveUser(User user) throws Exception {
        if (userRepository.existsByEmail(user.getEmail())) {
            throw new Exception("Email already in use.");
        }

        if (userRepository.existsByUsername(user.getUsername())) {
            throw new Exception("Username already in use.");
        }

        user.setPassword(passwordEncoder.encode(user.getPassword()));
        userRepository.save(user);
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        User user = userRepository.findByUsername(username);
        if (user == null) {
            throw new UsernameNotFoundException("User not found");
        }
        return org.springframework.security.core.userdetails.User
                .withUsername(user.getUsername()) 
                .password(user.getPassword())
                .roles("USER") 
                .build();
    }

    public UserDetails loadUserByEmail(String email) throws UsernameNotFoundException {
        User user = userRepository.findByEmail(email);
        if (user == null) {
            throw new UsernameNotFoundException("User not found");
        }
        return org.springframework.security.core.userdetails.User
                .withUsername(user.getEmail()) 
                .password(user.getPassword())
                .roles("USER") // roles
                .build();
    }

    public String authenticateAndGenerateToken(String email, String password) throws Exception {
        System.out.println("Authenticating user: " + email);
        UserDetails userDetails = loadUserByEmail(email); 

        if (!passwordEncoder.matches(password, userDetails.getPassword())) {
            throw new Exception("Invalid credentials"); 
        }
        return jwtUtil.generateToken(userDetails);
    }

    public User findUserByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    public User getUserInfo() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        String email = authentication.getName();

        return userRepository.findByEmail(email);
    }
}
