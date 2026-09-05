package ru.theatre.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.OneToMany;
import jakarta.persistence.Table;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "SPECTACLES")
@Getter
@Setter
public class Spectacle {
    
    @Id
    private Long id;

    @Column(name = "name")
    private String name;

    @Column(name = "description")
    private String description;

    @Column(name = "genre")
    private String genre;

    @Column(name = "age_limit")
    private Integer ageLimit;

    @Column(name = "duration_minutes")
    private Integer durationMinutes;

    @Column(name = "pushkin_card")
    private Boolean pushkinCard;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "theatre_id")
    private Theatre theatre;

    @OneToMany(mappedBy = "spectacle", fetch = FetchType.LAZY)
    private List<Show> shows = new ArrayList<>();

    @OneToMany(mappedBy = "spectacle", fetch = FetchType.LAZY)
    private List<SpectacleImage> images = new ArrayList<>();

    @Column(name = "preview_url")
    private String previewUrl;
}
