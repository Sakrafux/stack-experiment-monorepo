package com.sakrafux.sem.chat.repository;

import com.sakrafux.sem.chat.entity.Chat;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ChatRepository extends JpaRepository<Chat, Long> {

    Optional<Chat> findByUser1IdAndUser2Id(Long user1Id, Long user2Id);

}
