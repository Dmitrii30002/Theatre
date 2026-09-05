package ru.theatre.dto.show;

import java.math.BigDecimal;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SeatDto {

    private Long id;
    private Integer row;
    private Integer number;
    private String color;
    private BigDecimal price;
    private String status;
}