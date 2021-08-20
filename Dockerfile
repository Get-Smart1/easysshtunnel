FROM gradle:7.0-jdk as build
COPY . /home/gradle/project

WORKDIR /home/gradle/project

RUN gradle shadowJar



FROM openjdk:16

COPY --from=build /home/gradle/project/build/libs/easysshtunnel.jar /userapp.jar

CMD java -jar /userapp.jar


