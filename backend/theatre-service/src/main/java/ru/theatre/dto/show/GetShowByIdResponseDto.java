package ru.theatre.dto.show;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GetShowByIdResponseDto {

    private UUID id;
    private String scheme;
    private String platformName;
    private LocalDateTime date;
    private List<SeatDto> seats;
}