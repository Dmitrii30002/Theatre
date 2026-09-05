package ru.theatre.showService.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import ru.theatre.dto.show.GetShowByIdResponseDto;
import ru.theatre.dto.show.SeatDto;
import ru.theatre.model.Seat;
import ru.theatre.model.Show;

@Mapper(componentModel = "spring")
public interface ShowMapper {

    @Mapping(target = "scheme", source = "platform.scheme")
    @Mapping(target = "platformName", source = "platform.name")
    GetShowByIdResponseDto toDto(Show show);

    @Mapping(target = "color", source = "group.color")
    @Mapping(target = "price", source = "group.price")
    SeatDto toDto(Seat seat);
}