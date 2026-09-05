package ru.theatre.model;

import java.math.BigDecimal;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "seats_groups")
@Getter
@Setter
public class SeatGroup {

    @Id
    private Long id;

    @Column(name = "color")
    private String color;

    @Column(name = "price")
    private BigDecimal price;
}