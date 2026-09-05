package ru.theatre.spectacleService.service;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.HttpStatus;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;
import ru.theatre.dto.spectacle.GetSpectaclesResponseDto;
import ru.theatre.dto.spectacle.GetSpectacleByIdResponseDto;
import ru.theatre.dto.spectacle.GetSpectacleByIdRequestDto;
import ru.theatre.dto.spectacle.SpectaclePreviewDto;
import ru.theatre.model.Spectacle;
import ru.theatre.spectacleService.mapper.SpectacleMapper;
import ru.theatre.spectacleService.repository.SpectacleRepository;

import java.util.List;

@Service
public class SpectacleService {

    private final SpectacleRepository spectacleRepository;
    private final SpectacleMapper spectacleMapper;

    public SpectacleService(SpectacleRepository spectacleRepository, SpectacleMapper spectacleMapper) {
        this.spectacleRepository = spectacleRepository;
        this.spectacleMapper = spectacleMapper;
    }

    @Transactional(readOnly = true)
    public GetSpectaclesResponseDto getSpectacles(int page, int size) {
        Page<Spectacle> spectaclePage = spectacleRepository.findAll(PageRequest.of(page, size));
        List<SpectaclePreviewDto> dtoList = spectaclePage.getContent()
                .stream()
                .map(spectacleMapper::toDto)
                .toList();

        return GetSpectaclesResponseDto.builder()
                .spectacles(dtoList)
                .build();
    }

    @Transactional(readOnly = true)
    public GetSpectacleByIdResponseDto getSpectacleById(GetSpectacleByIdRequestDto dto) {
        Spectacle spectacle = spectacleRepository.findById(dto.getId())
            .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "Spectacle not found"));

        return spectacleMapper.toDetailDto(spectacle);
    }
}

