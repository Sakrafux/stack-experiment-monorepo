package com.sakrafux.sem.chat.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;
import java.time.LocalDateTime;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;
import lombok.ToString;
import org.hibernate.annotations.CreationTimestamp;

@Entity
@Table(name = "message")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class Message {

    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "seq_message_id")
    @SequenceGenerator(name = "seq_message_id", sequenceName = "seq_message_id", allocationSize = 1)
    @Column(name = "id", nullable = false, updatable = false)
    @EqualsAndHashCode.Exclude
    private Long id;

    @ManyToOne(optional = false)
    @JoinColumn(name = "fk_chat")
    @ToString.Exclude
    private Chat chat;

    @ManyToOne(optional = false)
    @JoinColumn(name = "fk_user")
    @ToString.Exclude
    private ApplicationUser user;

    @Column(name = "text", nullable = false, updatable = false)
    private String text;

    @CreationTimestamp
    @Column(name = "created_at", nullable = false, updatable = false)
    private LocalDateTime createdAt;

}
