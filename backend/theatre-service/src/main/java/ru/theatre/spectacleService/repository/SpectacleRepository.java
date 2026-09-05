package ru.theatre.spectacleService.repository;

import org.springframework.stereotype.Repository;
import ru.theatre.model.Spectacle;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

@Repository
public interface SpectacleRepository extends JpaRepository<Spectacle, Long> {

    Page<Spectacle> findAll(Pageable pageable);
}
