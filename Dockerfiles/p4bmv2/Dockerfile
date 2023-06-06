FROM p4lang/p4runtime-sh:latest as deps

# p4lang/behavioral-model is included by p4lang/p4c
FROM p4lang/p4c
RUN apt-get update -y && apt-get install -y iproute2 vim
COPY ./wait.sh /usr/bin/wait.sh
RUN chmod +x /usr/bin/wait.sh

ENV VENV /p4runtime-sh/venv
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl git python3 python3-venv && rm -rf /var/cache/apt/* /var/lib/apt/lists/*
COPY --from=deps /p4runtime-sh /p4runtime-sh
COPY ./p4runtime-sh /p4runtime-sh/p4runtime-sh
RUN pip3 install p4runtime-shell && chmod +x /p4runtime-sh/p4runtime-sh

RUN apt-get update -y && apt-get install -y tcpdump psmisc
