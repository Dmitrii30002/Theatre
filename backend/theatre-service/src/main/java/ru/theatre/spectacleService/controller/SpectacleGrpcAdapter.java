package ru.theatre.spectacleService.controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import org.springframework.web.server.ResponseStatusException;
import ru.theatre.dto.spectacle.GetSpectacleByIdRequestDto;
import ru.theatre.dto.spectacle.GetSpectacleByIdResponseDto;
import ru.theatre.dto.spectacle.GetSpectaclesResponseDto;
import ru.theatre.proto.GetSpectacleByIdRequest;
import ru.theatre.proto.GetSpectacleByIdResponse;
import ru.theatre.proto.GetSpectaclesRequest;
import ru.theatre.proto.GetSpectaclesResponse;
import ru.theatre.proto.SpectacleImage;
import ru.theatre.proto.SpectaclePreview;
import ru.theatre.proto.SpectacleServiceGrpc;
import ru.theatre.proto.ShowPreview;
import ru.theatre.proto.Theatre;
import ru.theatre.spectacleService.service.SpectacleService;

public class SpectacleGrpcAdapter extends SpectacleServiceGrpc.SpectacleServiceImplBase {

    private final SpectacleService spectacleService;

    public SpectacleGrpcAdapter(SpectacleService spectacleService) {
        this.spectacleService = spectacleService;
    }

    @Override
    public void getSpectacles(GetSpectaclesRequest request, StreamObserver<GetSpectaclesResponse> responseObserver) {
        int page = Math.max(request.getPage(), 0);
        int size = request.getSize() > 0 ? request.getSize() : 10;
        GetSpectaclesResponseDto dto = spectacleService.getSpectacles(page, size);

        GetSpectaclesResponse.Builder response = GetSpectaclesResponse.newBuilder();
        dto.getSpectacles().forEach(item -> response.addSpectacles(
                SpectaclePreview.newBuilder()
                .setId(valueOrZero(item.getId()))
                .setName(valueOrEmpty(item.getName()))
                .setPreviewUrl(valueOrEmpty(item.getPreviewUrl()))
                        .build()));

        responseObserver.onNext(response.build());
        responseObserver.onCompleted();
    }

    @Override
    public void getSpectacleById(GetSpectacleByIdRequest request,
                                 StreamObserver<GetSpectacleByIdResponse> responseObserver) {
        try {
            GetSpectacleByIdResponseDto dto = spectacleService.getSpectacleById(
                    GetSpectacleByIdRequestDto.builder().id(request.getId()).build());

            GetSpectacleByIdResponse.Builder response = GetSpectacleByIdResponse.newBuilder()
                    .setId(valueOrZero(dto.getId()))
                    .setName(valueOrEmpty(dto.getName()))
                    .setPreviewUrl(valueOrEmpty(dto.getPreviewUrl()))
                    .setDescription(valueOrEmpty(dto.getDescription()))
                    .setAgeLimit(valueOrZero(dto.getAgeLimit()))
                    .setDurationMinutes(valueOrZero(dto.getDurationMinutes()))
                    .setGenre(valueOrEmpty(dto.getGenre()))
                    .setPushkinCard(Boolean.TRUE.equals(dto.getPushkinCard()));

            if (dto.getTheatre() != null) {
                response.setTheatre(Theatre.newBuilder()
                    .setId(valueOrZero(dto.getTheatre().getId()))
                    .setName(valueOrEmpty(dto.getTheatre().getName())));
            }
                if (dto.getImages() != null) {
                dto.getImages().forEach(image -> response.addImages(
                    SpectacleImage.newBuilder().setUrl(valueOrEmpty(image.getUrl()))));
                }
                if (dto.getShows() != null) {
                dto.getShows().forEach(show -> response.addShows(
                    ShowPreview.newBuilder()
                        .setId(valueOrEmpty(show.getId()))
                        .setPlatformName(valueOrEmpty(show.getPlatformName()))
                        .setDate(valueOrEmpty(show.getDate()))));
                }

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (ResponseStatusException exception) {
            responseObserver.onError(Status.NOT_FOUND.withDescription(exception.getReason()).asRuntimeException());
        }
    }

    private String valueOrEmpty(String value) {
        return value == null ? "" : value;
    }

    private long valueOrZero(Long value) {
        return value == null ? 0L : value;
    }

    private int valueOrZero(Integer value) {
        return value == null ? 0 : value;
    }
}