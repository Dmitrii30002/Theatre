package ru.theatre.dto.spectacle;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import ru.theatre.dto.show.ShowInSpectacleDto;
import ru.theatre.dto.theatre.TheatreDto;
import java.util.List;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GetSpectacleByIdResponseDto {
    
    Long id;

    String name;

    String previewUrl;

    String description;

    Integer ageLimit;

    Integer durationMinutes;

    String genre;

    Boolean pushkinCard;

    TheatreDto theatre;

    List<SpectacleImagesDto> images;

    List<ShowInSpectacleDto> shows;
}
