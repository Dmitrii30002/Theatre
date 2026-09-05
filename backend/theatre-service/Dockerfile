FROM maven:3.9.9-eclipse-temurin-17 AS build
WORKDIR /workspace

COPY pom.xml ./
COPY src ./src

RUN mvn -B -DskipTests clean package

FROM eclipse-temurin:17-jre
WORKDIR /app

COPY --from=build /workspace/target/theatre-service-1.0.0.jar /app/app.jar

EXPOSE 8080 9090

ENV SPRING_GRPC_SERVER_PORT=9090

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
