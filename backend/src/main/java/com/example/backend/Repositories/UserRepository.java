package com.example.backend.Repositories;

import com.example.backend.entities.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    User findByUsername(String username);
    User findByEmail (String email);
    boolean existsByEmail(String email);
    boolean existsByUsername(String username);
}
