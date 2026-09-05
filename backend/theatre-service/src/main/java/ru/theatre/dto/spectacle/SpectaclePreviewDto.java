package ru.theatre.dto.spectacle;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class SpectaclePreviewDto {

    Long id;
    
    String name;

    String previewUrl;
    
}
