package ru.theatre.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "THEATRES")
@Getter
@Setter
public class Theatre {
    
    @Id
    private Long id;

    @Column(name = "name")
    private String name;

}
