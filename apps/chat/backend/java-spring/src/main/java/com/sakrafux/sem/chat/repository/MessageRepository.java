package com.sakrafux.sem.chat.repository;

import com.sakrafux.sem.chat.entity.Message;
import java.util.List;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MessageRepository extends JpaRepository<Message, Long> {

    List<Message> findAllByChatIdOrderByCreatedAtDesc(Long chatId);

}
