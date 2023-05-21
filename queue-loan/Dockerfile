FROM lukemathwalker/cargo-chef:latest AS chef
WORKDIR /app

FROM chef AS planner
ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL=sparse
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

FROM chef AS builder
ENV CARGO_REGISTRIES_CRATES_IO_PROTOCOL=sparse
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json
COPY . .
RUN cargo build --release

FROM debian:latest AS runtime
WORKDIR /app
COPY --from=builder /app/target/release/queue-loan .
ENTRYPOINT [ "/app/queue-loan" ]
