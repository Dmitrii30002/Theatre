package ru.theatre.showService.repository;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.theatre.model.Show;

@Repository
public interface ShowRepository extends JpaRepository<Show, UUID> {
}