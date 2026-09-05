package ru.theatre.dto.spectacle;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class GetSpectaclesRequestDto {
    Integer page;

    Integer size;
}
