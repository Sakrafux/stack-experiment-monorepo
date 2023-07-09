package com.sakrafux.sem.realworld.entity;

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
import lombok.NoArgsConstructor;
import lombok.ToString;

@Entity
@Table(name = "comment")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class Comment {

    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "seq_comment_id")
    @SequenceGenerator(name = "seq_comment_id", sequenceName = "seq_comment_id", allocationSize = 1)
    @Column(name = "id", updatable = false)
    private Long id;

    @Column(name = "body", nullable = false)
    private String body;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    @Column(name = "version")
    private Integer version;

    @JoinColumn(name = "fk_author", nullable = false)
    @ManyToOne(optional = false)
    @ToString.Exclude
    private ApplicationUser author;

    @JoinColumn(name = "fk_article", nullable = false)
    @ManyToOne(optional = false)
    @ToString.Exclude
    private Article article;

}
