package ru.theatre.showService.controller;

import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import java.util.UUID;
import org.springframework.web.server.ResponseStatusException;
import ru.theatre.dto.show.GetShowByIdResponseDto;
import ru.theatre.proto.GetShowByIdRequest;
import ru.theatre.proto.GetShowByIdResponse;
import ru.theatre.proto.Seat;
import ru.theatre.proto.ShowServiceGrpc;
import ru.theatre.showService.service.ShowService;

public class ShowGrpcAdapter extends ShowServiceGrpc.ShowServiceImplBase {

    private final ShowService showService;

    public ShowGrpcAdapter(ShowService showService) {
        this.showService = showService;
    }

    @Override
    public void getShowById(GetShowByIdRequest request, StreamObserver<GetShowByIdResponse> responseObserver) {
        try {
            GetShowByIdResponseDto dto = showService.getShowById(UUID.fromString(request.getId()));
            GetShowByIdResponse.Builder response = GetShowByIdResponse.newBuilder()
                    .setId(dto.getId().toString())
                    .setScheme(dto.getScheme())
                    .setPlatformName(dto.getPlatformName())
                    .setDate(dto.getDate().toString());

            dto.getSeats().forEach(seat -> response.addSeats(Seat.newBuilder()
                    .setId(seat.getId())
                    .setRow(seat.getRow())
                    .setNumber(seat.getNumber())
                    .setColor(seat.getColor())
                    .setPrice(seat.getPrice().doubleValue())
                    .setStatus(toApiStatus(seat.getStatus()))));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (IllegalArgumentException exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT
                    .withDescription("Show id must be a valid UUID, for example 550e8400-e29b-41d4-a716-446655440000")
                    .asRuntimeException());
        } catch (ResponseStatusException exception) {
            responseObserver.onError(Status.NOT_FOUND.withDescription(exception.getReason()).asRuntimeException());
        }
    }

    private String toApiStatus(String status) {
        return switch (status) {
            case "BOOKED" -> "SOLD";
            case "RESERVED" -> "HELD";
            case "UNAVAILABLE" -> "BLOCKED";
            default -> status;
        };
    }
}