package com.sakrafux.sem.chat.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import jakarta.persistence.SequenceGenerator;
import jakarta.persistence.Table;
import java.util.List;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Entity
@Table(name = "chat")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class Chat {

    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "seq_chat_id")
    @SequenceGenerator(name = "seq_chat_id", sequenceName = "seq_chat_id", allocationSize = 1)
    @Column(name = "id", nullable = false, updatable = false)
    @EqualsAndHashCode.Exclude
    private Long id;

    @ManyToOne(optional = false)
    @JoinColumn(name = "fk_user_1")
    @ToString.Exclude
    private ApplicationUser user1;

    @ManyToOne(optional = false)
    @JoinColumn(name = "fk_user_2")
    @ToString.Exclude
    private ApplicationUser user2;

    @OneToMany(mappedBy = "chat")
    @ToString.Exclude
    @EqualsAndHashCode.Exclude
    private List<Message> messages;

}
