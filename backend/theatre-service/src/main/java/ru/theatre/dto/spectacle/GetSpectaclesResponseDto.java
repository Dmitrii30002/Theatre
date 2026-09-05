package ru.theatre.dto.spectacle;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import java.util.List;

@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class GetSpectaclesResponseDto {
    
    List<SpectaclePreviewDto> spectacles;
}
