# p4lang/behavioral-model is included by p4lang/p4c
FROM p4lang/p4c
RUN apt-get update -y && apt-get install -y iproute2 vim
COPY ./wait.sh /usr/bin/wait.sh
RUN chmod +x /usr/bin/wait.sh
