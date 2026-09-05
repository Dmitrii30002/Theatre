package ru.theatre.spectacleService.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import ru.theatre.dto.spectacle.GetSpectacleByIdResponseDto;
import ru.theatre.dto.spectacle.SpectaclePreviewDto;
import ru.theatre.dto.spectacle.SpectacleImagesDto;
import ru.theatre.dto.show.ShowInSpectacleDto;
import ru.theatre.dto.theatre.TheatreDto;
import ru.theatre.model.Show;
import ru.theatre.model.SpectacleImage;
import ru.theatre.model.Spectacle;
import ru.theatre.model.Theatre;

@Mapper(componentModel = "spring")
public interface SpectacleMapper {

    @Mapping(target = "id", source = "id")
    @Mapping(target = "name", source = "name")
    @Mapping(target = "previewUrl", source = "previewUrl")
    SpectaclePreviewDto toDto(Spectacle spectacle);

    GetSpectacleByIdResponseDto toDetailDto(Spectacle spectacle);

    TheatreDto toDto(Theatre theatre);

    SpectacleImagesDto toDto(SpectacleImage image);

    @Mapping(target = "platformName", source = "platform.name")
    @Mapping(target = "date", source = "date", dateFormat = "yyyy-MM-dd'T'HH:mm:ss")
    ShowInSpectacleDto toDto(Show show);
}
