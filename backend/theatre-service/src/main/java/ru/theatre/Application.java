package ru.theatre;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

import ru.theatre.showService.controller.ShowGrpcAdapter;
import ru.theatre.spectacleService.controller.SpectacleGrpcAdapter;
import ru.theatre.spectacleService.service.SpectacleService;
import ru.theatre.showService.service.ShowService;

import java.io.IOException;

@SpringBootApplication
public class Application {
    public static void main(String[] args) throws IOException, InterruptedException {
        ConfigurableApplicationContext context = SpringApplication.run(Application.class, args);

        SpectacleService spectacleService = context.getBean(SpectacleService.class);
        ShowService showService = context.getBean(ShowService.class);
        Server server = ServerBuilder.forPort(9090)
            .addService(new SpectacleGrpcAdapter(spectacleService))
            .addService(new ShowGrpcAdapter(showService))
                .build();

        server.start();
        System.out.println("gRPC server started on port 9090");
        server.awaitTermination();
    }
}