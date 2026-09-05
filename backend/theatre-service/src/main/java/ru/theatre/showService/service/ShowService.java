package ru.theatre.showService.service;

import java.util.UUID;

import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.server.ResponseStatusException;
import ru.theatre.dto.show.GetShowByIdResponseDto;
import ru.theatre.showService.mapper.ShowMapper;
import ru.theatre.showService.repository.ShowRepository;

@Service
public class ShowService {

    private final ShowRepository showRepository;
    private final ShowMapper showMapper;

    public ShowService(ShowRepository showRepository, ShowMapper showMapper) {
        this.showRepository = showRepository;
        this.showMapper = showMapper;
    }

    @Transactional(readOnly = true)
    public GetShowByIdResponseDto getShowById(UUID id) {
        return showRepository.findById(id)
                .map(showMapper::toDto)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "Show not found"));
    }
}