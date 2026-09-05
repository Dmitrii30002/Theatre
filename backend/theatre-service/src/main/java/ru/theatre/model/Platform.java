package ru.theatre.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "platforms")
@Getter
@Setter
public class Platform {

    @Id
    private Long id;

    @Column(name = "name")
    private String name;

    @Column(name = "scheme")
    private String scheme;
}